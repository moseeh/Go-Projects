package utils

import (
	"errors"
	"strconv"
)

func HexToRGB(hex string) ([3]int, error) {
	var rgb [3]int
	if hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) != 6 {
		return rgb, errors.New("invalid hex color")
	}

	for i := 0; i < 3; i++ {
		val, err := strconv.ParseInt(hex[i*2:i*2+2], 16, 64)
		if err != nil {
			return rgb, err
		}
		rgb[i] = int(val)
	}

	return rgb, nil
}
