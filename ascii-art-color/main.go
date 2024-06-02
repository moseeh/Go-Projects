package main

import (
	"color/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	//check enough command line arguments
	if len(os.Args[1:]) < 1 || len(os.Args[1:]) > 3 || (!strings.HasPrefix(os.Args[1], "--color=") && len(os.Args[1:]) != 1) || (strings.HasPrefix(os.Args[1], "--color=") && len(os.Args[1:]) == 1) {
		fmt.Print("Usage: go run . [OPTION] [STRING]\n\n")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
	inputslice := utils.ParseFlags()
	color := *utils.ColorPtr
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
	utils.DisplayText(color, inputslice, contentLines)
}
