package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func BuildMainLayout(window fyne.Window) *fyne.Container {
	viewport := BuildViewport()
	return container.NewMax(viewport)
}
