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
		cmd.outputCmd(c, flags, toDescriptCmd)
	}

	return nil
}

func (cmd *HelpCommand) outputUnsupportCmd(stdout Stdout, paths []string) {
}

func (cmd *HelpCommand) outputCmd(c *Context, data *helpCommandFlag, toDescriptCmd *innerCommand) {

	flags := c.anaylseCmdSupportFlags(*toDescriptCmd)

	tempData := &TempData_OutputCmdHelp{
		Description:       cmd.app.Usage,
		CmdPath:           strings.Join(data.CmdPaths, " "),
		Flags:             []*TempData_Meta{},
		SupportGlobalFlag: false,
		GlobalFlags:       []*TempData_Meta{},
	}

	tempData.Flags = cmd.fmtFlags(flags)

	globalFlags := []*Flag{}
	for _, gf := range cmd.app.GlobalFlags {
		globalFlags = append(globalFlags, gf.Flag)
	}
	tempData.GlobalFlags = cmd.fmtFlags(globalFlags)

	if len(tempData.GlobalFlags) > 0 {
		tempData.SupportGlobalFlag = true
	}

	value := AnalyseTemplate(Tag_OutputCmdHelp, tempData)
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
			Name:    cmd.fmtFlagKeys(flag),
			Default: fmt.Sprintf("%v", flag.Default),
			Usage:   flag.Usage,
		}

		if flag.Require {
			meta.Default = ""
		}

		if len(meta.Default) > 0 {
			meta.Default = fmt.Sprintf("[%s:%s]", defaultStr, meta.Default)
		}

		fNameMax = Ext_Max(fNameMax, len(meta.Name))
		fDefaultMax = Ext_Max(fDefaultMax, len(meta.Default))
		fUsageMax = Ext_Max(fUsageMax, len(meta.Usage))

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

func (cmd *HelpCommand) fmtFlagKeys(flag *Flag) string {

	keys := []string{flag.Name}
	keys = append(keys, flag.Aliases...)

	for i, key := range keys {
		if len(key) > 1 {
			keys[i] = fmt.Sprintf("--%s", key)
		} else {
			keys[i] = fmt.Sprintf("-%s", key)
		}
	}

	return strings.Join(keys, ",")
}
