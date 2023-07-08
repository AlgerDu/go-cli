package cli

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	HelpCommand struct {
		*BaseCommand

		app *innerApp
	}
)

var (
	cmdName_Help = "help"
	helpCmdFlags = map[string]bool{
		"-h":     true,
		"--help": true,
	}
)

func newHelp(
	app *innerApp,
) *HelpCommand {

	return &HelpCommand{
		BaseCommand: &BaseCommand{
			Meta: &CommandMeta{
				Name: cmdName_Help,
			},
		},
		app: app,
	}
}

func (cmd *HelpCommand) Action(c *Context, flags any) error {
	toDescriptCmd, exist := cmd.app.findCmd(c.CommandPaths...)
	if !exist {
		cmd.outputUnsupportCmd(c.Stdout, c.CommandPaths)
		return nil
	}

	if len(toDescriptCmd.Children) > 0 {
		cmd.outputSubCmds(c, toDescriptCmd)
	} else {
		cmd.outputCmd(c, toDescriptCmd)
	}

	return nil
}

func (cmd *HelpCommand) outputUnsupportCmd(stdout Stdout, paths []string) {

	stdout.
		Println("ERROR:").
		Scope(DefaultScopeWord).
		Printfln("unsupport cmd %v", paths).
		NewLline()
}

func (cmd *HelpCommand) outputCmd(c *Context, toDescriptCmd *innerCommand) {

	flags := c.anaylseCmdSupportFlags(*toDescriptCmd)

	data := &TempData_OutputCmdHelp{
		Description:       cmd.app.Usage,
		CmdPath:           strings.Join(c.CommandPaths, " "),
		Flags:             []*TempData_Meta{},
		SupportGlobalFlag: false,
		GlobalFlags:       []*TempData_Meta{},
	}

	fNameMax := 0
	fDefaultMax := 0
	fUsageMax := 0

	for _, flag := range flags {

		meta := &TempData_Meta{
			Name:    strings.Join(append([]string{flag.Name}, flag.Aliases...), ","),
			Usage:   flag.Usage,
			Default: fmt.Sprintf("%v", flag.Default),
		}

		if len(meta.Name) > fNameMax {
			fNameMax = len(meta.Name)
		}
		if len(meta.Default) > fDefaultMax {
			fDefaultMax = len(meta.Default)
		}
		if len(meta.Usage) > fUsageMax {
			fUsageMax = len(meta.Usage)
		}

		data.Flags = append(data.Flags, meta)
	}

	for _, meta := range data.Flags {

		meta.Name = fmt.Sprintf("%-"+strconv.Itoa(fNameMax+4)+"s", meta.Name)
		meta.Default = fmt.Sprintf("%-"+strconv.Itoa(fDefaultMax+8)+"s", meta.Default)
		meta.Usage = fmt.Sprintf("%-"+strconv.Itoa(fNameMax)+"s", meta.Usage)
	}

	value := AnalyseTemplate(tag_OutputCmdHelp, data)
	c.Stdout.Print(value)
}

func (cmd *HelpCommand) outputSubCmds(c *Context, toDescriptCmd *innerCommand) error {
	return nil
}
