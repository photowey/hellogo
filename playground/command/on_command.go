package command

type OnCommand struct {
	Device
}

func (c *OnCommand) Execute() error {
	c.Device.On()

	return nil
}
