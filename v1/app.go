package cli

type innerApp struct {
	Usage   string
	Version string

	Action CommandAction

	Cmds map[string]*innerCommandMeta
}

func (app *innerApp) Run(args []string) error {
	return nil
}
