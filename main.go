package main

import (
	_ "embed"
	"os"

	"github.com/sebakri/dnv/internal/cmd"
	"github.com/sebakri/dnv/internal/log"
	"github.com/urfave/cli/v2"
)

const (
	command = "dnv"
)

func main() {
	app := &cli.App{
		Name: command,
		Commands: []*cli.Command{
			cmd.InitCommand(),
			cmd.GenerateCommand(),
			cmd.StatusCommand(),
		},
		Before: func(c *cli.Context) error {
			debug, exists := os.LookupEnv("DNV_DEBUG")
			if exists {
				log.SetDebugEnabled(debug == "true")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
