package password

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerator_Generate(t *testing.T) {
	t.Run("min letters count", func(t *testing.T) {
		g := NewGenerator()
		tests := []GenerateParams{
			{Length: 5, MinLettersCount: 5, IncludeLower: true, IncludeUpper: true, IncludeDigits: true, IncludeSymbols: true},
			{Length: 5, MinLettersCount: 4, IncludeLower: true, IncludeUpper: true, IncludeDigits: true, IncludeSymbols: true},
			{Length: 5, MinLettersCount: 3, IncludeLower: true, IncludeUpper: true, IncludeDigits: true, IncludeSymbols: true},
			{Length: 5, MinLettersCount: 2, IncludeLower: true, IncludeUpper: true, IncludeDigits: true, IncludeSymbols: true},
			{Length: 5, MinLettersCount: 1, IncludeLower: true, IncludeUpper: true, IncludeDigits: true, IncludeSymbols: true},
		}

		for _, tt := range tests {
			pass, err := g.Generate(tt)
			require.NoError(t, err)
			require.Len(t, pass, int(tt.Length))
			var lettersCount uint
			for _, letter := range pass {
				if slices.Contains(lowerLettersDict, letter) || slices.Contains(upperLettersDict, letter) {
					lettersCount++
				}
			}
			require.GreaterOrEqual(t, lettersCount, tt.MinLettersCount)
		}
	})

	t.Run("only digits", func(t *testing.T) {
		g := NewGenerator()
		tests := []GenerateParams{
			{Length: 1, IncludeDigits: true},
			{Length: 5, IncludeDigits: true},
			{Length: 10, IncludeDigits: true},
			{Length: 25, IncludeDigits: true},
			{Length: 50, IncludeDigits: true},
		}

		for _, tt := range tests {
			pass, err := g.Generate(tt)
			require.NoError(t, err)
			require.Len(t, pass, int(tt.Length))
			for _, letter := range pass {
				if !slices.Contains(digitsDict, letter) {
					t.Fatalf("letter %q is not a digit", letter)
				}
			}
		}
	})

	t.Run("only symbols", func(t *testing.T) {
		g := NewGenerator()
		tests := []GenerateParams{
			{Length: 1, IncludeSymbols: true},
			{Length: 5, IncludeSymbols: true},
			{Length: 10, IncludeSymbols: true},
			{Length: 25, IncludeSymbols: true},
			{Length: 50, IncludeSymbols: true},
		}

		for _, tt := range tests {
			pass, err := g.Generate(tt)
			require.NoError(t, err)
			require.Len(t, pass, int(tt.Length))
			for _, letter := range pass {
				if !slices.Contains(symbolsDict, letter) {
					t.Fatalf("letter %q is not a symbol", letter)
				}
			}
		}
	})

	t.Run("custom dictionary", func(t *testing.T) {
		tests := []struct {
			Lower   string
			Upper   string
			Digits  string
			Symbols string
			Length  uint
		}{
			{Length: 15, Lower: "a", Upper: "A", Digits: "5", Symbols: "_"},
			{Length: 25, Lower: "njsdkhgodsu", Upper: "MNVIUSHK", Digits: "12579", Symbols: "$#_|"},
			{Length: 35, Lower: "pobjdh", Upper: "BHFDY", Digits: "76230", Symbols: ")(_-="},
		}

		for _, tt := range tests {
			g := NewGenerator().
				WithLowerLetters(tt.Lower).
				WithUpperLetters(tt.Upper).
				WithDigits(tt.Digits).
				WithSymbols(tt.Symbols)
			pass, err := g.Generate(GenerateParams{
				Length:         tt.Length,
				IncludeLower:   true,
				IncludeUpper:   true,
				IncludeDigits:  true,
				IncludeSymbols: true,
			})
			require.NoError(t, err)
			require.Len(t, pass, int(tt.Length))
			combinedDict := []rune(tt.Lower + tt.Upper + tt.Digits + tt.Symbols)
			for _, letter := range pass {
				if !slices.Contains(combinedDict, letter) {
					t.Fatalf("letter %q should not appear", letter)
				}
			}
		}
	})
}
