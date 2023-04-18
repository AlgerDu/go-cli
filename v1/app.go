package cli

import "fmt"

type innerApp struct {
	Name    string
	Usage   string
	Version string

	Cmds map[string]*innerCommand
}

func (app *innerApp) Run(args []string) error {

	fmt.Printf("%v", args)

	//paths, flags := anaylseArgs(args)

	return nil
}

func (app *innerApp) isHelp(path []string, flags map[string]string) bool {

	return false
}
