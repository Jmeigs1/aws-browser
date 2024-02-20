package pointer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// PointerButton is a button with pointer mouse coursor when enabled.
type PointerButton struct {
	widget.Button
}

// NewPointerButton creates a new button widget with the set label and tap
// handler.
func NewPointerButton(text string, icon fyne.Resource, onTapped func()) *PointerButton {
	btn := &PointerButton{}
	btn.ExtendBaseWidget(btn)
	btn.Text = text
	btn.OnTapped = onTapped
	btn.Icon = icon
	return btn
}

// Cursor returns the cursor type of this widget.
func (b *PointerButton) Cursor() desktop.Cursor {
	if !b.Disabled() {
		return desktop.PointerCursor
	}
	return desktop.DefaultCursor
}
