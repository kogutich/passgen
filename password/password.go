package password

import (
	"errors"
	"fmt"
)

var (
	ErrWrongLength        = errors.New("length must be greater than zero")
	ErrMinLettersGTLength = errors.New("minimum letters count is greater than length")
	ErrMinLettersMismatch = errors.New("letters dict is empty, bun minimum letters > 0")
	ErrEmptyDict          = errors.New("empty dict")
	ErrReadFromRand       = errors.New("failed to read random sequence")
)

var (
	lowerLettersDict = []rune("abcdefghijklmnopqrstuvwxyz")
	upperLettersDict = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digitsDict       = []rune("0123456789")
	symbolsDict      = []rune("~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/")
)

// Generator generates secure passwords using crypro/rand library.
type Generator struct {
	rnd     *rand
	lower   []rune
	upper   []rune
	digits  []rune
	symbols []rune
}

// NewGenerator creates a new password generator.
func NewGenerator() *Generator {
	return &Generator{
		lower:   lowerLettersDict,
		upper:   upperLettersDict,
		digits:  digitsDict,
		symbols: symbolsDict,
		rnd:     newRand(),
	}
}

// WithLowerLetters sets custom lowercase letters dictionary.
func (g *Generator) WithLowerLetters(lower string) *Generator {
	g.lower = []rune(lower)
	return g
}

// WithUpperLetters sets custom uppercase letters dictionary.
func (g *Generator) WithUpperLetters(upper string) *Generator {
	g.upper = []rune(upper)
	return g
}

// WithDigits sets custom digits dictionary.
func (g *Generator) WithDigits(digits string) *Generator {
	g.digits = []rune(digits)
	return g
}

// WithSymbols sets custom symbols dictionary.
func (g *Generator) WithSymbols(symbols string) *Generator {
	g.symbols = []rune(symbols)
	return g
}

type GenerateParams struct {
	Length          uint
	MinLettersCount uint // from lower_dict + upper_dict
	IncludeLower    bool
	IncludeUpper    bool
	IncludeDigits   bool
	IncludeSymbols  bool
}

// Generate produces a new password.
// It may return an error in case of:
// - invalid read from rand device (i.e. /dev/urandom)
// - wrong parameters.
func (g *Generator) Generate(params GenerateParams) (pass string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%w: %v", ErrReadFromRand, r)
			return
		}
	}()
	lettersDict, othersDict, err := g.buildDicts(params)
	if err != nil {
		return "", err
	}
	var lettersCount, othersCount uint
	if len(othersDict) == 0 {
		lettersCount = params.Length
	} else if len(lettersDict) != 0 {
		lettersCount = g.rnd.UintN(params.Length-params.MinLettersCount+1) + params.MinLettersCount
	}
	othersCount = params.Length - lettersCount

	result := make([]rune, 0, params.Length)
	for i := uint(0); i < lettersCount; i++ {
		result = append(result, lettersDict[g.rnd.IntN(len(lettersDict))])
	}
	for i := uint(0); i < othersCount; i++ {
		result = append(result, othersDict[g.rnd.IntN(len(othersDict))])
	}
	g.rnd.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return string(result), err
}

func (g *Generator) buildDicts(params GenerateParams) ([]rune, []rune, error) {
	if params.Length == 0 {
		return nil, nil, ErrWrongLength
	}
	if params.MinLettersCount > params.Length {
		return nil, nil, ErrMinLettersGTLength
	}
	var letters, others []rune
	if params.IncludeLower {
		letters = append(letters, g.lower...)
	}
	if params.IncludeUpper {
		letters = append(letters, g.upper...)
	}
	if params.IncludeDigits {
		others = append(others, g.digits...)
	}
	if params.IncludeSymbols {
		others = append(others, g.symbols...)
	}
	if len(letters) == 0 && len(others) == 0 {
		return nil, nil, ErrEmptyDict
	}
	if len(letters) == 0 && params.MinLettersCount > 0 {
		return nil, nil, ErrMinLettersMismatch
	}
	return letters, others, nil
}
