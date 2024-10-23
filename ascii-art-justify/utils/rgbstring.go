package utils

import (
	"errors"
	"strconv"
	"strings"
)

// RGBStringToRGB converts an rgb(r, g, b) string to an RGB array
func RGBStringToRGB(rgbStr string) ([3]int, error) {
	var rgb [3]int
	rgbStr = strings.TrimPrefix(rgbStr, "rgb(")
	rgbStr = strings.TrimSuffix(rgbStr, ")")
	parts := strings.Split(rgbStr, ",")
	if len(parts) != 3 {
		return rgb, errors.New("invalid rgb color")
	}
	// take the rgb string and convert to int for the [3]int
	for i := 0; i < 3; i++ {
		val, err := strconv.Atoi(strings.TrimSpace(parts[i]))
		if err != nil {
			return rgb, err
		}
		rgb[i] = val
	}

	return rgb, nil
}
