package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildStatusBar() *fyne.Container {
	// Placeholder status labels
	fps := widget.NewLabel("FPS: 0")
	cameraPos := widget.NewLabel("Camera: (0, 0, 0)")

	// Update status dynamically in the future
	return container.NewHBox(fps, cameraPos)
}
