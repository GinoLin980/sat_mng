package internal

import "os"

func SaveWords() error {
	if len(PendingWords) == 0 {
		return ErrEmptyPendingWords
	}

	file, err := os.OpenFile(Filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	toAppend := buildWordsToSave(PendingWords)

	if _, err := file.WriteString(toAppend); err != nil {
		return err
	}

	Saved = true

	return nil
}

func buildWordsToSave(words []Vocabulary) string {
	result := ""

	for _, word := range words {
		result += word.Word + " " + word.WordType + " " + word.Description + "\n"
	}

	return result
}
