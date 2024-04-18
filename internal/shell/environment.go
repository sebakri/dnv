package shell

import (
	"os"

	"github.com/sebakri/dnv/internal/config"
	"github.com/sebakri/dnv/internal/log"
)

type ReplaceValue struct {
	New string
	Old string
}

type Environment struct {
	LoadedFrom string
	Variables  map[string]ReplaceValue
}

func CreateEnvironment(cfgs []*config.Config) *Environment {
	env := &Environment{
		LoadedFrom: cfgs[len(cfgs)-1].File,
		Variables:  make(map[string]ReplaceValue),
	}

	for _, cfg := range cfgs {
		for key, value := range cfg.EnvironmentVariables {
			ev, exists := os.LookupEnv(key)

			if exists {
				log.Debug("Found existing environment variable: ", key)
				if _, ok := env.Variables[key]; !ok {
					log.Debug("Replacing environment variable: ", key)
					env.Variables[key] = ReplaceValue{Old: ev, New: value}
				} else {
					log.Debug("Environment variable already replaced: ", key)
					oldValue := env.Variables[key].Old
					env.Variables[key] = ReplaceValue{Old: oldValue, New: value}
				}
			} else {
				log.Debug("Adding environment variable: ", key)
				env.Variables[key] = ReplaceValue{Old: "", New: value}
			}
		}
	}

	return env
}
