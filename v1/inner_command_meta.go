package cli

type innerCommandMeta struct {
	*CommandMeta

	Action       CommandAction
	DefaultFlags any
	Children     map[string]*innerCommandMeta
}

func (command *innerCommandMeta) AddSucCommand(
	subCommand Command,
	options ...AddCommandOption,
) {

	innerSubCommand := &innerCommandMeta{
		CommandMeta:  subCommand.GetDescripton(),
		Action:       subCommand.Action,
		DefaultFlags: subCommand.GetDefaultFlags(),
		Children:     map[string]*innerCommandMeta{},
	}

	command.Children[innerSubCommand.Name] = innerSubCommand

	if len(options) > 0 {
		for _, optionAction := range options {
			optionAction(innerSubCommand)
		}
	}
}
