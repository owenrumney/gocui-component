package main

import (
	"github.com/awesome-gocui/gocui"
)

func main() {

	gui, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		panic(err)
	}

	defer gui.Close()

	settings := NewSettingsWidget("settings", 10, 10, 40, 10, gui)
	settings.Draw()

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}
