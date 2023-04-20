package cli

import "os"

type Context struct {
	WorkDir string

	CommandPaths []string
	Flags        map[string]string
}

func newContext() *Context {

	wd, _ := os.Getwd()

	return &Context{
		WorkDir: wd,

		CommandPaths: []string{},
		Flags:        map[string]string{},
	}
}
