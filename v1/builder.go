package cli

type defaultBuilder struct {
	app *innerApp
}

func NewBuilder() AppBuilder {
	return &defaultBuilder{
		app: &innerApp{
			Cmds: map[string]*innerCommandMeta{},
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

func (builder *defaultBuilder) SetAction(action CommandAction) AppBuilder {
	return builder
}

func (builder *defaultBuilder) AddCommand(command Command, opt ...AddCommandOption) AppBuilder {

	return builder
}

func (builder *defaultBuilder) Build() App {
	return builder.app
}
