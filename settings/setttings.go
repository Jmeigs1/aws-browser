package settings

import (
	"aws-browser/components/pointer"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func toggleNav(w fyne.Window) func() {
	return func() {
		dialog.ShowInformation("Information", "Navbar toggle - not implmented", w)
	}
}

func selectProfile(w fyne.Window) func() {
	return func() {
		dialog.ShowInformation("Information", "Navbar toggle - not implmented", w)
	}
}

func Render(w fyne.Window) *fyne.Container {
	toolbar := widget.NewToolbar(
		pointer.NewToolbarButton(theme.ListIcon(), toggleNav(w)),
		widget.NewToolbarSpacer(),
		pointer.NewToolbarButton(theme.SettingsIcon(), selectProfile(w)),
	)
	ctr := container.New(layout.NewVBoxLayout(), toolbar)
	return ctr
}
