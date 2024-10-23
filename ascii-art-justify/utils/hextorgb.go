package utils

import (
	"errors"
	"strconv"
	"strings"
)

func HexToRGB(hex string) ([3]int, error) {
	var rgb [3]int

	// Remove the leading '#' if it exists
	hex = strings.TrimPrefix(hex, "#")

	// Normalize the hex string
	if len(hex) == 3 {
		expanded := ""
		for _, c := range hex {
			expanded += string(c) + string(c)
		}
		hex = expanded
	}

	// Check if the length is valid
	if len(hex) != 6 {
		return rgb, errors.New("invalid hex color")
	}

	// Convert each pair of hex digits to an integer
	for i := 0; i < 3; i++ {
		val, err := strconv.ParseInt(hex[i*2:i*2+2], 16, 64)
		if err != nil {
			return rgb, errors.New("invalid hex color")
		}
		rgb[i] = int(val)
	}

	return rgb, nil
}
