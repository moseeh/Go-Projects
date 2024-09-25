package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"my-ls-1/models"
)

func ListDir(path string, options models.Options) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		// If it's a file, just print its info and return
		fileInfos := []models.FileInfo{getFileInfo(path, fileInfo, options)}
		if options.Long {
			printLongFormat(fileInfos, path)
		} else {
			printShortFormat(fileInfos)
		}
		return nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	fileInfos := getFileInfos(path, entries, options)
	sortEntries(fileInfos, options)

	if options.Long {
		printLongFormat(fileInfos, path)
	} else {
		printShortFormat(fileInfos)
	}

	if options.Recursive {
		for _, info := range fileInfos {
			if info.IsDir && info.Name != "." && info.Name != ".." && (options.All || !strings.HasPrefix(info.Name, ".")) {
				fullPath := filepath.Join(path, info.Name)
				fmt.Printf("\n%s:\n", fullPath)
				err := ListDir(fullPath, options)
				if err != nil {
					fmt.Fprintf(os.Stderr, "my-ls: %v\n", err)
				}
			}
		}
	}

	return nil
}
