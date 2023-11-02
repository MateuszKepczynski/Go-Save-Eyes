package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type MainWindow struct {
	window  fyne.Window
	objects []fyne.CanvasObject
}

func CreateWindow(mainApp fyne.App, text string, width, height float32) MainWindow {
	mW := MainWindow{
		window: mainApp.NewWindow(text),
	}

	mW.window.Resize(fyne.Size{Width: width, Height: height})

	return mW
}

func (mw *MainWindow) AddObjects(objects ...fyne.CanvasObject) {
	mw.objects = append(mw.objects, objects...)
}

func (mw *MainWindow) Run() {
	mw.window.SetContent(container.NewVBox(mw.objects...))
	mw.window.ShowAndRun()
}
