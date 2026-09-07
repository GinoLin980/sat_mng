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

func NewWord(app fyne.App) func() {
	return func() {
		newWordWindow := app.NewWindow("Adding a New Word")
		newWordWindow.Resize(fyne.NewSize(300, 200))
		newWordWindow.SetFixedSize(true)

		newWordWindow.SetContent(newWordUI(newWordWindow))
		newWordWindow.Show()
	}
}

// private UI builder
func newWordUI(wd fyne.Window) fyne.CanvasObject {
	word := internal.Vocabulary{}

	entry := widget.NewEntry()
	entry.PlaceHolder = "Enter a new word here"
	entry.OnChanged = func(s string) { word.Word = strings.TrimSpace(s) }

	wordTypeSelect := widget.NewSelect([]string{"Noun", "Verb", "Adjective"}, func(s string) {
		switch s {
		case "Noun":
			word.WordType = "n"
		case "Verb":
			word.WordType = "v"
		case "a":
			word.WordType = "a"
		}
	})

	descriptionEntry := widget.NewMultiLineEntry()
	descriptionEntry.PlaceHolder = "Enter description here"
	descriptionEntry.OnChanged = func(s string) { word.Description = strings.TrimSpace(s) }

	return container.NewPadded(
		container.NewVBox(
			entry,
			layout.NewSpacer(),
			wordTypeSelect,
			layout.NewSpacer(),
			descriptionEntry,
			container.NewHBox(newWordCancelButton(wd, &word), newWordConfirmButton(wd, &word)),
		),
	)
}

func newWordCancelButton(wd fyne.Window, word *internal.Vocabulary) fyne.CanvasObject {
	return widget.NewButton("Cancel",
		func() {
			if word.Word != "" || word.WordType != "" || word.Description != "" {
				fmt.Println(word)
				dialog.ShowConfirm(
					"Are you sure?",
					"Content exists in \"Word\" or \"Description\"\nAre you sure you want to discard them",
					func(b bool) {
						if b {
							wd.Close()
							return
						}
					}, // if choose yes
					wd,
				)
				// end of flow to prevent window closed by command below
				return
			}

			// if choose yes
			wd.Close()
		},
	)
}

func newWordConfirmButton(wd fyne.Window, word *internal.Vocabulary) fyne.CanvasObject {
	return widget.NewButton("Confirm", func() {
		if word == nil {
			dialog.ShowError(fmt.Errorf("nil pointer provided in newWordConfirmButton!"), wd)
			return
		}

		if word.Word == "" || word.WordType == "" || word.Description == "" {
			dialog.ShowError(fmt.Errorf("Missing field(s)\nWord: %s\nWord Type: %s\nDescription: %s", word.Word, word.WordType, word.Description), wd)
			return
		}

		content = append(content, []string{word.Word, word.WordType, word.Description}) // update the table items
		contentForLookup[strings.ToLower(word.Word)] = *word
		internal.PendingWords[strings.ToLower(word.Word)] = *word
		UpdateTable(wd)
		internal.Saved = false

		wd.Close()
	})
}
