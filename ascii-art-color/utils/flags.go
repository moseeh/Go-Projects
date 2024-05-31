package utils

import (
	"flag"
)

var ColorPtr = flag.String("color", "white", "coloring the output")

func ParseFlags() []string{
	flag.Parse()
	return flag.Args()
}
