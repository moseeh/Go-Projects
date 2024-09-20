package main

import (
	"fmt"
	"os"

	"my-ls-1/utils"
)

func main() {
	options, paths := utils.ParseArgs(os.Args[1:])
	if len(paths) == 0 {
		paths = []string{"."}
	}
	utils.SortPath(paths)

	for i, path := range paths {
		if i > 0 {
			fmt.Println()
		}
		if len(paths) > 1 {
			fmt.Printf("%s:\n", path)
		}
		err := utils.ListDir(path, options, "")
		if err != nil {
			fmt.Fprintf(os.Stderr, "my-ls: %s: %v\n", path, err)
		}
	}
}
