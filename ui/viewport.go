package ui

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func BuildViewport() *fyne.Container {
	// A placeholder white viewport
	viewport := canvas.NewRectangle(&color.RGBA{255, 255, 255, 255})
	return container.New(layout.NewMaxLayout(), viewport)
}
