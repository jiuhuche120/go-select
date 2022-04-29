package main

import "github.com/urfave/cli/v2"

func useCMD() *cli.Command {
	return &cli.Command{
		Name:   "use",
		Usage:  "change the version of global golang compiler",
		Action: use,
	}
}

func use(ctx *cli.Context) error {
	return nil
}
