package cmd

import (
	"os"

	"github.com/sebakri/dnv/internal/env"
	"github.com/urfave/cli/v2"
)

func CleanCommand() *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the current session",
		Hidden:  true,
		Action: func(c *cli.Context) error {
			cleanSessionFolder()
			return nil
		},
	}
}

func cleanSessionFolder() {
	sessionFolder := env.GetDNV().SessionFolder

	if _, err := os.Stat(sessionFolder); os.IsNotExist(err) {
		return
	}

	os.RemoveAll(sessionFolder)
}
