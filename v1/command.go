package cli

type CommandAction func(c *Context, flags any) error

type CommandDescription struct {
	Name    string
	Usage   string
	Aliases []string
}

func (cmd *CommandDescription) GetDescripton() *CommandDescription {
	return cmd
}
