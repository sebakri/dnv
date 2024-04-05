package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sebakri/dnv/internal/log"
	"github.com/sebakri/dnv/internal/pwsh"
	"github.com/urfave/cli/v2"
)

var (
	supportedShells = []string{
		"pwsh",
	}
)

func InitCommand() *cli.Command {
	return &cli.Command{
		Name:      "init",
		Usage:     "initialize dnv for a specific shell",
		Args:      true,
		ArgsUsage: fmt.Sprintf("<%s>", strings.Join(supportedShells, "|")),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "enable debug mode",
				Value: false,
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 || c.Args().Len() > 1 {
				return cli.ShowCommandHelp(c, "init")
			}

			switch c.Args().First() {
			case "pwsh":
				fmt.Fprintf(os.Stdout, pwsh.InitScript())
			}
			return nil
		},
		Before: func(c *cli.Context) error {
			log.SetDebugEnabled(c.Bool("debug"))
			return nil
		},
	}
}
