package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	standardFile := "banners/standard.txt"
	sampleFile := "sample.txt"

	// Create a mapping from standard.txt
	artToChar := parseStandardFile(standardFile)

	// Decode the sample.txt using the mapping
	decodedString := decodeSampleFile(sampleFile, artToChar)

	// Print the result
	fmt.Println(decodedString)
}

func parseStandardFile(filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	artToChar := make(map[string]string)
	var art []string
	charCode := 32 // Start with ASCII code 32

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(art) == 8 { // Each character's art is 8 lines
				artToChar[strings.Join(art, "\n")] = string(rune(charCode))
				art = []string{}
				charCode++
			}
		} else {
			art = append(art, line)
		}
	}

	if len(art) == 8 { // Last character case
		artToChar[strings.Join(art, "\n")] = string(rune(charCode))
	}

	return artToChar
}

func decodeSampleFile(filename string, artToChar map[string]string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Each character's ASCII art is 8 lines high
	artHeight := 8
	var decodedString strings.Builder

	// Iterate through the lines and match segments to characters
	lineLength := len(lines[0])
	for i := 0; i < lineLength; {
		//var segment []string
		maxWidth := 0
		for j := 0; j < artHeight; j++ {
			if i < len(lines[j]) {
				//segment = append(segment, lines[j][i:])
				if maxWidth < len(lines[j][i:]) {
					maxWidth = len(lines[j][i:])
				}
			} else {
				//segment = append(segment, "")
			}
		}
		//charArt := strings.Join(segment, "\n")
		found := false
		for width := 1; width <= maxWidth; width++ {
			subArt := []string{}
			for j := 0; j < artHeight; j++ {
				if i+width <= len(lines[j]) {
					subArt = append(subArt, lines[j][i:i+width])
				} else {
					subArt = append(subArt, lines[j][i:])
				}
			}
			charArt := strings.Join(subArt, "\n")
			if val, exists := artToChar[charArt]; exists {
				decodedString.WriteString(val)
				i += width
				found = true
				break
			}
		}
		if !found {
			i++ // If no match is found
		}
	}

	return decodedString.String()
}
