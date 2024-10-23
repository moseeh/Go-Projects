package utils

import "strings"

func ReplaceSpecial(input string) string {
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "    ")
	input = strings.ReplaceAll(input, "\\r", "\r")
	input = strings.ReplaceAll(input, "\\v", "\v")
	input = strings.ReplaceAll(input, "\\f", "\f")
	input = strings.ReplaceAll(input, "\\a", "\a")
	input = strings.ReplaceAll(input, "\\b", "\b")
	input = strings.ReplaceAll(input, "\\v", "\v")

	return input
}
