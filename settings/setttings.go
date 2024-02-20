package settings

import (
	"aws-browser/components/pointer"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Render() *fyne.Container {
	toolbar := widget.NewToolbar(
		pointer.NewToolbarButton(theme.ListIcon(), func() {}),
		widget.NewToolbarSpacer(),
		pointer.NewToolbarButton(theme.SettingsIcon(), func() {}),
	)
	ctr := container.New(layout.NewVBoxLayout(), toolbar)
	return ctr
}
