package utils

import (
	"fmt"
	"strings"
)

func DisplayText(input string, contentLines []string) (st string, err error) {
	input = strings.ReplaceAll(input, "\\n", "\r\n")
	input = strings.ReplaceAll(input, "\\t", "    ")

	wordslice := strings.Split(input, "\r\n")

	for _, word := range wordslice {
		if word == "" {
			st += fmt.Sprintln()
		} else {
			if IsEnglish(word) {
				st += PrintWord(word, contentLines)
			} else {
				return "", fmt.Errorf("invalid input: %s ", word)
			}
		}
	}
	return st, nil
}
