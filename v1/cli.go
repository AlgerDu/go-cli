package cli

type CommandMeta struct {
	Name    string
	Usage   string
	Aliases []string
}

type CommandAction func(*Context, any) error

type Command interface {
	GetDescripton() *CommandMeta
	GetDefaultFlags() any
	Action(c *Context, flags any) error
}

type AppMeta struct {
	Usage   string
	Version string
}

type AddCommandOption func(CommandSettor)

type CommandSettor interface {
	AddSucCommand(Command, ...AddCommandOption)
}

type App interface {
	Run(args []string) error
}

type AppBuilder interface {
	SetVersion(version string) AppBuilder
	SetUsage(usage string) AppBuilder
	SetAction(action CommandAction) AppBuilder
	AddCommand(Command, ...AddCommandOption) AppBuilder
	Build() App
}
