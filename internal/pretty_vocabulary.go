package internal

import "fmt"

// return a
// Word: {Word}
// Word Type: {WordType}
// Description: {Description}
func (v Vocabulary) Pretty() string {
	return fmt.Sprintf("Word: %s\nWord Type: %s\nDescription: %s", v.Word, v.WordType, v.Description)
}
