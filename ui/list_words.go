package ui

import (
	"errors"
	"fmt"
	"os"
	"sat_word_list/internal"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var table *widget.Table

func ListWords(filename string, wd fyne.Window, app fyne.App) fyne.CanvasObject {

	var err error // prevent using content, err := readwords, as it will create local content var
	contentForLookup, content, err = internal.ReadWords(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// no file found, end of program
			dl := dialog.NewError(err, wd)
			dl.SetOnClosed(app.Quit) // app will quit instantly on prompt's close
			dl.Show()
			return widget.NewLabel("unable to ReadWords")
		} else if errors.Is(err, internal.ErrReadInvalidWords) {
			dialog.ShowError(fmt.Errorf("File: %s\nError: %s", filename, err), wd)
		} else {
			dialog.ShowError(fmt.Errorf("File: %s\nError: %s", filename, err), wd)
			return widget.NewLabel("unable to ReadWords")
		}
	}

	table = widget.NewTable(
		func() (int, int) {
			return len(content), len(content[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			if id.Row < len(content) && id.Col < len(content[id.Row]) {
				label.SetText(content[id.Row][id.Col])
			} else {
				label.SetText("")
			}
		},
	)

	for col := 0; col < len(content[0]); col++ {
		maxWidth := 0
		for _, row := range content {
			if col >= len(row) {
				continue
			}
			size := fyne.MeasureText(row[col], float32(theme.TextSize()), fyne.TextStyle{})
			if int(size.Width) > maxWidth {
				maxWidth = int(size.Width)
			}
		}
		table.SetColumnWidth(col, float32(maxWidth)+theme.Padding()*4)
	}
	return table
}

func UpdateTable(wd fyne.Window) {
	if table == nil {
		dialog.ShowError(errors.New("Table isn't initialized, can't be saved or updated"), wd)
		return
	}

	table.Refresh()
}
