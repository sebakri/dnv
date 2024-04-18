package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sebakri/dnv/internal/config"
	"github.com/sebakri/dnv/internal/env"
	"github.com/sebakri/dnv/internal/log"
	"github.com/sebakri/dnv/internal/pwsh"
	"github.com/sebakri/dnv/internal/shell"
	"github.com/urfave/cli/v2"
)

type LoadCtx struct {
	EnvLoaded                bool
	CurrentFolderContainsEnv bool
	CurrentFolderInsideEnv   bool
}

func LoadCommand() *cli.Command {
	return &cli.Command{
		Name:    "load",
		Aliases: []string{"l"},
		Usage:   "Load a new environment",
		Action: func(c *cli.Context) error {
			sh := getShell(env.GetDNV().Shell)
			if sh == nil {
				log.Debug("Shell not supported")
				return nil
			}

			sh.Init()

			currentFolder, err := os.Getwd()
			if err != nil {
				log.Debug("Error getting current folder: ", err)
				return nil
			}

			sg := sh.ScriptGenerator()
			ctx := LoadCtx{
				EnvLoaded:                false,
				CurrentFolderContainsEnv: false,
				CurrentFolderInsideEnv:   false,
			}

			existing := sh.LoadEnvironment()

			ctx.EnvLoaded = existing != nil

			_, err = os.Stat(filepath.Join(currentFolder, ".dnv"))
			ctx.CurrentFolderContainsEnv = err == nil

			if ctx.EnvLoaded {
				ctx.CurrentFolderInsideEnv = isSubfolder(existing.LoadedFrom, currentFolder)
			}

			if ctx.CurrentFolderContainsEnv {
				env := loadEnvironment(currentFolder)
				if env != nil {
					log.Debug("Loading environment")
					sg.LoadEnvironment(env)
					if err := sh.SaveEnvironment(env); err != nil {
						log.Debug("Error saving environment: ", err)
					}
				}
			}

			// if existing != nil {
			// 	log.Debug("Unloading environment")
			// 	sg.AddComment("Unloading previous environment from " + existing.LoadedFrom)
			// 	sg.UnloadEnvironment(existing)

			// 	if currentFolder == existing.LoadedFrom || isSubfolder(existing.LoadedFrom, currentFolder) {
			// 		log.Debug(currentFolder, existing.LoadedFrom)
			// 		log.Debug("Current folder is a subfolder of loaded environment. No need to reload.")
			// 	}
			// }

			// if _, err := os.Stat(filepath.Join(currentFolder, ".dnv")); os.IsNotExist(err) {
			// 	log.Debug("No .dnv folder found")
			// }

			// env := loadEnvironment(currentFolder)
			// if env != nil {
			// 	log.Debug("Loading environment")
			// 	sg.AddComment("Loading new environment from " + env.LoadedFrom)
			// 	sg.LoadEnvironment(env)
			// } else {
			// 	log.Debug("No environment found")
			// }

			// if err := sh.SaveEnvironment(env); err != nil {
			// 	log.Debug("Error saving environment: ", err)
			// 	return nil
			// }

			log.Debugf("Ctx: %+v \n", ctx)
			log.Debugf("PreviousEnv: %+v \n", existing)
			log.Debugf("Env: %+v \n", sh.LoadEnvironment())
			log.Debug("Script: \n", sg.Script())
			fmt.Fprint(os.Stdout, sg.Script())

			return nil
		},
	}
}

func isSubfolder(parent string, child string) bool {
	return strings.HasPrefix(child, parent)
}

func loadEnvironment(directory string) *shell.Environment {
	cfgs := loadConfigs(directory)

	if len(cfgs) == 0 {
		log.Debug("No config file found")

		return nil
	}

	return shell.CreateEnvironment(cfgs)
}

func loadConfigs(directory string) []*config.Config {
	configs := config.LookupConfigs(directory)

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

	return cfgs
}

func getShell(name string) shell.Shell {
	switch name {
	case "pwsh":
		return pwsh.NewShell()
	default:
		return nil
	}
}
