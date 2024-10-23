package utils

import (
	"fmt"
	"strings"
)

func DisplayText(color string, inputslice []string, contentLines []string) string {
	var str strings.Builder

	// Initialize words to be colored and input
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
		return ""
	}
	if input == "\\n" || input == "\n" {
		fmt.Println()
		return ""
	}

	input = ReplaceSpecial(input)
	tobecolored = ReplaceSpecial(tobecolored)
	wordslice := strings.Split(input, "\\n")
	tbcsplit := strings.Split(tobecolored, "\\n")

	for i, word := range wordslice {
		if word == "" {
			str.WriteString("\n")
		} else {
			if IsEnglish(word) && IsEnglish(tobecolored) {
				if i < len(tbcsplit) && len(tbcsplit) > 1 {
					// Coloring elements of each word on each line
					str.WriteString(PrintWord(color, word, tbcsplit[i], contentLines))
				}
				if i >= len(tbcsplit) && len(tbcsplit) > 1 {
					tobecolored = ""
					// Unrepresented lines are left uncolored
					str.WriteString(PrintWord(color, word, tobecolored, contentLines))
				}
				// Color each element in every line
				if len(tbcsplit) == 1 {
					str.WriteString(PrintWord(color, word, tobecolored, contentLines))
				}
				if i != len(wordslice)-1 {
					str.WriteString("\n")
				}
			} else {
				fmt.Println("Invalid input:", word, tobecolored)
			}
		}
	}
	return str.String()
}
