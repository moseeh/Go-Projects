package utils

import (
	"fmt"
	"strconv"

	"my-ls-1/models"
)

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
		if entry.IsDir {
			fmt.Printf("%s%s %3s %-8s %-8s %8s %s \033[38;2;0;120;255m%s\033[0m\n",
				indent, modeStr, linkCount, entry.User, entry.Group, size, timeStr, entry.Name)
		} else {
			fmt.Printf("%s%s %3s %-8s %-8s %8s %s %s\n",
				indent, modeStr, linkCount, entry.User, entry.Group, size, timeStr, entry.Name)
		}

	}
}
