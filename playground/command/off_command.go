package command

type OffCommand struct {
	Device
}

func (c *OffCommand) Execute() error {
	c.Device.Off()

	return nil
}
