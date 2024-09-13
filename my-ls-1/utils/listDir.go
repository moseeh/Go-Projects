package utils

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
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

func printShortFormat(entries []models.FileInfo, indent string) {
	for _, entry := range entries {
		fmt.Printf("%s%s  ", indent, entry.Name)
	}
	fmt.Println()
}

func printLongFormat(entries []models.FileInfo, indent string) {
	var totalBlocks int64
	for _, entry := range entries {
		totalBlocks += (entry.Size + 511) / 512
	}
	fmt.Printf("%stotal %d\n", indent, totalBlocks)

	for _, entry := range entries {
		modeStr := entry.Mode.String()
		linkCount := strconv.Itoa(entry.Links)
		size := strconv.FormatInt(entry.Size, 10)
		timeStr := entry.ModTime.Format("Jan _2 15:04")
		fmt.Printf("%s%s %3s %-8s %-8s %8s %s %s\n",
			indent, modeStr, linkCount, entry.User, entry.Group, size, timeStr, entry.Name)
	}
}

func getUserName(uid int) string {
	u, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		return strconv.Itoa(uid)
	}
	return u.Username
}

func getGroupName(gid int) string {
	g, err := user.LookupGroupId(strconv.Itoa(gid))
	if err != nil {
		return strconv.Itoa(gid)
	}
	return g.Name
}
