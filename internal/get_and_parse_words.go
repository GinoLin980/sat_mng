package internal

import (
	"errors"
	"os"
	"strings"
)

// [][]string is 2D array
func ReadWords(filename string) ([][]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	content := string(bytes)
	if content == "" {
		return nil, errors.New("Empty file!")
	}

	return parse(content), nil
}

func parse(content string) [][]string {
	result := [][]string{}
	for line := range strings.Lines(content) {
		result = append(result, strings.SplitN(line, " ", 3))
	}

	return result
}
