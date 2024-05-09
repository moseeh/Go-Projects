package utils

import (
	"strings"
)

func SplitFile(s string) []string {
	if strings.Contains(s, "o") {
		return strings.Split(s, "\r\n")
	}
	return strings.Split(s, "\n")
}
