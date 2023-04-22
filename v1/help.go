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

	c.Stdout.
		Println(cmd.app.Name).
		NewLline().
		Println("usage:").
		Println(cmd.app.Usage)

	return nil
}
