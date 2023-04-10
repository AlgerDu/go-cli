package cli

type defaultBuilder struct {
	cmds []Command
}

func NewBuilder() AppBuilder {
	return &defaultBuilder{
		cmds: []Command{},
	}
}

func (builder *defaultBuilder) AddCommand(command Command) {

}

func (*defaultBuilder) Build() App {
	return nil
}
