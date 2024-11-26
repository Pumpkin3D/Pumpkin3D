package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/webview"
)

func BuildIDE() *fyne.Container {
	// Embed a WebView for Monaco Editor
	editor := widget.NewButton("Open IDE", func() {
		w := webview.New(true)
		defer w.Destroy()
		w.SetTitle("Pumpkin3D IDE")
		w.SetSize(800, 600, webview.HintNone)
		w.Navigate("http://localhost:8080/editor.html") // Load the Monaco Editor page
		w.Run()
	})

	return container.NewVBox(editor)
}
