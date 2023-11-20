package cli

type defaultBuilder struct {
	app *innerApp
}

func NewBuilder(name string) AppBuilder {
	return &defaultBuilder{
		app: newInnerApp(),
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

	UseEN()
	builder.useHelp()
	builder.useVersion()

	return builder.app
}

func (builder *defaultBuilder) useHelp() {

	builder.app.GlobalFlags = append(builder.app.GlobalFlags, helpGloblaFlag)

	innerHelpCmd := newHelp(builder.app)
	builder.AddCommand(innerHelpCmd)
}

func (builder *defaultBuilder) useVersion() {

	I18n[I18nTag(cmdName_Version)] = i18n_en_version

	builder.app.GlobalFlags = append(builder.app.GlobalFlags, versionGloblaFlag)

	innerHelpCmd := newVersion(builder.app)
	builder.AddCommand(innerHelpCmd)
}
