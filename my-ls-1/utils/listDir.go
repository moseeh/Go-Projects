package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"my-ls-1/models"
)

func ListDir(path string, options models.Options, indent string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	fileInfos := make([]models.FileInfo, 0, len(entries))
	for _, entry := range entries {
		if !options.All && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return err
		}

		fileInfo := models.FileInfo{
			Name:    entry.Name(),
			Mode:    info.Mode(),
			Size:    info.Size(),
			ModTime: info.ModTime(),
			IsDir:   entry.IsDir(),
		}

		if options.Long {
			stat := info.Sys().(*syscall.Stat_t)
			fileInfo.Links = int(stat.Nlink)
			fileInfo.User = getUserName(int(stat.Uid))
			fileInfo.Group = getGroupName(int(stat.Gid))
		}

		fileInfos = append(fileInfos, fileInfo)
	}

	sortEntries(fileInfos, options)

	if options.Long {
		printLongFormat(fileInfos, indent)
	} else {
		printShortFormat(fileInfos, indent)
	}

	if options.Recursive {
		for _, info := range fileInfos {
			if info.IsDir && (options.All || !strings.HasPrefix(info.Name, ".")) {
				fullPath := filepath.Join(path, info.Name)
				fmt.Printf("\n%s%s:\n", indent, fullPath)
				err := ListDir(fullPath, options, indent+"  ")
				if err != nil {
					fmt.Fprintf(os.Stderr, "my-ls: %v\n", err)
				}
			}
		}
	}

	return nil
}
