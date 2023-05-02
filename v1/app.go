package cli

import (
	"fmt"
	"strings"
)

type innerApp struct {
	*innerCommand

	Version string
}

func (app *innerApp) Run(args []string) error {

	fmt.Printf("args: %v\n", args)

	context := newContext()
	context.CommandPaths, context.Flags = app.anaylseArgs(args)

	if app.isHelp(context) {

		helpCmd := newHelp(app)
		helpCmd.Action(context, nil)
	}

	return nil
}

func (app *innerApp) anaylseArgs(args []string) ([]string, map[string]string) {

	args = args[1:]

	paths := []string{}
	flags := map[string]string{}
	count := len(args)

	for i := 0; i < count; i++ {

		word := args[i]
		if strings.HasPrefix(word, "-") {
			if i+1 < count {
				nextWord := args[i+1]
				if strings.HasPrefix(nextWord, "-") {

					flags[word] = ""
				} else {

					flags[word] = nextWord
					i++
				}
			} else {
				flags[word] = ""
			}

		} else {
			paths = append(paths, word)
		}

	}

	return paths, flags
}

func (app *innerApp) isHelp(context *Context) bool {

	for _, path := range context.CommandPaths {
		if path == cmdName_Help {
			return true
		}
	}

	for flag := range context.Flags {
		value, exist := helpCmdFlags[flag]
		if exist && value {
			return true
		}
	}

	return false
}

func (app *innerApp) findCmd(c *Context) (*innerCommand, bool) {

	cmd := app.innerCommand

	if len(c.CommandPaths) == 0 {
		return cmd, true
	}

	for _, path := range c.CommandPaths {

		findChild := false

		for _, childCmd := range cmd.Children {
			if childCmd.Check(path) {
				findChild = true
				cmd = childCmd
				break
			}
		}

		if !findChild {
			return nil, false
		}
	}

	return cmd, true
}
