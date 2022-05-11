package command

type Button struct {
	Command
}

func (b *Button) Press() error {
	return b.Command.Execute()
}
