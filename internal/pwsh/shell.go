package pwsh

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sebakri/dnv/internal/shell"
)

type sh struct {
	id            string
	name          string
	sessionFolder string
	envFile       string
}

const (
	name    = "pwsh"
	envFile = "env.json"
)

func NewShell() shell.Shell {
	id := strconv.Itoa(os.Getppid())
	return &sh{
		id:            id,
		name:          name,
		envFile:       envFile,
		sessionFolder: filepath.Join(os.TempDir(), "dnv", name+"-"+id),
	}
}

func (sh *sh) ID() string {
	return sh.id
}

func (sh *sh) Name() string {
	return sh.name
}

func (sh *sh) SessionFolder() string {
	return sh.sessionFolder
}

func (sh *sh) SaveEnvironment(env *shell.Environment) error {
	json, err := json.Marshal(env)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(sh.sessionFolder, sh.envFile), json, 0644)
}

func (sh *sh) LoadEnvironment() *shell.Environment {
	var env shell.Environment

	envFile := filepath.Join(sh.sessionFolder, sh.envFile)

	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		return nil
	}

	content, err := os.ReadFile(envFile)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(content, &env)
	if err != nil {
		return nil
	}

	return &env
}

func (sh *sh) ScriptGenerator() shell.ScriptGenerator {
	ctx := shell.GenerateContext{
		SessionFolder: sh.sessionFolder,
		ShellID:       sh.name,
		SessionID:     sh.id,
	}

	return NewScriptGenerator(ctx)
}

func (sh *sh) Init() error {
	if _, err := os.Stat(sh.sessionFolder); os.IsNotExist(err) {
		err := os.MkdirAll(sh.sessionFolder, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
