package utils

import (
	"fmt"
	"strings"
)

// printWord prints a word using predefined lines from the content file
func PrintWord(color string, word string, tobecolored string, contentLines []string) {
	linesOfSlice := make([]string, 9)
	start := 0

	for start < len(word) {
		if strings.HasPrefix(word[start:], tobecolored) {
			for i := 0; i < 9; i++ {
				for _, v := range tobecolored {
					linesOfSlice[i] += Color(contentLines[int(v-32)*9+i], color)
				}
			}
			start += len(tobecolored)
		} else {
			v := rune(word[start])
			for i := 0; i < 9; i++ {
				linesOfSlice[i] += contentLines[int(v-32)*9+i]
			}
			start++
		}
	}

	fmt.Println(strings.Join(linesOfSlice, "\n"))
}

