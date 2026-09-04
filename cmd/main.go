package main

// entry point of the program
// internal is everything except ui
// ui contains all components

import (
	"sat_word_list/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	wd := a.NewWindow("SAT Word List Data Manager")

	// make window fixed and a nice ratio
	wd.Resize(fyne.NewSize(900, 500))
	wd.SetFixedSize(true)

	wd.SetContent(ui.ListWords("words.dat", wd))

	// run the program and block the process
	wd.ShowAndRun()
}
