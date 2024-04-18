package shell

type Shell interface {
	ID() string
	Name() string
	SessionFolder() string

	Init() error

	SaveEnvironment(*Environment) error
	LoadEnvironment() *Environment

	ScriptGenerator() ScriptGenerator
}
