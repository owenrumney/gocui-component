package component

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

// HeadingField struct
type HeadingField struct {
	*gocui.Gui
	label string
	ctype ComponentType
	*Position
	*Attributes
	handlers Handlers
}

func (c *HeadingField) IsFocusable() bool {
	return false
}

func (c *HeadingField) Focus() {
	// TODO implement me
	c.Gui.SetCurrentView(c.label)
}

func (c *HeadingField) UnFocus() {
	// TODO implement me
	c.Gui.SetCurrentView(c.label)
}

func (c *HeadingField) AddHandlerOnly(key Key, handler Handler) {
	c.handlers[key] = handler
}

// NewHeading new heading
func NewHeading(gui *gocui.Gui, label string, x, y int) *HeadingField {
	labelWidth := len(label)
	p := &Position{
		X: x,
		Y: y,
		W: x + labelWidth + 1,
		H: y + 2,
	}

	c := &HeadingField{
		Gui:      gui,
		label:    label,
		Position: p,
		Attributes: &Attributes{
			textColor:   gocui.ColorYellow | gocui.AttrBold,
			textBgColor: gocui.ColorDefault,
		},
		handlers: make(Handlers),
		ctype:    TypeHeading,
	}

	return c
}

// GetLabel get heading label
func (c *HeadingField) GetLabel() string {
	return c.label
}

// GetPosition get heading position
func (c *HeadingField) GetPosition() *Position {
	return c.Position
}

// GetType get component type
func (c *HeadingField) GetType() ComponentType {
	return c.ctype
}

// Draw draw label and heading
func (c *HeadingField) Draw() {
	// draw label
	if v, err := c.Gui.SetView(c.label, c.X, c.Y, c.W, c.H, 0); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}

		v.Frame = false
		v.FgColor = c.textColor
		v.BgColor = c.textBgColor
		fmt.Fprint(v, c.label)
	}

}

// Close closes the heading
func (c *HeadingField) Close() {
	views := []string{
		c.label,
	}

	for _, v := range views {
		if err := c.DeleteView(v); err != nil {
			if err != gocui.ErrUnknownView {
				panic(err)
			}
		}
	}
}
