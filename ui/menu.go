package ui

import (
	"fyne.io/fyne/v2/widget"
)

func BuildMenu() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(widget.NewIcon(nil), func() {
			// New project
		}),
		widget.NewToolbarAction(widget.NewIcon(nil), func() {
			// Open project
		}),
		widget.NewToolbarAction(widget.NewIcon(nil), func() {
			// Save project
		}),
	)
}
