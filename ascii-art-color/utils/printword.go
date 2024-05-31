package utils

import (
	"fmt"
	"strings"
)

// printWord prints a word using predefined lines from the content file
func PrintWord(color string, word string, tobecolored string, contentLines []string) {
	linesOfSlice := make([]string, 9)
	for _, v := range word {
		isColored := strings.ContainsRune(tobecolored, v)
		if isColored{
			for i := 1; i <= 9; i++ {
				linesOfSlice[i-1] += Color(contentLines[int(v-32)*9+i], color)
			}
		} else {
			for i := 1; i <= 9; i++ {
				linesOfSlice[i-1] += contentLines[int(v-32)*9+i]
			}
		}

	}
	fmt.Println(strings.Join(linesOfSlice, "\n"))
}
