package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildPropertiesPanel() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Properties"),
		widget.NewForm(
			widget.NewFormItem("Name", widget.NewEntry()),
			widget.NewFormItem("Position X", widget.NewEntry()),
			widget.NewFormItem("Position Y", widget.NewEntry()),
		),
	)
}
