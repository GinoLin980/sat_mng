package ui

import (
	"sat_word_list/internal"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ListWords(filename string, wd fyne.Window) fyne.CanvasObject {
	content, err := internal.ReadWords(filename)
	if err != nil {
		dialog.ShowError(err, wd)
		return nil
	}

	table := widget.NewTable(
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
		table.SetColumnWidth(col, float32(maxWidth+theme.Padding()*4))
	}
	return table
}
