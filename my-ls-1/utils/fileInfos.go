package utils

import (
	"io/fs"
	"os"
	"strings"
	"syscall"

	"my-ls-1/models"
)

// getFileInfos retrieves file information for each directory entry in the given path.
// It optionally includes special entries like "." and ".." if the `All` option is set.
func getFileInfos(path string, entries []fs.DirEntry, options models.Options) []models.FileInfo {
	fileInfos := make([]models.FileInfo, 0, len(entries))

	addSpecialEntry := func(name string) {
		info, err := os.Stat(path + "/" + name)
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

// getFileInfo constructs a FileInfo object based on the given file name and os.FileInfo.
// If the `Long` option is enabled, it extracts additional metadata such as link count, user, and group.
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
		fileInfo.BlockSize = stat.Blocks
		fileInfo.User = getUserName(int(stat.Uid))
		fileInfo.Group = getGroupName(int(stat.Gid))
	}

	return fileInfo
}
