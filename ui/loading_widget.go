package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// NewLoadingWidget retourne un widget de chargement
func NewLoadingWidget() fyne.CanvasObject {
	label := widget.NewLabel("Chargement...")
	bar := widget.NewProgressBarInfinite()

	return container.NewVBox(
		label,
		bar,
	)
}
