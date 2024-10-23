package utils

import (
	"fmt"
	"strings"
)

const Reset = "\033[0m"

func Color(s, color string) string {
	// get rgb values for coloring
	c, err := FindColor(strings.ToLower(color))
	if err != nil {
		return "COLOR NOT FOUND: "

	}

	// return fmt.Sprintf("%s%s%s", c, s, Reset)
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s%s", c[0], c[1], c[2], s, Reset)
}
