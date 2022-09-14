package main

import (
	"github.com/awesome-gocui/gocui"
	component "github.com/owenrumney/gocui-form"
)

type SettingsWidget struct {
	name string
	x, y int
	w, h int
	form *component.Form
}

func NewSettingsWidget(name string, x, y, w, h int, gui *gocui.Gui) *SettingsWidget {
	return &SettingsWidget{
		name: name,
		x:    x,
		y:    y,
		w:    w,
		h:    h,
		form: component.NewForm(gui, name, x, y, w, h),
	}
}

func (s *SettingsWidget) Draw() {
	s.form.AddHeading("General Options")
	s.form.AddCheckBox("  Enable Debugging", 25, true)
	s.form.AddInputField("  Cache Directory", 25, 25, "/home/owen/.cache/trivy")
	s.form.AddHeading("Scanning Options")
	s.form.AddCheckBox("  Scan Vulnerabilities", 25, true)
	s.form.AddCheckBox("  Scan Misconfigs", 25, true)
	s.form.AddCheckBox("  Show Secrets", 25, false)
	s.form.AddCheckBox("  Ignore Unfixed", 25, false)

	s.form.Draw()
}
