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

type Variables struct {
	Added    map[string]string
	Replaced map[string]ReplaceValue
}

type Environment struct {
	Variables Variables
}

func CreateEnvironment(cfgs []*config.Config) *Environment {
	env := &Environment{
		Variables: Variables{
			Added:    make(map[string]string),
			Replaced: make(map[string]ReplaceValue),
		},
	}

	for _, cfg := range cfgs {
		for key, value := range cfg.EnvironmentVariables {
			ev, exists := os.LookupEnv(key)

			if exists {
				log.Debug("Found existing environment variable: ", key)
				if _, ok := env.Variables.Replaced[key]; !ok {
					log.Debug("Replacing environment variable: ", key)
					env.Variables.Replaced[key] = ReplaceValue{Old: ev, New: value}
				} else {
					log.Debug("Environment variable already replaced: ", key)
					oldValue := env.Variables.Replaced[key].Old
					env.Variables.Replaced[key] = ReplaceValue{Old: oldValue, New: value}
				}
			} else {
				log.Debug("Adding environment variable: ", key)
				env.Variables.Added[key] = value
			}
		}
	}

	return env
}
