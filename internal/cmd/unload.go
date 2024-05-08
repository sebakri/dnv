package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sebakri/dnv/internal/env"
	"github.com/sebakri/dnv/internal/log"
	"github.com/urfave/cli/v2"
)

func UnloadCommand() *cli.Command {
	return &cli.Command{
		Name:    "unload",
		Aliases: []string{"u"},
		Usage:   "Unload the current environment",
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
				ctx.CurrentFolderInsideEnv = isSubfolder(filepath.Dir(existing.LoadedFrom), currentFolder)
			}

			if ctx.EnvLoaded && ctx.CurrentFolderContainsEnv {
				log.Debug("Existing Env + Current Folder Contains Env")
				sg.UnloadEnvironment(existing)
				cleanSessionFolder()
			}

			if ctx.EnvLoaded && !ctx.CurrentFolderContainsEnv && !ctx.CurrentFolderInsideEnv {
				log.Debug("Existing Env + Current Folder Does Not Contain Env + Current Folder Not Inside Env")
				sg.UnloadEnvironment(existing)
				cleanSessionFolder()
			}

			log.Debugf("Ctx: %+v \n", ctx)
			log.Debugf("PreviousEnv: %+v \n", existing)
			log.Debugf("Env: %+v \n", sh.LoadEnvironment())
			log.Debug("Script: \n", sg.Script())

			fmt.Fprint(os.Stdout, sg.Script())

			return nil
		},
	}
}
