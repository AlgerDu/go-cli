package cli

var (
	template_en_HelpForCmd = `
{{ .Description }}

Usage: {{ .CmdPath }} COMMAND [OPTIONS]

Commands: 
{{- range .SubCommands }}
    {{ .Name }}{{ .Default }}{{ .Usage }}
{{- end }}

{{- if .SupportGlobalFlag }}
Global Flags:
{{ - range .GlobalFlags }}
    {{ .Name }} {{ .Usgae }}
{{- end }}
{{- end }}
`
)

type TempData_HelpForCmd struct {
	Description string
	CmdPath     string
	SubCommands []struct {
		Name    string
		Default string
		Usage   string
	}
	SupportGlobalFlag bool
	GlobalFlags       []struct {
		Name  string
		Usage string
	}
}
