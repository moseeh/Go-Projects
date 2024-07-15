package utils

import (
	"fmt"
	"os"
)

func ReadBannerFile(banner string) ([]byte, error) {
	var filename string
	switch banner {
	case "standard":
		filename = "banners/standard.txt"
	case "shadow":
		filename = "banners/shadow.txt"
	case "thinkertoy":
		filename = "banners/thinkertoy.txt"
	default:
		return nil, fmt.Errorf("invalid banner type")
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return content, nil
}
