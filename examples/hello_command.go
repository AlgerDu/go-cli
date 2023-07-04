package main

import (
	"fmt"

	"github.com/AlgerDu/go-cli/v1"
)

type (
	HelloCommandFlags struct {
		Name      string `cli:"name"`
		ClassRoom string
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
			Meta: &cli.CommandMeta{
				Name:  "hello",
				Usage: "say hello to someone",
			},
			DefaultFlags: defaultHelloCommandFlags,
		},
	}
)

func (cmd *HelloCommand) Action(c *cli.Context, flags any) error {
	f := flags.(HelloCommandFlags)
	fmt.Printf("hello %s", f.Name)
	return nil
}
