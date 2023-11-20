package cli

type (
	CommandMeta struct {
		Name    string
		Usage   string
		Aliases []string
	}

	PipelineAction func(*Context) error

	Command interface {
		GetDescripton() *CommandMeta
		GetDefaultFlags() any
		Action(c *Context) error
	}

	AddCommandOption func(CommandSettor)

	CommandSettor interface {
		AddSucCommand(Command, ...AddCommandOption)
	}

	App interface {
		Run(args []string) error
	}

	AppBuilder interface {
		SetVersion(version string) AppBuilder
		SetUsage(usage string) AppBuilder
		AddCommand(Command, ...AddCommandOption) AppBuilder
		Build() App
	}
)
