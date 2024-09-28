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
	newPaths := []string{}
	files := false
	for _, path := range paths {
		fileinfo, err := os.Lstat(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "my-ls: %s: %v\n", path, err)
			return
		}
		if fileinfo.IsDir() {
			newPaths = append(newPaths, path)
		} else {
			files = true
			err := utils.ListDir(path, options, files)
			if err != nil {
				fmt.Fprintf(os.Stderr, "my-ls: %s: %v\n", path, err)
				return
			}
		}
	}
	files = false

	for i, path := range newPaths {
		if i == 0 && len(paths) > 1 {
			fmt.Println()
		}
		if i > 0 {
			fmt.Println()
		}
		if len(paths) > 1 {
			fmt.Printf("%s:\n", path)
		}
		err := utils.ListDir(path, options, files)
		if err != nil {
			fmt.Fprintf(os.Stderr, "my-ls: %s: %v\n", path, err)
			return
		}
	}
}
