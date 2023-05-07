package main

import (
	"os"

	"github.com/AlgerDu/go-cli/v1"
)

func main() {

	builder := cli.NewBuilder("hello").
		SetUsage("say hello to someone").
		AddCommand(helloCommand, func(cs cli.CommandSettor) {
			cs.AddSucCommand(helloCommand)
		})

	app := builder.Build()
	app.Run(os.Args)
}
