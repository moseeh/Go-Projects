package main

import (
	"color/utils"
	"fmt"
	"os"
)

func main() {
	// check enough command line arguments
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return
	}
	color, input := utils.ParseArguments(os.Args[1:])
	// reads contents of the provided
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("invalid text file")
		return
	}
	contentLines := utils.SplitFile(string(content))
	// checks if the file provided contains 856 lines
	if len(contentLines) != 856 {
		fmt.Println("invalid text file")
		return
	}
	utils.DisplayText(color, input, contentLines)
}
