package command

import (
	`fmt`
)

type Tv struct {
	Running bool
}

func (t *Tv) On() {
	t.Running = true
	fmt.Println("Turning tv on")
}

func (t *Tv) Off() {
	t.Running = false
	fmt.Println("Turning tv off")
}
