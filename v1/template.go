package cli

import (
	"bytes"
	"fmt"
	"text/template"
)

func AnalyseTemplate(tag I18nTag, data any) string {

	templateStr, exist := I18n[tag]
	if !exist {
		return fmt.Sprintf("tag [%s] is not exist", tag)
	}

	t := template.New("")
	t, err := t.Parse(templateStr)
	if err != nil {
		return err.Error()
	}

	b := &bytes.Buffer{}
	err = t.Execute(b, data)
	if err != nil {
		return err.Error()
	}

	return b.String()
}

type (
	TempData_Meta struct {
		Name    string
		Usage   string
		Default string
	}

	TempData_OutputCmdHelp struct {
		Description       string
		CmdPath           string
		Flags             []*TempData_Meta
		SupportGlobalFlag bool
		GlobalFlags       []*TempData_Meta
	}

	TempData_OutputSubCmdHelp struct {
		Description       string
		CmdPath           string
		SubCommands       []*TempData_Meta
		SupportGlobalFlag bool
		GlobalFlags       []*TempData_Meta
	}
)

var (
	template_en_OutputCmdHelp = `{{ .Description }}

Usage: {{ .CmdPath }} COMMAND [OPTIONS]

Flags: 
{{- range .Flags }}
  {{ .Name }}{{ .Default }}{{ .Usage }}
{{- end }}

{{- if .SupportGlobalFlag }}

Global Flags:
{{- range .GlobalFlags }}
  {{ .Name }}{{ .Usage }}
{{- end }}
{{- end }}
`

	template_en_OutputSubCmdHelp = `{{ .Description }}

Usage: {{ .CmdPath }} COMMAND [OPTIONS]

Commands: 
{{- range .SubCommands }}
    {{ .Name }}{{ .Usage }}
{{- end }}

{{- if .SupportGlobalFlag }}
Global Flags:
{{- range .GlobalFlags }}
    {{ .Name }}{{ .Usage }}
{{- end }}
{{- end }}
`
)
