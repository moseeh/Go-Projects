package utils

import (
	"fmt"
	"strings"
)

// DisplayText generates ASCII art from the input string using the provided content lines.
func DisplayText(input string, contentLines []string) (st string, err error) {
	// Splitting with Windows-style newline "\r\n"
	wordslice := strings.Split(input, "\r\n")

	for _, word := range wordslice {
		if word == "" {
			st += fmt.Sprintln()
		} else {
			if IsEnglish(word) {
				st += PrintWord(word, contentLines)
			} else {
				return "", fmt.Errorf("invalid input: %s", word)
			}
		}
	}
	return st, nil
}
