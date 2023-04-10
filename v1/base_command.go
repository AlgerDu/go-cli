package cli

type BaseCommand struct {
	DefaultFlags any
	Meta         *CommandMeta
}

func (cmd *BaseCommand) GetDescripton() *CommandMeta {
	return cmd.Meta
}

func (cmd *BaseCommand) GetDefaultFlags() any {
	return cmd.DefaultFlags
}
