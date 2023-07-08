package cli

import (
	"testing"
)

func TestI18n_HelpForCmd(t *testing.T) {

	UseEN()

	data := &TempData_OutputCmdHelp{
		Description: "abc",
		CmdPath:     "hello",
		Flags: []*TempData_Meta{
			{
				Name:  "student",
				Usage: "hhhhhh",
			},
		},
		SupportGlobalFlag: true,
		GlobalFlags: []*TempData_Meta{
			{
				Name:  "-h",
				Usage: "help",
			},
		},
	}

	value := AnalyseTemplate(tag_OutputCmdHelp, data)
	t.Log(value)
}
