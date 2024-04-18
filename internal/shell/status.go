package shell

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sebakri/dnv/internal/env"
	"github.com/sebakri/dnv/internal/log"
)

type EnvStatus struct {
	Added    []string
	Replaced []string
}

type Status struct {
	Env *Environment
}

func NewStatusFromEnvironment(env *Environment) Status {
	return Status{
		Env: env,
	}
}

func (s Status) Short() string {
	return fmt.Sprintf("+%d", len(s.Env.Variables))
}

func (s Status) String() string {
	var status []string

	for key := range s.Env.Variables {
		status = append(status, fmt.Sprintf("+%s", key))
	}

	return strings.Join(status, " ")
}

func UpdateStatus(newEnv *Environment) {
	statusFile := filepath.Join(env.GetDNV().SessionFolder, "status")

	newStatus := Status{
		Env: newEnv,
	}

	statusJson, err := json.Marshal(newStatus)

	if err != nil {
		log.Debug("Error marshalling status: ", err)
		return
	}

	os.WriteFile(statusFile, statusJson, 0644)
}

func CurrentStatus() (*Status, error) {
	statusFile := filepath.Join(env.GetDNV().SessionFolder, "status")
	content, err := os.ReadFile(statusFile)
	if err != nil {
		return nil, err
	}

	var status Status
	if err := json.Unmarshal(content, &status); err != nil {
		return nil, err
	}

	return &status, nil
}
