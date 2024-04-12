package pwsh

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sebakri/dnv/internal/shell"
)

type ScriptGenerator struct {
	ctx shell.GenerateContext
}

func NewScriptGenerator(ctx shell.GenerateContext) *ScriptGenerator {
	return &ScriptGenerator{
		ctx: ctx,
	}
}

func (sg *ScriptGenerator) Ctx() shell.GenerateContext {
	return sg.ctx
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

func (sg ScriptGenerator) SaveScript(script []byte, name string) error {
	if err := os.WriteFile(sg.scriptPath(name), script, 0644); err != nil {
		return err
	}
	return nil
}

func (sg *ScriptGenerator) AppendToScript(script []byte, name string) error {
	f, err := os.OpenFile(sg.scriptPath(name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	if _, err := os.Stat(sg.scriptPath(name)); os.IsNotExist(err) {
		return false
	}

	return true
}

func (sg *ScriptGenerator) AddComment(comment string) []byte {
	return []byte("# " + comment)
}

func (sg *ScriptGenerator) PrependToScript(script []byte, name string) error {
	content, err := os.ReadFile(sg.scriptPath(name))
	if err != nil {
		return err
	}

	if err := os.WriteFile(sg.scriptPath(name), append(script, content...), 0644); err != nil {
		return err
	}
	return nil
}

func (sg *ScriptGenerator) ScriptExtension() string {
	return ".ps1"
}

func (sg ScriptGenerator) scriptPath(name string) string {
	return filepath.Join(
		sg.Ctx().SessionFolder,
		strings.Join([]string{sg.Ctx().ShellID, sg.Ctx().SessionID, name + sg.ScriptExtension()}, "-"))
}
