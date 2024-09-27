package utils

import (
	"fmt"
	"os"
	"strconv"

	"my-ls-1/models"
)

func printLongFormat(entries []models.FileInfo, path string, files bool) {
	var totalBlocks int64
	sizeLen, linkLen, userLen, groupLen := 0, 0, 0, 0
	devLen := 8 // For displaying major, minor as " 123,  456"

	// First pass: calculate the max lengths for size, link, user, group, and major/minor device numbers
	for _, entry := range entries {
		totalBlocks += (entry.Size + 4095) / 4096 * 4
		size := strconv.FormatInt(entry.Size, 10)
		linkCount := strconv.Itoa(entry.Links)
		sizeLen = max(sizeLen, len(size))
		linkLen = max(linkLen, len(linkCount))
		userLen = max(userLen, len(entry.User))
		groupLen = max(groupLen, len(entry.Group))

		// Check for device files and calculate their major/minor lengths
		if entry.Mode&os.ModeCharDevice != 0 || entry.Mode&os.ModeDevice != 0 {
			stat := getDeviceStat(path + "/" + entry.Name)
			major, minor := majorMinor(stat.Rdev)
			devLen = max(devLen, len(strconv.Itoa(int(major)))+len(strconv.Itoa(int(minor)))) // for " major, minor"
		}
	}
	// Ensure sizeLen accounts for either the normal file size or the device major/minor size
	sizeLen = max(sizeLen, devLen)

	if !files {
		fmt.Printf("total %d\n", totalBlocks)
	}

	// Second pass: print the actual entries
	for _, entry := range entries {
		modeStr := modeToString(entry.Mode)
		linkCount := strconv.Itoa(entry.Links)
		size := strconv.FormatInt(entry.Size, 10)
		timeStr := entry.ModTime.Format("Jan _2 15:04")
		color := getFileColor(entry.Mode, entry.Name)

		if entry.Mode&os.ModeCharDevice != 0 || entry.Mode&os.ModeDevice != 0 { // Check for device files
			stat := getDeviceStat(path + "/" + entry.Name)
			major, minor := majorMinor(stat.Rdev)
			// Align major/minor numbers within the same column width as file sizes
			fmt.Printf("%-10s %*s %-*s %-*s %*d, %*d %s %s%s\033[0m",
				modeStr, linkLen, linkCount, userLen, entry.User, groupLen, entry.Group,
				3, major, 3, minor, timeStr, color, entry.Name)
		} else {
			// For regular files, use the size length calculated earlier
			fmt.Printf("%-10s %*s %-*s %-*s %*s %s %s%s\033[0m",
				modeStr, linkLen, linkCount, userLen, entry.User, groupLen, entry.Group,
				sizeLen, size, timeStr, color, entry.Name)
		}

		// If it's a symlink, print the target
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
