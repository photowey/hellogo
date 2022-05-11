package command

import (
	`testing`
)

func TestCommand(t *testing.T) {
	tv := &Tv{}

	onCommand := &OnCommand{
		Device: tv,
	}

	offCommand := &OffCommand{
		Device: tv,
	}

	onButton := &Button{
		Command: onCommand,
	}
	_ = onButton.Press()

	offButton := &Button{
		Command: offCommand,
	}
	_ = offButton.Press()
}
