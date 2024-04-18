package shell

type GenerateContext struct {
	SessionFolder string
	ShellID       string
	SessionID     string
}

type ScriptGenerator interface {
	Ctx() GenerateContext
	AddEnvironmentVariable(name string, value string) []byte
	RemoveEnvironmentVariable(name string) []byte
	AddToPath(path string) []byte
	RemoveFromPath(path string) []byte
	SaveScript(script []byte, path string) error
	AppendToScript(script []byte, path string) error
	PrependToScript(script []byte, path string) error
	ScriptExists(path string) bool
	AddComment(comment string)
	ScriptExtension() string

	UnloadEnvironment(env *Environment)
	LoadEnvironment(env *Environment)
	Script() string
}

type Script struct {
	Content []byte
}

func (s *Script) AddLine(line []byte) {
	s.Content = append(s.Content, line...)
	s.Content = append(s.Content, "\n"...)
}
