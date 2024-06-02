package utils

import (
	"fmt"
	"strings"
)

func DisplayText(color string, inputslice []string, contentLines []string) {
	tobecolored := ""
	input := ""
	if len(inputslice) == 2 {
		tobecolored = inputslice[0]
		input = inputslice[1]
	} else {
		tobecolored = inputslice[0]
		input = inputslice[0]
	}
	if input == "" {
		return
	}
	if input == "\\n" || input == "\n" {
		fmt.Println()
		return
	}
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "    ")
	tobecolored = strings.ReplaceAll(tobecolored, "\n", "\\n")
	tobecolored = strings.ReplaceAll(tobecolored, "\\t", "    ")
	// split the input string with the "\\n" into a slice strings
	wordslice := strings.Split(input, "\\n")

	for _, word := range wordslice {
		if word == "" {
			fmt.Println()
		} else {
			if IsEnglish(word) {
				PrintWord(color, word, tobecolored, contentLines)
			} else {
				fmt.Println("Invalid input:", word)
			}
		}
	}
}
