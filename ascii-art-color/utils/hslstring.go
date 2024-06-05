package utils

import (
	"errors"
	"strconv"
	"strings"
)

// HSLStringToRGB converts an hsl(h, s, l) string to an RGB array
func HSLStringToRGB(hslStr string) ([3]int, error) {
	var rgb [3]int
	hslStr = strings.TrimPrefix(hslStr, "hsl(")
	hslStr = strings.TrimSuffix(hslStr, ")")
	parts := strings.Split(hslStr, ",")
	if len(parts) != 3 {
		return rgb, errors.New("invalid hsl color")
	}

	h, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	s, err2 := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(parts[1], "%")))
	l, err3 := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(parts[2], "%")))
	if err1 != nil || err2 != nil || err3 != nil {
		return rgb, errors.New("invalid hsl color")
	}

	return HSLToRGB(h, s, l), nil
}
