package adapter

import "fmt"

type windwosAdapter struct {
	windowsMachine *windows
}

func (w *windwosAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts lightning singnal to USB.")
	w.windowsMachine.insertIntoUSBPort()
}
