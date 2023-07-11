package cli

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	helpCommandFlag struct {
		CmdPaths []string
	}

	HelpCommand struct {
		*BaseCommand

		app *innerApp
	}
)

var (
	cmdName_Help = "help"

	helpGloblaFlag = &GlobalFlag{
		Flag: &Flag{
			FieldName: "",
			Name:      cmdName_Help,
			Aliases:   []string{"h"},
		},
		Action: helpGlobalFlagAction,
	}
)

func helpGlobalFlagAction(context *Context) error {

	context.Value = &helpCommandFlag{
		CmdPaths: context.CommandPaths,
	}
	context.CommandPaths = []string{cmdName_Help}

	return nil
}

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

func (cmd *HelpCommand) Action(c *Context) error {

	flags := c.Value.(*helpCommandFlag)

	toDescriptCmd, exist := cmd.app.findCmd(flags.CmdPaths...)
	if !exist {
		cmd.outputUnsupportCmd(c.Stdout, flags.CmdPaths)
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

	data.Flags = cmd.fmtFlags(flags)
	//data.GlobalFlags = cmd.fmtFlags(cmd.app.GlobalFlags)

	if len(data.GlobalFlags) > 0 {
		data.SupportGlobalFlag = true
	}

	value := AnalyseTemplate(Tag_OutputCmdHelp, data)
	c.Stdout.Print(value)
}

func (cmd *HelpCommand) outputSubCmds(c *Context, toDescriptCmd *innerCommand) error {
	return nil
}

func (cmd *HelpCommand) fmtFlags(flags []*Flag) []*TempData_Meta {

	metas := []*TempData_Meta{}

	fNameMax := 0
	fDefaultMax := 0
	fUsageMax := 0
	defaultStr := I18n[Tag_Default]

	for _, flag := range flags {

		meta := &TempData_Meta{
			Name:    strings.Join(append([]string{flag.Name}, flag.Aliases...), ","),
			Default: fmt.Sprintf("%v", flag.Default),
			Usage:   flag.Usage,
		}

		if flag.Require {
			meta.Default = ""
		}

		if len(meta.Default) > 0 {
			meta.Default = fmt.Sprintf("[%s:%s]", defaultStr, meta.Default)
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

		metas = append(metas, meta)
	}

	fNameMax = fNameMax + 4
	if fDefaultMax > 0 {
		fDefaultMax = fDefaultMax + 8
	}

	for _, meta := range metas {

		meta.Name = fmt.Sprintf("%-"+strconv.Itoa(fNameMax)+"s", meta.Name)
		meta.Default = fmt.Sprintf("%-"+strconv.Itoa(fDefaultMax)+"s", meta.Default)
		meta.Usage = fmt.Sprintf("%-"+strconv.Itoa(fUsageMax)+"s", meta.Usage)
	}

	return metas
}
