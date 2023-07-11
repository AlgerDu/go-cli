package cli

type innerCommand struct {
	*CommandMeta

	Action       PipelineAction
	DefaultFlags any
	Children     map[string]*innerCommand
}

func NewInnerCommand(cmd Command) *innerCommand {

	meta := cmd.GetDescripton()
	if meta == nil {
		meta = &CommandMeta{
			Name: "__APP__",
		}
	}

	return &innerCommand{
		CommandMeta:  meta,
		Action:       cmd.Action,
		DefaultFlags: cmd.GetDefaultFlags(),
		Children:     map[string]*innerCommand{},
	}
}

func (command *innerCommand) AddSucCommand(
	subCommand Command,
	options ...AddCommandOption,
) {

	innerSubCommand := NewInnerCommand(subCommand)
	command.Children[innerSubCommand.Name] = innerSubCommand

	if len(options) > 0 {
		for _, optionAction := range options {
			optionAction(innerSubCommand)
		}
	}
}

func (cmd *innerCommand) Check(path string) bool {

	if cmd.Name == path {
		return true
	}

	for _, alias := range cmd.Aliases {
		if alias == path {
			return true
		}
	}

	return false
}
