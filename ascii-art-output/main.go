package main

import (
	"fmt"
	"os"

	"ascii-art/utils"
)

func main() {
	// checks if there input provived
	if len(os.Args) < 2 {
		fmt.Println("Please provide text to display.")
		return
	}
	args := utils.ParseFlag()
	banner := *utils.OutputPtr
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
	utils.DisplayText(args[0], banner, contentLines)
}
