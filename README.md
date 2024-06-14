# passgen

A simple tool for generating passwords, can be used either as a command-line interface (CLI) or as a library.  
Uses the crypto/rand library.

## CLI

Installation:

```sh
go install -ldflags="-s -w" github.com/kogutich/passgen/cmd/passgen@latest
```

Usage:

```text
Usage: passgen --length=UINT [flags]

Tool for password generation.

Flags:
  -h, --help                   Show context-sensitive help.
  -l, --length=UINT            Password length.
  -n, --no-trailing-newline    Do not output the trailing newline.
      --min-letters-count=0    Minimum letters count (lowercase or uppercase).
      --exclude-lower          Exclude lowercase letters.
      --exclude-upper          Exclude uppercase letters.
      --exclude-digits         Exclude digits.
      --exclude-symbols        Exclude symbols.
      --lower=STRING           Lowercase letters dictionary.
      --upper=STRING           Uppercase letters dictionary.
      --digits=STRING          Digits dictionary.
      --symbols=STRING         Symbols dictionary.
```

Examples:

```sh
# all character kinds with default dictionaries
$ passgen -l 12
mp+lqdNEh6fn

# only digits
$ passgen -l 4 --exclude-lower --exclude-upper --exclude-symbols
9028

# custom dictionaries
$ passgen -l 30 --lower "abc" --upper "ABC" --digits "01" --symbols "_<>"
AbA<Cc<0>A>B<>10_<>c<c0aa0_<A1
```

## Library

Installation:

```sh
go get github.com/kogutich/passgen
```

Usage:

```go
import "github.com/kogutich/passgen/password"

...

gen := password.NewGenerator().
    WithLowerLetters("abc").
    WithUpperLetters("DEF").
    WithSymbols("_@").
    WithDigits("123")

pass, err := gen.Generate(password.GenerateParams{
    Length:          36,
    MinLettersCount: 10,
    IncludeLower:    true,
    IncludeUpper:    true,
    IncludeDigits:   true,
    IncludeSymbols:  true,
})
```
