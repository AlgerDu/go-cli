package commands

import (
	"fmt"

	"github.com/AlgerDu/go-cli/v1"
)

type (
	HelpCommandFlags struct {
		Path string
	}

	HelpCommand struct {
		*cli.BaseCommand

		commands map[string]cli.Command
		appMeta  *cli.AppMeta
	}
)

func NewHelp(
	appMeta *cli.AppMeta,
	commands map[string]cli.Command,
) *HelpCommand {

	return &HelpCommand{
		BaseCommand: &cli.BaseCommand{
			DefaultFlags: &HelpCommandFlags{},
			Meta: &cli.CommandMeta{
				Name: "help",
			},
		},
		commands: commands,
		appMeta:  appMeta,
	}
}

func (cmd *HelpCommand) Action(c *cli.Context, flags any) error {

	f := flags.(HelpCommandFlags)

	fmt.Println(f)

	return nil
}
