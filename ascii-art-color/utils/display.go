package utils

import (
	"fmt"
	"strings"
)

func DisplayText(color string, input string, contentLines []string) {
	if input == "" {
		return
	}
	if input == "\\n" || input == "\n" {
		fmt.Println()
		return
	}
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "    ")
	// split the input string with the "\\n" into a slice strings
	wordslice := strings.Split(input, "\\n")

	for _, word := range wordslice {
		if word == "" {
			fmt.Println()
		} else {
			if IsEnglish(word) {
				PrintWord(color, word, contentLines)
			} else {
				fmt.Println("Invalid input:", word)
			}
		}
	}
}
