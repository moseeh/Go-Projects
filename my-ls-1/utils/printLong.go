package utils

import (
	"fmt"
	"strconv"

	"my-ls-1/models"
)

func printLongFormat(entries []models.FileInfo, indent string) {
	var totalBlocks int64
	sizeLen, linkLen := 0, 0
	for _, entry := range entries {
		totalBlocks += (entry.Size + 511) / 512
		size := strconv.FormatInt(entry.Size, 10)
		if sizeLen < len(size) {
			sizeLen = len(size)
		}
		linkCount := strconv.Itoa(entry.Links)
		if linkLen < len(linkCount) {
			linkLen = len(linkCount)
		}
	}
	fmt.Printf("%stotal %d\n", indent, totalBlocks)
	for _, entry := range entries {
		modeStr := entry.Mode.String()
		linkCount := strconv.Itoa(entry.Links)
		size := strconv.FormatInt(entry.Size, 10)
		timeStr := entry.ModTime.Format("Jan _2 15:04")
		if entry.IsDir {
			fmt.Printf("%s%s %*s %-8s %s %*s %s \033[38;2;0;120;255m%s\033[0m\n",
				indent, modeStr, linkLen, linkCount, entry.User, entry.Group, sizeLen, size, timeStr, entry.Name)
		} else {
			fmt.Printf("%s%s %*s %-8s %s %*s %s %s\n",
				indent, modeStr, linkLen, linkCount, entry.User, entry.Group, sizeLen, size, timeStr, entry.Name)
		}

	}
}
