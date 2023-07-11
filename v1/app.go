package cli

import (
	"fmt"
	"strings"
)

type innerApp struct {
	*innerCommand

	Version     string
	GlobalFlags []*GlobalFlag

	pipelines []PipelineAction
}

func newInnerApp() *innerApp {
	app := &innerApp{
		innerCommand: &innerCommand{
			CommandMeta:  &CommandMeta{},
			DefaultFlags: nil,
			Children:     map[string]*innerCommand{},
		},
		Version:     "1.0.0",
		GlobalFlags: []*GlobalFlag{},
	}

	app.pipelines = []PipelineAction{
		app.checkGlobalFlags,
		app.runCmd,
	}

	return app
}

func (app *innerApp) Run(args []string) error {

	fmt.Printf("args: %v\n", args)

	context := newContext()
	context.CommandPaths, context.UserSetFlags = app.anaylseArgs(args)

	var err error
	for _, pipelineAction := range app.pipelines {
		err = pipelineAction(context)
		if err != nil {
			break
		}
	}

	return err
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

	fmtFlags := map[string]string{}
	for key, value := range flags {
		fmtKey := strings.TrimLeft(key, "-")
		fmtFlags[fmtKey] = value
	}

	return paths, fmtFlags
}

func (app *innerApp) checkGlobalFlags(context *Context) error {

	mapFlags := map[string]*GlobalFlag{}
	for _, flag := range app.GlobalFlags {
		mapFlags[flag.Name] = flag
		if flag.Aliases != nil && len(flag.Aliases) > 0 {
			for _, aliase := range flag.Aliases {
				mapFlags[aliase] = flag
			}
		}
	}

	for key := range context.UserSetFlags {
		flag, exist := mapFlags[key]
		if exist {
			return flag.Action(context)
		}
	}
	return nil
}

func (app *innerApp) runCmd(context *Context) error {

	cmdPaths := context.CommandPaths
	cmd, exist := app.findCmd(cmdPaths...)
	if !exist {
		return nil
	}

	return cmd.Action(context)
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
