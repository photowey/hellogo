package adapter

import (
	"fmt"
)

type Win struct{}

func NewWin() *Win {
	return &Win{}
}

func (w *Win) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
