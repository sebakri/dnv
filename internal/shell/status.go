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
	Env EnvStatus
}

func (s Status) Short() string {
	return fmt.Sprintf("+%d <>%d", len(s.Env.Added), len(s.Env.Replaced))
}

func (s Status) String() string {
	var status []string

	for _, added := range s.Env.Added {
		status = append(status, fmt.Sprintf("+%s", added))
	}

	for _, replaced := range s.Env.Replaced {
		status = append(status, fmt.Sprintf("<>%s", replaced))
	}

	return strings.Join(status, " ")
}

func UpdateStatus(newStatus Status) {
	statusFile := filepath.Join(env.GetDNV().SessionFolder, "status")

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
