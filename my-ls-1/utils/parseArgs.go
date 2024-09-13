package utils

import (
	"strings"

	"my-ls-1/models"
)

func ParseArgs(args []string) (models.Options, []string) {
	var options models.Options
	var paths []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			for _, ch := range arg[1:] {
				switch ch {
				case 'l':
					options.Long = true
				case 'R':
					options.Recursive = true
				case 'a':
					options.All = true
				case 'r':
					options.Reverse = true
				case 't':
					options.SortByTime = true
				}
			}
		} else {
			paths = append(paths, arg)
		}
	}
	return options, paths
}
