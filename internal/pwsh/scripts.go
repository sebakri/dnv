package pwsh

import (
	"bytes"
	_ "embed"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"

	"github.com/sebakri/dnv/internal/log"
)

//go:embed templates/init.ps1
var initTemplate string

// Hook returns the PowerShell hook template.
func InitScript() string {
	var out bytes.Buffer

	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	template.Must(template.New("initTemplate").Funcs(sprig.FuncMap()).Parse(initTemplate)).Execute(&out, struct {
		UnloadCommand string
		LoadCommand   string
		CleanCommand  string
		Debug         bool
	}{
		UnloadCommand: exe + " unload",
		LoadCommand:   exe + " load",
		CleanCommand:  exe + " clean",
		Debug:         log.DebugEnabled(),
	})

	return out.String()
}
