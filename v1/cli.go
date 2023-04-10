package cli

type CommandMeta struct {
	Parent  string
	Name    string
	Usage   string
	Aliases []string
}

type Command interface {
	GetDescripton() *CommandMeta
	GetDefaultFlags() any
	Action(c *Context, flags any) error
}

type App interface {
	AddCommand(command Command)
	Run(args []string)
}

type AppBuilder interface {
	AddCommand(command Command)
	Build() App
}
