package cmd

import (
	"fmt"
	"os"

	"github.com/sebakri/dnv/internal/env"
	"github.com/sebakri/dnv/internal/log"
	"github.com/sebakri/dnv/internal/shell"
	"github.com/urfave/cli/v2"
)

func StatusCommand() *cli.Command {
	return &cli.Command{
		Name:    "status",
		Aliases: []string{"s"},
		Usage:   "show current status",
		Action: func(c *cli.Context) error {
			sh := getShell(env.GetDNV().Shell)
			if sh == nil {
				log.Debug("Shell not supported")
				return nil
			}

			existing := sh.LoadEnvironment()

			if existing == nil {
				fmt.Fprintln(os.Stdout, "No environment loaded")
				return nil
			}

			status := shell.NewStatusFromEnvironment(existing)

			fmt.Fprintln(os.Stdout, status.String())

			return nil
		},
	}
}
