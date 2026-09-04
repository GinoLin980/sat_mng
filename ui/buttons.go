package ui

import (
	"sat_word_list/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// set initial value to true, as unmodified means the disk and ram are synced, which equal to saved
var saved bool = true

func Buttons(wd fyne.Window, app fyne.App) []fyne.CanvasObject {
	return []fyne.CanvasObject{
		widget.NewButton("New Word", nil),
		widget.NewButton("Check Meaning", nil),
		widget.NewButton("Save", nil),
		widget.NewButton("Quit", ui.Quit(saved, wd, app)),
	}
}

// 1) To add a new entry.
// ·  Upon clicking this option – a new entry window pops up that requests the word (text input to be typed in), a selection of verb/noun/adjective, and a text input for the meaning of the word, and a submit button. Upon clicking this submit button, the word and its properties must be added to the list of words with the system.

// 2) To check meaning.
// ·  Upon clicking this option – a new window pops up and requests a word (text input to be typed in) and a submit button. Upon clicking this button, the system searches and displays the word and its properties (if found) on the display window or displays an appropriate error message. Note – the words should be searched case insensitive.

// 3) To save the updated word list of words for later use (note if you run the application again – it should include the new added entries).

// 4) Quit Application
