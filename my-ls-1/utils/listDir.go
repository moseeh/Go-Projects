package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"my-ls-1/models"
)

func ListDir(path string, options models.Options, files bool) error {
	// Use filepath.Clean to handle '.' and '..' but preserve multiple slashes
	cleanPath := filepath.Clean(path)

	fileInfo, err := os.Stat(cleanPath)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		// If it's a file, just print its info and return
		fileInfos := []models.FileInfo{getFileInfo(cleanPath, fileInfo, options)}
		if options.Long {
			printLongFormat(fileInfos, path, files)
		} else {
			printShortFormat(fileInfos)
		}
		return nil
	}

	entries, err := os.ReadDir(cleanPath)
	if err != nil {
		return err
	}

	fileInfos := getFileInfos(cleanPath, entries, options)
	sortEntries(fileInfos, options)

	if options.Long {
		if options.Recursive {
			fmt.Printf("\n%s:\n", path)
		}
		printLongFormat(fileInfos, path, files)
	} else {
		if options.Recursive {
			fmt.Printf("\n%s:\n", path)
		}
		printShortFormat(fileInfos)
	}

	if options.Recursive {
		for _, info := range fileInfos {
			if info.IsDir && info.Name != "." && info.Name != ".." && (options.All || !strings.HasPrefix(info.Name, ".")) {
				fullPath := strings.TrimSuffix(path, "/") + "/" + info.Name
				err := ListDir(fullPath, options, files)
				if err != nil {
					fmt.Fprintf(os.Stderr, "my-ls: %v\n", err)
				}
			}
		}
	}

	return nil
}
