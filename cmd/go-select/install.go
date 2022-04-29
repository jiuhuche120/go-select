package main

import "github.com/urfave/cli/v2"

func installCMD() *cli.Command {
	return &cli.Command{
		Name:   "install",
		Usage:  "list and install available golang versions",
		Action: install,
	}
}

func install(ctx *cli.Context) error {
	return nil
}
