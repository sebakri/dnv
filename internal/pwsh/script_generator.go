package pwsh

import (
	"os"
	"path/filepath"

	"github.com/sebakri/dnv/internal/env"
)

type ScriptGenerator struct {
}

func NewScriptGenerator() *ScriptGenerator {
	return &ScriptGenerator{}
}

func (sg *ScriptGenerator) AddEnvironmentVariable(name string, value string) []byte {
	return []byte("$env:" + name + " = \"" + value + "\"")
}

func (sg *ScriptGenerator) RemoveEnvironmentVariable(name string) []byte {
	return []byte("Remove-Item env:\\" + name)
}

func (sg *ScriptGenerator) AddToPath(path string) []byte {
	return []byte("$env:PATH += \";" + path + "\"")
}

func (sg *ScriptGenerator) RemoveFromPath(path string) []byte {
	return []byte("$env:PATH = $env:PATH -replace \";" + path + "\", \"\"")
}

func (sg *ScriptGenerator) SaveScript(script []byte, name string) error {
	sessionFolder, err := loadSessionFolder()

	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(sessionFolder, name+sg.ScriptExtension()), script, 0644); err != nil {
		return err
	}
	return nil
}

func (sg *ScriptGenerator) AppendToScript(script []byte, name string) error {
	sessionFolder, err := loadSessionFolder()

	if err != nil {
		return err
	}

	f, err := os.OpenFile(filepath.Join(sessionFolder, name+sg.ScriptExtension()), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(script); err != nil {
		return err
	}
	return nil
}

func (sg *ScriptGenerator) ScriptExists(name string) bool {
	sessionFolder := env.GetDNV().SessionFolder

	if _, err := os.Stat(filepath.Join(sessionFolder, name+sg.ScriptExtension())); os.IsNotExist(err) {
		return false
	}

	return true
}

func (sg *ScriptGenerator) AddComment(comment string) []byte {
	return []byte("# " + comment)
}

func (sg *ScriptGenerator) PrependToScript(script []byte, name string) error {
	sessionFolder, err := loadSessionFolder()

	if err != nil {
		return err
	}

	content, err := os.ReadFile(filepath.Join(sessionFolder, name+sg.ScriptExtension()))
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(sessionFolder, name+sg.ScriptExtension()), append(script, content...), 0644); err != nil {
		return err
	}
	return nil
}

func (sg *ScriptGenerator) ScriptExtension() string {
	return ".ps1"
}

func loadSessionFolder() (string, error) {
	sessionFolder := env.GetDNV().SessionFolder

	if _, err := os.Stat(sessionFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(sessionFolder, 0755); err != nil {
			return "", err
		}
	}

	return sessionFolder, nil
}
