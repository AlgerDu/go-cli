package cli

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

	cmd.outputVersion(c.Stdout)

	toDescriptCmd, exist := cmd.app.findCmd(c)
	if !exist {
		cmd.outputUnsupportCmd(c.Stdout, c.CommandPaths)
		return nil
	}

	cmd.outputCmd(c, toDescriptCmd)

	return nil
}

func (cmd *HelpCommand) outputVersion(stdout Stdout) {

	stdout.
		Println("VERSION:").
		Scope(DefaultScopeWord).
		Println(cmd.app.Version).
		NewLline()
}

func (cmd *HelpCommand) outputUnsupportCmd(stdout Stdout, paths []string) {

	stdout.
		Println("ERROR:").
		Scope(DefaultScopeWord).
		Printfln("unsupport cmd %v", paths).
		NewLline()
}

func (cmd *HelpCommand) outputCmd(c *Context, toDescriptCmd *innerCommand) {

	c.Stdout.
		Println("NAME:").
		Scope(DefaultScopeWord).
		Println(c.CommandPaths).
		NewLline()

	c.Stdout.
		Println("USAGE:").
		Scope(DefaultScopeWord).
		Println(toDescriptCmd.Usage).
		NewLline()
}

func (cmd *HelpCommand) outputSubCmds(c *Context) error {
	return nil
}

func (cmd *HelpCommand) outputCmdFlags(c *Context, toDescriptCmd *innerCommand) error {
	return nil
}
