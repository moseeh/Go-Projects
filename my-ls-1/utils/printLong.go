package utils

import (
	"fmt"
	"strconv"

	"my-ls-1/models"
)

func printLongFormat(entries []models.FileInfo) {
	var totalBlocks int64
	sizeLen, linkLen, userLen, groupLen := 0, 0, 0, 0

	for _, entry := range entries {
		totalBlocks += (entry.Size + 1023) / 1024 // Round up to the nearest 1024 bytes
		size := strconv.FormatInt(entry.Size, 10)
		linkCount := strconv.Itoa(entry.Links)
		sizeLen = max(sizeLen, len(size))
		linkLen = max(linkLen, len(linkCount))
		userLen = max(userLen, len(entry.User))
		groupLen = max(groupLen, len(entry.Group))
	}

	fmt.Printf("total %d\n", totalBlocks)

	for _, entry := range entries {
		modeStr := entry.Mode.String()
		linkCount := strconv.Itoa(entry.Links)
		size := strconv.FormatInt(entry.Size, 10)
		timeStr := entry.ModTime.Format("Jan _2 15:04")
		color := getFileColor(entry.Mode, entry.Name)

		fmt.Printf("%-11s %*s %-*s %-*s %*s %s %s%s\033[0m\n",
			modeStr, linkLen, linkCount, userLen, entry.User, groupLen, entry.Group,
			sizeLen, size, timeStr, color, entry.Name)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
