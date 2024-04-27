package main

import (
	"fmt"
	"os"
	"strings"

	utils "ascii/utils"
)

func main() {
	// checks if there input provived
	if len(os.Args) < 2 {
		fmt.Println("Please provide text to display.")
		return
	}
	// reads contents of the provided
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		panic(err)
	}
	// splits the content of the file into lines
	contentLines := strings.Split(string(content), "\n")
	// concatenates command-line arguments into a singe string
	word := strings.Join(os.Args[1:], " ")
	if word == "\\n" {
		fmt.Println()
		return
	}
	// Replace occurrences of "\t" with spaces
	word = strings.ReplaceAll(word, "\\t", "    ")

	// Split the input string by "\n" to separate words
	words := strings.Split(word, "\\n")

	// looping through each word
	for _, input := range words {
		if input == "" {
			return
		} else {
			utils.PrintWord(input, contentLines)
		}
	}
}
