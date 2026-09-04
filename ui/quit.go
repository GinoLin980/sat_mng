package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func Quit(saved bool, wd fyne.Window, app fyne.App) func() {
	customQuit := func(yes bool) {
		if yes {
			if err := save(); err != nil {
				dialog.ShowError(err, wd)
				return
			}
		}
	}

	// returning a function to match fyne.v2.widget.NewButton(string, func()) interface
	// if we didn't wrap it with func(){} then the call in button will be a nil return value and be called immediately when the UI is rendered.
	return func() {
		if !saved {
			dialog.ShowConfirm(
				"Something Not Saved", // Title
				"You didn't save the updated words, Do you want to save it.", // Describtion
				customQuit, // callback function which receive bool
				wd,         // parent window to show the dialog
			)
		}

		app.Quit()
	}
}
