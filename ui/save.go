package ui

import (
	"sat_word_list/internal"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func SaveUI(wd fyne.Window) func() {
	return func() {
		if err := internal.SaveWords(); err != nil {
			dialog.ShowError(err, wd)
			return
		}

		dialog.ShowInformation("Successed", "Successed to save words!", wd)
	}

}
