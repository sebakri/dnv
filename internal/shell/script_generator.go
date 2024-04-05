package shell

import "github.com/sebakri/dnv/internal/pwsh"

type ScriptGenerator interface {
	AddEnvironmentVariable(name string, value string) []byte
	RemoveEnvironmentVariable(name string) []byte
	AddToPath(path string) []byte
	RemoveFromPath(path string) []byte
	SaveScript(script []byte, path string) error
	AppendToScript(script []byte, path string) error
	PrependToScript(script []byte, path string) error
	ScriptExists(path string) bool
	AddComment(comment string) []byte
	ScriptExtension() string
}

func GetScriptGenerator(shellId string) ScriptGenerator {
	switch shellId {
	case "pwsh":
		return pwsh.NewScriptGenerator()
	default:
		return nil
	}

}
