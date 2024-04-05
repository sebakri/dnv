package shell

import (
	"github.com/sebakri/dnv/internal/log"
)

func GenerateScripts(sg ScriptGenerator, env *Environment) error {
	loadScript := new(Script)
	unloadScript := new(Script)

	for key, value := range env.Variables.Added {
		loadScript.AddLine(sg.AddEnvironmentVariable(key, value))
		unloadScript.AddLine(sg.RemoveEnvironmentVariable(key))
	}

	for key, value := range env.Variables.Replaced {
		loadScript.AddLine(sg.AddEnvironmentVariable(key, value.New))
		unloadScript.AddLine(sg.AddEnvironmentVariable(key, value.Old))
	}

	status := createStatus(env)

	log.Debug(status.String())
	log.Debug(status.Short())

	sg.SaveScript(loadScript.Content, "load")

	if sg.ScriptExists("unload") {
		sg.PrependToScript(unloadScript.Content, "unload")
	} else {
		sg.SaveScript(unloadScript.Content, "unload")
	}

	return nil
}

func createStatus(env *Environment) Status {
	var status Status

	for key := range env.Variables.Added {
		status.Env.Added = append(status.Env.Added, key)
	}

	for key := range env.Variables.Replaced {
		status.Env.Replaced = append(status.Env.Replaced, key)
	}

	return status
}