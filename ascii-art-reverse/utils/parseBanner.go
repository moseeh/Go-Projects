package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// reads banner file and returns a map of the art as the key
func ParseBannerFile(filename string) map[string]string {
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
