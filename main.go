package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/go-save-eyes/app"
	"github.com/go-save-eyes/window"
)

func main() {
	mainApp := app.NewMainApp()

	mW := window.CreateWindow(mainApp.MainApp, "Go Save Eyes", 400, 150)

	timeLeftLabel := widget.NewLabel("Time to break: 1h0m0s")
	timeLeftLabel.Alignment = fyne.TextAlignCenter

	progress := widget.NewProgressBar()

	go mainApp.BreakInitializer(timeLeftLabel, progress) // need to be in routine

	mW.AddObjects(timeLeftLabel, progress)
	mW.Run()
}
