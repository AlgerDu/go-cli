package cli

import (
	"fmt"
	"strings"
)

type innerApp struct {
	*innerCommand

	Version     string
	GlobalFlags []*flag
}

func (app *innerApp) Run(args []string) error {

	fmt.Printf("args: %v\n", args)

	context := newContext()
	context.CommandPaths, context.UserSetFlags = app.anaylseArgs(args)

	cmdPaths := context.CommandPaths

	if app.isHelp(context) {
		cmdPaths = []string{"help"}
	}

	cmd, exist := app.findCmd(cmdPaths...)
	if !exist {
		return nil
	}

	return cmd.Action(context, nil)
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

	for flag := range context.UserSetFlags {
		value, exist := helpCmdFlags[flag]
		if exist && value {
			return true
		}
	}

	return false
}

func (app *innerApp) findCmd(cmdPaths ...string) (*innerCommand, bool) {

	cmd := app.innerCommand

	if len(cmdPaths) == 0 {
		return cmd, true
	}

	for _, path := range cmdPaths {

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
