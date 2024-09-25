package utils

import (
	"fmt"
	"os"
	"strconv"

	"my-ls-1/models"
)

func printLongFormat(entries []models.FileInfo, path string) {
	var totalBlocks int64
	sizeLen, linkLen, userLen, groupLen := 0, 0, 0, 0

	for _, entry := range entries {
		totalBlocks += (entry.Size + 4095) / 4096 * 4
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

		fmt.Printf("%-11s %*s %-*s %-*s %*s %s %s%s\033[0m",
			modeStr, linkLen, linkCount, userLen, entry.User, groupLen, entry.Group,
			sizeLen, size, timeStr, color, entry.Name)

		if entry.Mode&os.ModeSymlink != 0 {
			target, err := os.Readlink(path + "/" + entry.Name)
			if err == nil {
				targetColor := ""
				info, err := os.Stat(target)
				if err == nil {
					mod := info.Mode()
					targetColor = getFileColor(mod, target)
				}
				fmt.Printf(" -> %s%s\033[0m", targetColor, target)
			}

		}

		fmt.Println()
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
