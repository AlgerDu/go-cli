package main

import (
	"fmt"

	"github.com/AlgerDu/go-cli/v1"
)

type (
	HelloCommandFlags struct {
		Name      string
		ClassRoom string `flag:"usage:学生的教师信息"`
	}

	HelloCommand struct {
		*cli.BaseCommand
	}
)

var (
	defaultHelloCommandFlags = &HelloCommandFlags{
		Name: "ace",
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

func (cmd *HelloCommand) Action(c *cli.Context) error {
	f := c.Value.(*HelloCommandFlags)
	fmt.Printf("hello %s", f.Name)
	return nil
}
