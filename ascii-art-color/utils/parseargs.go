package utils

import (
	"fmt"
	"strings"
)

func ParseArguments(args []string) (string, string) {
	flag := args[0]
	color := ""
	input := ""
	if strings.HasPrefix(flag, "--color=") {
		color = flag[8:]
	} else {
		fmt.Println("incorrect color flag")
	}
	if len(args) == 2 {
		input = args[1]
	}
	return color, input
}
