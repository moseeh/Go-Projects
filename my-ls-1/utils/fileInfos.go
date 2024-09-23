package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"my-ls-1/models"
)

func getFileInfos(path string, entries []fs.DirEntry, options models.Options) []models.FileInfo {
	fileInfos := make([]models.FileInfo, 0, len(entries))

	addSpecialEntry := func(name string) {
		info, err := os.Stat(filepath.Join(path, name))
		if err == nil {
			fileInfos = append(fileInfos, getFileInfo(name, info, options))
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
			continue
		}

		fileInfos = append(fileInfos, getFileInfo(entry.Name(), info, options))
	}

	return fileInfos
}

func getFileInfo(name string, info os.FileInfo, options models.Options) models.FileInfo {
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

	return fileInfo
}
