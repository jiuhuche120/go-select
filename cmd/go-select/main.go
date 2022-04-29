package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-select"
	app.Usage = "A tool to quickly switch between Golang compiler versions."
	app.Compiled = time.Now()

	app.Commands = []*cli.Command{
		installCMD(),
		useCMD(),
		versionsCMD(),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
