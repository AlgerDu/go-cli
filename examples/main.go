package main

import (
	"os"

	"github.com/AlgerDu/go-cli/v1"
)

func main() {

	app := cli.New("1.0", "test")
	app.AddCommand(&HelloCommand{}, &HelloCommandFlags{})

	app.Run(os.Args)
}

type HelloCommandFlags struct {
	Name string `cli:"name"`
}

type HelloCommand struct {
	*cli.CommandDescription
}

func (cmd *HelloCommand) Action(c *cli.Context, flags any) error {
	return nil
}
