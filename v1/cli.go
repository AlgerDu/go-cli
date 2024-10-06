package cli

type (
	CommandMeta struct {
		Name    string
		Usage   string
		Aliases []string
	}

	PipelineAction func(*Context) error

	// 管道的配置
	PipelineSettings struct {
		BeferCheckGlobalFlags []PipelineAction // 可以用来做一些初始化的工作

		checkGlobalFlags []PipelineAction // 检查全局 flag
		findCmd          PipelineAction   // 查找要执行的命令
		resolveFlag      PipelineAction   // 将用户设置的 flags 解析为要执行命令的 flas 结构体，方便使用
		runCmd           PipelineAction   // 运行命令
	}

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
		SetPipeline(settings *PipelineSettings) AppBuilder
		AddCommand(Command, ...AddCommandOption) AppBuilder
		Build() App
	}
)
