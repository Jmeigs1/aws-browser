package main

import (
	appconfig "aws-browser/services/app-config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	appconfig.InitAWS()

	a := app.New()
	w := a.NewWindow("App Config Viewer")
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.SetMaster()
	w.Resize(fyne.NewSize(600, 530))

	w.SetContent(appconfig.Start())
	w.ShowAndRun()
}
