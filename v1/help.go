package cli

import (
	"fmt"
)

type (
	HelpCommandFlags struct {
		Path string
	}

	HelpCommand struct {
		*BaseCommand

		app *innerApp
	}
)

func NewHelp(
	app *innerApp,
) *HelpCommand {

	return &HelpCommand{
		BaseCommand: &BaseCommand{
			DefaultFlags: &HelpCommandFlags{},
			Meta: &CommandMeta{
				Name: "help",
			},
		},
		app: app,
	}
}

func (cmd *HelpCommand) Action(c *Context, flags any) error {

	f := flags.(HelpCommandFlags)

	fmt.Println(f)

	return nil
}
