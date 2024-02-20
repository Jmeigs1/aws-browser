package main

import (
	appconfig "aws-browser/services/app-config"
	"aws-browser/settings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	appconfig.InitAWS()

	a := app.New()
	w := a.NewWindow("App Config Viewer")
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.SetMaster()
	w.Resize(fyne.NewSize(600, 530))

	layout := container.New(
		layout.NewVBoxLayout(),
		settings.Render(),
		appconfig.Render(),
	)

	w.SetContent(layout)
	w.ShowAndRun()
}
