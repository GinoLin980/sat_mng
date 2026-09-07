package internal

import (
	"fmt"
	"os"
	"strings"
)

// [][]string is 2D array
func ReadWords(filename string) ([][]string, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	if info.Size() == 0 {
		return nil, ErrEmptyFile
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	content := string(bytes)
	if content == "" {
		return nil, ErrEmptyFile
	}

	return parse(content)
}

func parse(content string) ([][]string, error) {
	result := [][]string{}
	errText := ""
	lineNum := 1

	// split the words by spliting on space(at most 3 items in slice)
	for line := range strings.Lines(content) {
		// skip lines without a caracter
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}

		splited := strings.SplitN(line, " ", 3)
		if len(splited) != 3 {
			errText += fmt.Sprintf("line %d: invalid word %s\n", lineNum, line)
		} else {
			result = append(result, splited)
		}

		lineNum++
	}

	if errText == "" {
		return result, nil
	}

	return result, fmt.Errorf(
		"%w:\n%s",
		ErrReadInvalidWords,
		errText,
	)
}
