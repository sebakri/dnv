package cmd

import (
	"os"

	"github.com/sebakri/dnv/internal/config"
	"github.com/sebakri/dnv/internal/env"
	"github.com/sebakri/dnv/internal/log"
	"github.com/sebakri/dnv/internal/pwsh"
	"github.com/sebakri/dnv/internal/shell"
	"github.com/urfave/cli/v2"
)

func GenerateCommand() *cli.Command {
	return &cli.Command{
		Name:  "generate",
		Usage: "lookup and generate environment scripts",
		Action: func(c *cli.Context) error {

			var currentDirectory string

			if c.Args().Len() == 0 {
				currentDirectory, _ = os.Getwd()
			} else {
				currentDirectory = c.Args().First()
			}

			configs := config.LookupConfigs(currentDirectory)

			cfgs := make([]*config.Config, 0)

			for _, cfgPath := range configs {

				content, err := os.ReadFile(cfgPath)
				if err != nil {
					log.Debug("Error reading config: ", err)
					continue
				}

				cfg, err := config.Parse(string(content))

				cfg.File = cfgPath

				if err != nil {
					log.Debug("Error loading config: ", err)
					continue
				}

				cfgs = append([]*config.Config{cfg}, cfgs...)
			}

			if len(cfgs) == 0 {
				log.Debug("No config file found")
				return nil
			}

			gctx := shell.GenerateContext{
				SessionFolder: env.GetDNV().SessionFolder,
				ShellID:       env.GetDNV().Shell,
				SessionID:     env.GetDNV().SessionId,
			}

			sg := getScriptGenerator(env.GetDNV().Shell, gctx)

			if sg == nil {
				log.Debug("Shell not supported")
				return nil
			}

			env := shell.CreateEnvironment(cfgs)

			if err := shell.GenerateScripts(sg, env); err != nil {
				return err
			}

			return nil
		},
	}
}

func getScriptGenerator(shellId string, ctx shell.GenerateContext) shell.ScriptGenerator {
	switch shellId {
	case "pwsh":
		return pwsh.NewScriptGenerator(ctx)
	default:
		return nil
	}

}
