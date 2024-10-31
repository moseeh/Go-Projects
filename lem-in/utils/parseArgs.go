package utils

import "os"

// checks validity of command line arguments to see if
func ParseArgs() (string, string) {
	args := os.Args[1:]

	switch len(args) {
	case 0:
		return "", "Please provide an input file"
	case 1:
		return args[0], ""
	default:
		return "", "too many arguments"
	}
}
