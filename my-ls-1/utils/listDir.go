package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"my-ls-1/models"
)

func ListDir(path string, options models.Options) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		// If it's a file, just print its name and return
		fmt.Println(path)
		return nil
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	fileInfos := make([]models.FileInfo, 0, len(entries))
	addSpecialEntry := func(name string) {
		info, err := os.Stat(filepath.Join(path, name))
		if err == nil {
			fileInfo := models.FileInfo{
				Name:    name,
				Mode:    info.Mode(),
				Size:    info.Size(),
				ModTime: info.ModTime(),
				IsDir:   info.IsDir(),
			}
			if options.Long {
				stat := info.Sys().(*syscall.Stat_t)
				fileInfo.Links = int(stat.Nlink)
				fileInfo.User = getUserName(int(stat.Uid))
				fileInfo.Group = getGroupName(int(stat.Gid))
			}
			fileInfos = append(fileInfos, fileInfo)
		}
	}
	if options.All {
		addSpecialEntry(".")
		addSpecialEntry("..")
	}

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
		printLongFormat(fileInfos)
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
