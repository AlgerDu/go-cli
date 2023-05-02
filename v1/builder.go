package cli

type defaultBuilder struct {
	app *innerApp
}

func NewBuilder(name string) AppBuilder {
	return &defaultBuilder{
		app: &innerApp{
			innerCommand: &innerCommand{
				CommandMeta:  &CommandMeta{},
				DefaultFlags: nil,
				Children:     map[string]*innerCommand{},
			},
			Version: "1.0.0",
		},
	}
}

func (builder *defaultBuilder) SetVersion(version string) AppBuilder {
	builder.app.Version = version
	return builder
}

func (builder *defaultBuilder) SetUsage(usage string) AppBuilder {
	builder.app.Usage = usage
	return builder
}

func (builder *defaultBuilder) AddCommand(command Command, opt ...AddCommandOption) AppBuilder {

	innerCommand := NewInnerCommand(command)

	if innerCommand.Name == "" {
		builder.app.Action = innerCommand.Action
		innerCommand = builder.app.innerCommand

		for _, option := range opt {
			option(innerCommand)
		}
	} else {
		builder.app.AddSucCommand(command, opt...)
	}

	return builder
}

func (builder *defaultBuilder) Build() App {
	return builder.app
}
