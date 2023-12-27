package cli

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/AlgerDu/go-cli/v1/exts"
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

func (app *innerApp) anaylseArgs(args []string) ([]string, UserSetFlags) {

	args = args[1:]

	paths := []string{}
	flags := UserSetFlags{}
	count := len(args)

	for i := 0; i < count; i++ {

		word := args[i]
		if strings.HasPrefix(word, "-") {
			if i+1 < count {
				nextWord := args[i+1]
				if strings.HasPrefix(nextWord, "-") {
					flags.Set(word, "")
				} else {
					flags.Set(word, nextWord)
					i++
				}
			} else {
				flags.Set(word, "")
			}

		} else {
			paths = append(paths, word)
		}

	}

	fmtFlags := UserSetFlags{}
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

func (app *innerApp) findCmdPipelineAction(context *Context) error {

	cmd, exist := app.findCmd(context.CommandPaths...)
	if !exist {
		return fmt.Errorf("no command for %v", context.CommandPaths)
	}

	context.toRunCmd = cmd
	return nil
}

func (app *innerApp) resolveFlagStruct(context *Context) error {

	supportFlags := context.anaylseCmdSupportFlags(context.toRunCmd)
	ctxValue := exts.Reflect_New(reflect.TypeOf(context.toRunCmd.DefaultFlags).Elem()).Interface()

	for _, flag := range supportFlags {
		set := false
		userSetValaues := []string{}
		vs, exist := context.UserSetFlags[flag.Name]
		if exist {
			set = true
			userSetValaues = append(userSetValaues, vs...)
			delete(context.UserSetFlags, flag.Name)
		}
		for _, aliase := range flag.Aliases {
			vs, exist = context.UserSetFlags[aliase]
			if exist {
				set = true
				userSetValaues = append(userSetValaues, vs...)
				delete(context.UserSetFlags, aliase)
			}
		}

		if !set && flag.Require {
			return fmt.Errorf("flag %s must set value", flag.Name)
		}

		if !set && !flag.Multiple {
			userSetValaues = append(userSetValaues, fmt.Sprintf("%v", flag.Default))
		}

		if len(userSetValaues) > 1 && !flag.Multiple {
			return fmt.Errorf("flag %s not support multiple set", flag.Name)
		}

		for _, v := range userSetValaues {
			err := bindFlagsToStruct(v, flag, ctxValue)
			if err != nil {
				return err
			}
		}
	}

	if len(context.UserSetFlags) > 0 {
		return fmt.Errorf("cmd not support flag %s", context.UserSetFlags)
	}

	context.Value = ctxValue
	return nil
}

func (app *innerApp) runCmd(context *Context) error {
	return context.toRunCmd.Action(context)
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
