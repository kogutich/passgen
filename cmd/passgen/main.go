package main

import "github.com/alecthomas/kong"

func main() {
	ctx := kong.Parse(
		&CLI{},
		kong.Name("passgen"),
		kong.Description("Tool for password generation."),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
	)
	if err := ctx.Run(ctx); err != nil {
		ctx.FatalIfErrorf(err)
	}
}
