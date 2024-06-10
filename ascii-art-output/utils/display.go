package utils

import (
	"fmt"
	"strings"
)

func DisplayText(input string, banner string, contentLines []string) {
	s := ""
	if input == "" {
		return
	}
	if input == "\\n" || input == "\n" {
		s += "\n"
		return
	}
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "    ")
	// split the input string with the "\\n" into a slice strings
	wordslice := strings.Split(input, "\\n")

	for _, word := range wordslice {
		if word == "" {
			s += "\n"
		} else {
			if IsEnglish(word) {
				s += PrintWord(word, banner, contentLines)
			} else {
				fmt.Println("Invalid input:", word)
			}
		}
	}
	WriteFile(s, banner)
}
