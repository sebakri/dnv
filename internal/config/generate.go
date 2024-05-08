package config

import (
	"os"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue"
	"github.com/sebakri/dnv/internal/log"
)

func Parse(content string) (*Config, error) {
	var r cue.Context

	value := r.CompileString(content)
	if value.Err() != nil {
		return nil, value.Err()
	}

	var cfg Config

	if err := value.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LookupConfigs(path string) []string {
	configs := []string{}

	path = strings.TrimRight(path, string(filepath.Separator))

	file := filepath.Join(path, ".dnv")

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return configs
	}

	for {

		log.Debug("Looking for config file in: ", path)
		file := filepath.Join(path, ".dnv")

		if isRoot(path) { // reached root
			break
		}

		path = filepath.Dir(path)

		if _, err := os.Stat(file); os.IsNotExist(err) {
			continue
		}

		log.Debug("Found config file: ", file)

		configs = append(configs, file)
	}

	return configs
}

func isRoot(path string) bool {
	return path == filepath.Dir(path)
}
