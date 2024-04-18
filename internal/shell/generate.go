package shell

func GenerateScripts(sg ScriptGenerator, env *Environment) error {
	loadScript := new(Script)
	unloadScript := new(Script)

	for key, value := range env.Variables {
		loadScript.AddLine(sg.AddEnvironmentVariable(key, value.New))
		unloadScript.AddLine(sg.AddEnvironmentVariable(key, value.Old))
	}

	sg.SaveScript(loadScript.Content, "load")

	UpdateStatus(env)

	if sg.ScriptExists("unload") {
		sg.PrependToScript(unloadScript.Content, "unload")
	} else {
		sg.SaveScript(unloadScript.Content, "unload")
	}

	return nil
}
