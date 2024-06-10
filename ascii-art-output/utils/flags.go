package utils

import "flag"

var OutputPtr = flag.String("output", "banner.txt", "file to be written to")

func ParseFlag() []string{
	flag.Parse()
	return flag.Args()
}
