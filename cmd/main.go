package main

// entry point of the program
// internal is everything except ui
// ui contains all components

import (
	"sat_word_list/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

const Filename = "words.dat"

func main() {
	a := app.New()
	wd := a.NewWindow("SAT Word List Data Manager")
	wd.SetMaster() // closing this window quits the app, closing any other open windows (e.g. Add New Word)

	// make window fixed and a nice ratio
	wd.Resize(fyne.NewSize(900, 500))
	wd.SetFixedSize(true)

	wd.SetContent(container.NewBorder(
		container.NewHBox(
			layout.NewSpacer(), ui.Buttons(wd, a),
		), // push the buttons to the right most side
		nil,
		nil,
		nil,
		ui.ListWords("words.dat", wd, a), // center, fills rest of screen
	))

	// run the program and block the process
	wd.ShowAndRun()
}
