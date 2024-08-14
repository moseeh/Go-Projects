package utils

import (
	"os"
	"strings"
)

// checks validity of arguments passed
func CheckArgs() bool {
	length := len(os.Args)

	switch length {
	case 1:
		return false
	case 2:
		if strings.HasPrefix(os.Args[1], "-") && !strings.HasPrefix(os.Args[1], "--reverse=") {
			return false
		}
		return true
	case 3:
		if strings.HasPrefix(os.Args[1], "-") {
			return false
		}
		banner := strings.ToLower(os.Args[2])
		if banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
			return true
		} else {
			return false
		}
	}
	return false
}
