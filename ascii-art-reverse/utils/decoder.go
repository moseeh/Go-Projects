package utils

import (
	"strings"
)

// decodes each group of 8 lines to make the art
func decoder(lines []string, artToChar map[string]string) string {
	artHeight := 8
	var decodedString strings.Builder
	lineLength := len(lines[0])
	for i := 0; i < lineLength; {
		maxWidth := 0
		for j := 0; j < artHeight; j++ {
			if i < len(lines[j]) {
				if maxWidth < len(lines[j][i:]) {
					maxWidth = len(lines[j][i:])
				}
			}
		}
		// moves horizontally with the 8 lines to check if the art is available the artmap
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
			i++
		}
	}

	return decodedString.String()
}
