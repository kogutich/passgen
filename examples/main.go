package main

import (
	"fmt"

	"github.com/kogutich/passgen/password"
)

func main() {
	s, _ := standard()
	p, _ := pin()
	c, _ := custom()

	fmt.Printf("standard: %s\npin: %s\ncustom: %s\n", s, p, c)
}

func standard() (string, error) {
	gen := password.NewGenerator()
	return gen.Generate(password.GenerateParams{
		Length:          12,
		MinLettersCount: 6,
		IncludeLower:    true,
		IncludeUpper:    true,
		IncludeDigits:   true,
		IncludeSymbols:  true,
	})
}

func pin() (string, error) {
	gen := password.NewGenerator()
	return gen.Generate(password.GenerateParams{
		Length:        4,
		IncludeDigits: true,
	})
}

func custom() (string, error) {
	gen := password.NewGenerator().
		WithLowerLetters("abc").
		WithUpperLetters("DEF").
		WithSymbols("_@").
		WithDigits("123")
	return gen.Generate(password.GenerateParams{
		Length:         36,
		IncludeLower:   true,
		IncludeUpper:   true,
		IncludeDigits:  true,
		IncludeSymbols: true,
	})
}
