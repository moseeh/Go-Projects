package asciiArt

import (
	"fmt"
	"os"
	"strings"
)

func ArtMaker() {
	filename := BannerFile()

	if os.Args[1] == "\\n" {
		fmt.Println()
		return
	}

	// Load the banner map from the file
	bannerMap, err := LoadBannerMap(filename)
	if err != nil {
		fmt.Println("error loading banner map:", err)
		return
	}

	// Process the provided argument
	args := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	args = strings.ReplaceAll(args, "\\t", "    ")
	lines := strings.Split(args, "\n")

	// Generate the ASCII art for each line
	for _, line := range lines {
		PrintLineBanner(line, bannerMap)
	}

}
