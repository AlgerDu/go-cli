package cli

import (
	"testing"
)

func TestI18n_HelpForCmd(t *testing.T) {

	UseEN()

	data := &TempData_HelpForCmd{
		Description: "abc",
		CmdPath:     "hello",
		SubCommands: []struct {
			Name  string
			Usage string
		}{
			{
				Name:  "student",
				Usage: "hhhhhh",
			},
		},
		SupportGlobalFlag: true,
		GlobalFlags: []struct {
			Name  string
			Usage string
		}{
			{
				Name:  "-h",
				Usage: "help",
			},
		},
	}

	value := AnaylseTemplate(tag_HelpForCmd, data)
	t.Log(value)
}
