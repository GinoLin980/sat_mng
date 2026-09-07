package internal

import (
	"fmt"
	"os"
	"strings"
)

// [][]string is 2D array, because table ui need 2d array data
func ReadWords(filename string) (map[string]Vocabulary, [][]string, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return nil, nil, err
	}

	if info.Size() == 0 {
		return nil, nil, ErrEmptyFile
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	content := string(bytes)
	if content == "" {
		return nil, nil, ErrEmptyFile
	}

	return parse(content)
}

func parse(content string) (map[string]Vocabulary, [][]string, error) {
	resultForLookup := map[string]Vocabulary{}
	result := [][]string{}
	errText := ""
	lineNum := 1

	// split the words by spliting on space(at most 3 items in slice)
	for line := range strings.Lines(content) {
		// skip lines without a caracter
		line := strings.TrimSpace(line)
		lineNum++

		if line == "" {
			continue
		}

		splited := strings.SplitN(line, " ", 3)
		if len(splited) != 3 {
			errText += fmt.Sprintf("line %d: invalid word %s\n", lineNum, line)
		} else {
			splited[0] = strings.TrimSpace(splited[0])
			splited[1] = strings.TrimSpace(splited[1])
			splited[2] = strings.TrimSpace(splited[2])
			resultForLookup[strings.ToLower(splited[0])] = Vocabulary{
				Word:        splited[0],
				WordType:    splited[1],
				Description: splited[2],
			}
			result = append(result, splited)
		}
	}
	if errText == "" {
		return resultForLookup, result, nil
	}

	return resultForLookup, result, fmt.Errorf(
		"%w:\n%s",
		ErrReadInvalidWords,
		errText,
	)
}
