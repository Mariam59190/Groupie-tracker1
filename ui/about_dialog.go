
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// ShowAboutDialog affiche une fenêtre "À propos"
func ShowAboutDialog(w fyne.Window) {
	dialog.ShowInformation(
		"À propos",
		"Groupie Tracker\n\nApplication développée en Go avec Fyne.\nProjet pédagogique.",
		w,
	)
}
