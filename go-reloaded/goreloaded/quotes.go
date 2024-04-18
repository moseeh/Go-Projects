package goreloaded

import "strings"

func QuotationMarks(s string) string {
	s = strings.ReplaceAll(s, "' ", " '")
	s = strings.ReplaceAll(s, " '", "'")
	return s
}
