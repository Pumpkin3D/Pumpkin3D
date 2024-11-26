package main

import (
	"log"
	"pumpkin3d/renderer"
	"pumpkin3d/ui"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"pumpkin3d.online/server"
)

func main() {
	// Start OpenGL renderer in a separate goroutine
	go func() {
		if err := renderer.Run(); err != nil {
			log.Fatalf("Failed to start renderer: %v", err)
		}
	}()

	// Start the HTTP server for Monaco Editor
	go func() {
		if err := server.StartServer(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Initialize application
	app := app.New()
	window := app.NewWindow("Pumpkin3D.online Game Engine")

	// Build main UI
	mainLayout := ui.BuildMainLayout(window)
	window.SetContent(container.NewBorder(ui.BuildMenu(), ui.BuildStatusBar(), ui.BuildSidebar(), ui.BuildIDE(), mainLayout))

	// Configure and launch
	window.Resize(fyne.NewSize(1400, 900))
	window.ShowAndRun()
}
