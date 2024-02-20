package pointer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ToolbarButton struct {
	Icon        fyne.Resource
	OnActivated func()
}

func (t *ToolbarButton) ToolbarObject() fyne.CanvasObject {
	button := NewPointerButton("", t.Icon, t.OnActivated)
	button.Importance = widget.LowImportance

	return button
}

func NewToolbarButton(icon fyne.Resource, onActivated func()) *ToolbarButton {
	return &ToolbarButton{icon, onActivated}
}
