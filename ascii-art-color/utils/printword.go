package utils

import (
	"fmt"
	"strings"
)

// printWord prints a word using predefined lines from the content file
func PrintWord(color string, word string, tobecolored string, contentLines []string) {
	if tobecolored == "" {
		// If the substring to be colored is empty, just print the word without coloring
		linesOfSlice := make([]string, 8)
		for _, v := range word {
			for i := 1; i < 9; i++ {
				linesOfSlice[i-1] += contentLines[int(v-32)*9+i]
			}
		}
		fmt.Println(strings.Join(linesOfSlice, "\n"))
		return
	}
	linesOfSlice := make([]string, 8)
	start := 0
	if strings.Contains(word, tobecolored) && len(tobecolored) > 1 {
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
	} else {
		for _, v := range word {
			isColored := strings.ContainsRune(tobecolored, v)
			if isColored {
				for i := 1; i < 9; i++ {
					linesOfSlice[i-1] += Color(contentLines[int(v-32)*9+i], color)
				}
			} else {
				for i := 1; i < 9; i++ {
					linesOfSlice[i-1] += contentLines[int(v-32)*9+i]
				}
			}
		}
	}

	fmt.Println(strings.Join(linesOfSlice, "\n"))
}
