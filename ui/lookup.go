package ui

import (
	"fmt"
	"sat_word_list/internal"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Lookup(app fyne.App) func() {
	return func() {
		lookupWindow := app.NewWindow("Search a word")
		lookupWindow.Resize(fyne.NewSize(300, 200))
		lookupWindow.SetFixedSize(true)

		lookupWindow.SetContent(buildLookupWindow(app, lookupWindow))
		lookupWindow.Show()
	}
}

func buildLookupWindow(app fyne.App, lookupWindow fyne.Window) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.PlaceHolder = "What to search?"

	return container.NewPadded(
		container.NewVBox(
			entry,
			layout.NewSpacer(),
			widget.NewButton("Search", func() {
				if vocabulary, found := contentForLookup[strings.TrimSpace(strings.ToLower(entry.Text))]; found {
					buildFoundWindow(app, vocabulary)
					lookupWindow.Close()
					return
				}

				dialog.ShowError(fmt.Errorf("%s not found!", entry.Text), lookupWindow)
			}),
		),
	)
}

func buildFoundWindow(app fyne.App, vocalbulary internal.Vocabulary) {
	foundWindow := app.NewWindow("Result found!")
	foundWindow.Resize(fyne.NewSize(300, 200))
	foundWindow.SetFixedSize(true)

	foundWindow.SetContent(
		container.NewVBox(
			widget.NewLabel(vocalbulary.Pretty()),
			layout.NewSpacer(),
			widget.NewButton("Ok!", func() { foundWindow.Close() }),
		),
	)
	foundWindow.Show()
}
