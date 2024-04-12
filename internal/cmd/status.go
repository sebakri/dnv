package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sebakri/dnv/internal/env"
	"github.com/sebakri/dnv/internal/log"
	"github.com/sebakri/dnv/internal/shell"
	"github.com/urfave/cli/v2"
)

func StatusCommand() *cli.Command {
	return &cli.Command{
		Name:  "status",
		Usage: "show current status",
		Action: func(c *cli.Context) error {
			if _, err := os.Stat(filepath.Join(env.GetDNV().SessionFolder, "status")); os.IsNotExist(err) {
				log.Debug("Status file not found")
				return nil
			}

			status, err := shell.CurrentStatus()
			if err != nil {
				log.Debug("Error reading status file")
				return nil
			}

			fmt.Fprintln(os.Stdout, status.String())

			return nil
		},
	}
}
