package cli

import "fmt"

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

	fmt.Printf("%s\n", cmd.app.Name)
	fmt.Printf("\n")
	fmt.Printf("usage:\n")
	fmt.Printf("%s", cmd.app.Usage)

	return nil
}
