package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildSidebar() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Explorer"),
		widget.NewList(
			func() int { return 5 }, // Placeholder items
			func() fyne.CanvasObject { return widget.NewLabel("Item") },
			func(id widget.ListItemID, obj fyne.CanvasObject) {
				obj.(*widget.Label).SetText("Item " + string(rune(id)))
			},
		),
	)
}
