package cli

import (
	"os"
)

type (
	Stdout interface {
		Scope(word string) Stdout

		NewLline() Stdout

		Printf(format string, a ...any) Stdout
		Print(a ...any) Stdout

		Printfln(format string, a ...any) Stdout
		Println(a ...any) Stdout
	}

	Context struct {
		WorkDir string

		CommandPaths []string
		UserSetFlags map[string]string

		Stdout Stdout
	}
)

func newContext() *Context {

	wd, _ := os.Getwd()

	return &Context{
		WorkDir: wd,

		CommandPaths: []string{},
		UserSetFlags: map[string]string{},

		Stdout: newStdout(),
	}
}

func (context *Context) anaylseFlag() {

}
