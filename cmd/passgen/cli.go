package main

import (
	"io"

	"github.com/alecthomas/kong"
	"github.com/kogutich/passgen/password"
)

type CLI struct { //nolint:govet
	Length            uint   `help:"Password length."                    required:""                                            short:"l"`
	NoTrailingNewline bool   `help:"Do not output the trailing newline." short:"n"`
	MinLettersCount   uint   `default:"0"                                help:"Minimum letters count (lowercase or uppercase)."`
	ExcludeLower      bool   `help:"Exclude lowercase letters."`
	ExcludeUpper      bool   `help:"Exclude uppercase letters."`
	ExcludeDigits     bool   `help:"Exclude digits."`
	ExcludeSymbols    bool   `help:"Exclude symbols."`
	Lower             string `help:"Lowercase letters dictionary."`
	Upper             string `help:"Uppercase letters dictionary."`
	Digits            string `help:"Digits dictionary."`
	Symbols           string `help:"Symbols dictionary."`
}

func (c *CLI) Run(ctx *kong.Context) error {
	gen := password.NewGenerator()
	if c.Lower != "" {
		gen.WithLowerLetters(c.Lower)
	}
	if c.Upper != "" {
		gen.WithUpperLetters(c.Upper)
	}
	if c.Digits != "" {
		gen.WithDigits(c.Digits)
	}
	if c.Symbols != "" {
		gen.WithSymbols(c.Symbols)
	}
	pass, err := gen.Generate(password.GenerateParams{
		Length:          c.Length,
		MinLettersCount: c.MinLettersCount,
		IncludeLower:    !c.ExcludeLower,
		IncludeUpper:    !c.ExcludeUpper,
		IncludeDigits:   !c.ExcludeDigits,
		IncludeSymbols:  !c.ExcludeSymbols,
	})
	if err != nil {
		return err
	}
	if _, err = io.WriteString(ctx.Stdout, pass); err != nil {
		return err
	}
	if !c.NoTrailingNewline {
		_, err = ctx.Stdout.Write([]byte{'\n'})
	}
	return err
}
