package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// reading into the file
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		panic(err)
	}
	// store content into a slice of strings
	contentslice := strings.Split(string(content), "\n")
	word := ""
	stringslice := os.Args[1:]
	if len(stringslice) == 1 {
		word = stringslice[0]
		if word == "" {
			return
		}
		if word == "\\n" {
			fmt.Println()
			return
		}
	} else {
		word = strings.Join(stringslice, " ")
	}

	words := strings.Split(word, "\\n")

	for _, input := range words {
		if input == "" {
			fmt.Println()
		} else {
			linesOfSlice := make([]string, 9)
			for _, v := range input {
				for i := 1; i <= 9; i++ {
					linesOfSlice[i-1] += contentslice[int(v-32)*9+i]
				}
			}
			fmt.Print(strings.Join(linesOfSlice, "\n"))
		}
	}
}
