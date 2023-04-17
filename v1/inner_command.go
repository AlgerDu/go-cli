package cli

type innerCommand struct {
	*CommandMeta

	Action       CommandAction
	DefaultFlags any
	Children     map[string]*innerCommand
}

func NewInnerCommand(cmd Command) *innerCommand {
	return &innerCommand{
		CommandMeta:  cmd.GetDescripton(),
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
