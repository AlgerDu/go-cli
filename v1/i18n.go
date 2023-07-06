package cli

import (
	"bytes"
	"fmt"
	"text/template"
)

type (
	I18nTag string
)

var (
	I18n = map[I18nTag]string{}

	tag_HelpForCmd I18nTag = "help_for_cmd"
)

func UseEN() {
	I18n[tag_HelpForCmd] = template_en_HelpForCmd
}

func AnaylseTemplate(tag I18nTag, data any) string {

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
