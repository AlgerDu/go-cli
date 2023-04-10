package main

import (
	"fmt"

	"github.com/AlgerDu/go-cli/v1"
)

type (
	HelloCommandFlags struct {
		Name string `cli:"name"`
	}

	HelloCommand struct {
		*cli.BaseCommand
	}
)

var (
	defaultHelloCommandFlags = &HelloCommandFlags{
		Name: "",
	}

	helloCommand = &HelloCommand{
		BaseCommand: &cli.BaseCommand{
			DefaultFlags: defaultHelloCommandFlags,
			Meta:         nil,
		},
	}
)

func (cmd *HelloCommand) Action(c *cli.Context, flags any) error {
	f := flags.(HelloCommandFlags)
	fmt.Printf("hello %s", f.Name)
	return nil
}
