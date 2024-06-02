package utils

import (
	"fmt"
	"strings"
)

// printWord prints a word using predefined lines from the content file
func PrintWord(color string, word string, tobecolored string, contentLines []string) {
	linesOfSlice := make([]string, 8)
	start := 0

	for start < len(word) {
		if strings.HasPrefix(word[start:], tobecolored) {
			for i := 1; i < 9; i++ {
				for _, v := range tobecolored {
					linesOfSlice[i-1] += Color(contentLines[int(v-32)*9+i], color)
				}
			}
			start += len(tobecolored)
		} else {
			v := rune(word[start])
			for i := 1; i < 9; i++ {
				linesOfSlice[i-1] += contentLines[int(v-32)*9+i]
			}
			start++
		}
	}

	fmt.Println(strings.Join(linesOfSlice, "\n"))
}
