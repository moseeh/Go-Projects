package utils

import "flag"

var ReversePtr = flag.String("reverse", "", "reverse the ascii art")

// parses command line flags from os.Args[1:]
func ParseFlag() {
	flag.Parse()
}
