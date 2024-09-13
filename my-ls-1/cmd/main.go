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

	for _, path := range paths {
		err := utils.ListDir(path, options, "")
		if err != nil {
			fmt.Fprintf(os.Stderr, "my-ls: %v\n", err)
		}
	}
}
