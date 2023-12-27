package cli

type defaultBuilder struct {
	app              *innerApp
	pipelineSettings *PipelineSettings
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

func (builder *defaultBuilder) SetPipeline(settings *PipelineSettings) AppBuilder {
	builder.pipelineSettings = settings
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

	builder.buildPipelines()

	return builder.app
}

func (builder *defaultBuilder) useHelp() {

	builder.app.GlobalFlags = append(builder.app.GlobalFlags, helpGloblaFlag)

	cmd := newHelp(builder.app)
	builder.AddCommand(cmd)
}

func (builder *defaultBuilder) useVersion() {

	I18n[I18nTag(cmdName_Version)] = i18n_en_version

	builder.app.GlobalFlags = append(builder.app.GlobalFlags, versionGloblaFlag)

	cmd := newVersion(builder.app)
	builder.AddCommand(cmd)
}

func (builder *defaultBuilder) buildPipelines() {
	settings := builder.pipelineSettings
	if settings == nil {
		settings = &PipelineSettings{}
	}

	settings.checkGlobalFlags = []PipelineAction{builder.app.checkGlobalFlags}
	settings.findCmd = builder.app.findCmdPipelineAction
	settings.resolveFlag = builder.app.resolveFlagStruct
	settings.runCmd = builder.app.runCmd

	pipelines := []PipelineAction{}
	pipelines = append(pipelines, settings.BeferCheckGlobalFlags...)
	pipelines = append(pipelines, settings.checkGlobalFlags...)
	pipelines = append(pipelines, settings.findCmd)
	pipelines = append(pipelines, settings.resolveFlag)
	pipelines = append(pipelines, settings.runCmd)

	builder.app.pipelines = pipelines
}
