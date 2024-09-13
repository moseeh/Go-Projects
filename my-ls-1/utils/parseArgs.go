package utils

import (
	"fmt"
	"os"
	"strings"

	"my-ls-1/models"
)

func ParseArgs(args []string) (models.Options, []string) {
	var options models.Options
	var paths []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			switch arg {
			case "-l":
				options.Long = true
			case "-R", "--recursive":
				options.Recursive = true
			case "-a", "--all":
				options.All = true
			case "-r", "--reverse":
				options.Reverse = true
			case "-t", "--time":
				options.SortByTime = true
			default:
				fmt.Printf("Warning: Unknown flag %s\n", arg)
				os.Exit(0)
			}
		} else {
			paths = append(paths, arg)
		}
	}
	return options, paths
}
