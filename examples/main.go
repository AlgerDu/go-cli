package main

import (
	"fmt"
	"os"

	"github.com/AlgerDu/go-cli/v1"
)

func main() {

	builder := cli.NewBuilder("hello").
		SetUsage("say hello to someone").
		AddCommand(helloCommand)

	app := builder.Build()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Print(err)
	}
}
