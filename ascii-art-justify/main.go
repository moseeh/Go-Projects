package main

import (
	"fmt"
	"os"
	"strings"

	"justify/utils"
)

func main() {
	// check enough command line arguments
	if !utils.CheckArgs(os.Args[1:]) {
		fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
		fmt.Println("Example: go run . --align=right something standard")
		return
	}
	inputslice := utils.ParseFlags()
	color := *utils.ColorPtr
	output := *utils.OutputPtr
	align := *utils.AlignPtr
	asciiArt := ""
	var filename string
	inputslice, filename = utils.GetBannerFile(inputslice)
	var contentLines []string
	if utils.CheckFileSize(filename) {

		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("invalid text file")
			return
		}
		contentLines = utils.SplitFile(string(content))
		if color != "" {
			_, errcolor := utils.FindColor(strings.ToLower(color))
			if errcolor != nil {
				fmt.Println(errcolor, color)
				return
			}
		}
		asciiArt = utils.DisplayText(color, inputslice, contentLines)
	} else {
		fmt.Println("Invalid/Corrupted bannerfile file")
		return
	}
	if output != "" {
		err := utils.CreateFile(output, asciiArt)
		if err != nil {
			fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
			fmt.Println("Example: go run . --output=file.txt something standard")
			return
		}
	}
	if align != "" {
		utils.Align(align, asciiArt, inputslice[0], contentLines)
	} else {
		fmt.Println(asciiArt)
	}
}
