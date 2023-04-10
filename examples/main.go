package main

import (
	"os"

	"github.com/AlgerDu/go-cli/v1"
)

func main() {

	app := cli.New("1.0", "test")
	app.AddCommand(helloCommand)

	app.Run(os.Args)
}
