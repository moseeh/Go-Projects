package utils

import (
	"strings"

	"my-ls-1/models"
)

// check command line arguments for flags and paths
func ParseArgs(args []string) (models.Options, []string) {
	var options models.Options
	paths := []string{}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if len(arg) > 1 && arg[1] != '-' {
				// Handle combined short flags
				for _, flag := range arg[1:] {
					switch flag {
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
					default:
						paths = append(paths, arg)
					}
				}
			} else {
				// Handle long flags and single short flags
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
					paths = append(paths, arg)
				}
			}
		} else {
			paths = append(paths, arg)
		}
	}
	return options, paths
}
