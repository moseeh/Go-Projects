package asciiArt

import "fmt"

// Print the banner for a line of text
func PrintLineBanner(line string, bannerMap map[int][]string) {
	if line == "" {
		fmt.Println()
		return
	}

	output := make([]string, 8)

	for _, char := range line {
		banner, exists := bannerMap[int(char)]
		if !exists {
			fmt.Printf("Character %c not found in banner map\n", char)
			continue
		}

		for i := 0; i < 8; i++ {
			output[i] += banner[i]
		}
	}

	// Print the final output
	for _, outLine := range output {
		fmt.Println(outLine)
	}
}
