package cli

type Command interface {
	GetDescripton() *CommandDescription
	Action(c *Context, flags any) error
}

type CommandSettor interface {
	Parent(command string)
}

type App interface {
	AddCommand(command Command, defaultFlags any)
	Run(args []string)
}
