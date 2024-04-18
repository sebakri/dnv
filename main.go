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
			cmd.StatusCommand(),
			cmd.CleanCommand(),
			cmd.LoadCommand(),
			cmd.UnloadCommand(),
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Enable debug mode",
			},
		},
		Before: func(c *cli.Context) error {
			debug, exists := os.LookupEnv("DNV_DEBUG")
			if exists {
				log.SetDebugEnabled(debug == "true")
			}

			if c.Bool("debug") {
				log.SetDebugEnabled(true)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
