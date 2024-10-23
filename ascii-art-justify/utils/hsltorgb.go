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
	// getting hsl values to be converted to rgb
	h, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	s, err2 := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(parts[1], "%")))
	l, err3 := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(parts[2], "%")))
	if err1 != nil || err2 != nil || err3 != nil {
		return rgb, errors.New("invalid hsl color")
	}

	return HSLToRGB(h, s, l), nil
}

// HSLToRGB converts HSL values to RGB values
func HSLToRGB(h, s, l int) [3]int {
	var rgb [3]int
	h = h % 360
	s = s % 101
	l = l % 101

	c := (1.0 - float64(abs(2*l-100))/100.0) * float64(s) / 100.0
	x := c * (1.0 - float64(abs(h/60%2-1)))
	m := float64(l)/100.0 - c/2.0

	var r, g, b float64
	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	rgb[0] = int((r + m) * 255.0)
	rgb[1] = int((g + m) * 255.0)
	rgb[2] = int((b + m) * 255.0)

	return rgb
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
