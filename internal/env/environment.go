package env

import (
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

type Environment struct {
	ID                   string
	ConfigPath           string
	EnvironmentVariables map[string]string
}

func LookupEnvironmentRcs(path string) []string {
	const envrc = ".envrc"
	var rcs []string

	for {
		_, err := os.Stat(path + "/" + envrc)
		if err != nil {
			rcs = append(rcs, filepath.Join(path, envrc))
		}

		path = filepath.Dir(path)

		if path == "/" {
			break
		}
	}

	return rcs
}

func createEnvironment(path string) *Environment {
	ID := createHash()

	backupEnvironment := &Environment{
		ID:                   ID,
		ConfigPath:           path,
		EnvironmentVariables: make(map[string]string),
	}

	for _, e := range os.Environ() {
		kv := strings.Split(e, "=")
		backupEnvironment.EnvironmentVariables[kv[0]] = kv[1]
	}

	return backupEnvironment
}

func createHash() string {
	possibleChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	hash := ""
	for i := 0; i < 16; i++ {
		hash += string(possibleChars[rand.Intn(len(possibleChars))])
	}
	return hash
}
