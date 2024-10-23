package utils

import (
	"strings"
)

// CheckArgs checks if the input arguments meet the specified requirements
func CheckArgs(args []string) bool {
	argCount := len(args)

	switch argCount {
	case 1:
		// One argument: It cannot be a flag, just a string
		return !isFlag(args[0])
	case 2:
		// Two arguments: Either a flag and a string or a string and a banner
		if isFlag(args[0]) && !isFlag(args[1]) {
			return true
		} else if !isFlag(args[0]) {
			banner := strings.ToLower(args[1])
			return isBanner(banner)
		}
		return false
	case 3:
		// Three arguments: Either a flag and a string and a banner or flag (color) substring and a string
		if isFlag(args[0]) {
			if strings.HasPrefix(args[0], "--color=") {
				return !isFlag(args[1]) && !isFlag(args[2])
			} else {
				banner := strings.ToLower(args[2])
				return !isFlag(args[1]) && isBanner(banner)
			}
		}
		return false
	case 4:
		// Four arguments: Only flag (color) substring string and banner
		if isFlag(args[0]) && strings.HasPrefix(args[0], "--color=") {
			banner := strings.ToLower(args[3])
			return !isFlag(args[1]) && !isFlag(args[2]) && isBanner(banner)
		}
		return false
	default:
		// More than 4 arguments are invalid
		return false
	}
}

// isFlag checks if a string is a flag in the format --flagname=value
func isFlag(arg string) bool {
	return strings.HasPrefix(arg, "--")
}

// isBanner checks if a string is a valid banner
func isBanner(arg string) bool {
	return arg == "standard" || arg == "thinkertoy" || arg == "shadow"
}
