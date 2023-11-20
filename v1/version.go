package cli

type (
	versionCommand struct {
		*BaseCommand

		app *innerApp
	}
)

var (
	cmdName_Version = "version"
	i18n_en_version = "show version"

	versionGloblaFlag = &GlobalFlag{
		Flag: &Flag{
			FieldName: "",
			Name:      cmdName_Version,
			Aliases:   []string{"v"},
		},
		Action: versionGlobalFlagAction,
	}
)

func versionGlobalFlagAction(context *Context) error {

	context.CommandPaths = []string{cmdName_Version}

	return nil
}

func newVersion(
	app *innerApp,
) *versionCommand {

	return &versionCommand{
		BaseCommand: &BaseCommand{
			Meta: &CommandMeta{
				Name: cmdName_Version,
			},
		},
		app: app,
	}
}

func (cmd *versionCommand) Action(c *Context) error {

	tempData := &TempData_Version{
		Version: cmd.app.Version,
	}

	value := AnalyseTemplate(Tag_OutputVersion, tempData)
	c.Stdout.Print(value)

	return nil
}
