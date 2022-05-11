package adapter

import (
	`fmt`
)

type WinAdapter struct {
	win *Win
}

func NewWinAdapter(win *Win) *WinAdapter {
	return &WinAdapter{
		win,
	}
}

func (w *WinAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.win.insertIntoUSBPort()
}
