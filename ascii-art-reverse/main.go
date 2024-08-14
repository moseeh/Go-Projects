package main

import (
	"fmt"

	"ascii-art-reverse/asciiArt"
	"ascii-art-reverse/utils"
)

func main() {
	// Check if an argument is provided
	if !utils.CheckArgs() {
		fmt.Printf("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>\n")
		return
	}

	// initializing banner file paths
	standardFile := "banners/standard.txt"
	shadowFile := "banners/shadow.txt"
	thinkertoyFile := "banners/thinkertoy.txt"

	// parsing flags
	utils.ParseFlag()
	sampleFile := *utils.ReversePtr
	// if no flag is provided, switch to accept printable ascii input and display their banner files in the terminal

	if sampleFile == "" {
		// If no flag is provided, use the Filename function to decide the file to be used
		asciiArt.ArtMaker()
	} else {
		bannerFiles := []string{standardFile, shadowFile, thinkertoyFile}
		var decodedString string
		var success bool

		for _, bannerFile := range bannerFiles {
			artToChar := utils.ParseBannerFile(bannerFile)
			lines, err := utils.ParseFile(sampleFile)
			if err != nil {
				fmt.Println(err)
				return
			}
			decodedString = utils.DecodeFile(lines, artToChar)
			if decodedString != "" {
				success = true
				break
			}
		}

		if success {
			fmt.Print(decodedString)
		} else {
			fmt.Println("Error: Could not decode the ASCII art using any of the banner files.")
		}
	}
}
