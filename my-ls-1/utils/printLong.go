package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"my-ls-1/models"
)

func printLongFormat(entries []models.FileInfo, path string, files bool) {
	var totalBlocks int64
	sizeLen, linkLen, userLen, groupLen, majorLen, minorLen := 0, 0, 0, 0, 0, 0

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
			a := len(strconv.Itoa(int(major))) + len(strconv.Itoa(int(minor))) + 2
			if a >= sizeLen {
				sizeLen = a
				majorLen = len(strconv.Itoa(int(major)))
				minorLen = len(strconv.Itoa(int(minor)))
			}
		}
	}

	if !files {
		fmt.Printf("total %d\n", totalBlocks)
	}

	// Second pass: print the actual entries
	for _, entry := range entries {
		modeStr := modeToString(entry.Mode)
		linkCount := strconv.Itoa(entry.Links)
		size := strconv.FormatInt(entry.Size, 10)
		timeStr := formatTime(entry.ModTime)
		color := getFileColor(entry.Mode, entry.Name)

		if entry.Mode&os.ModeCharDevice != 0 || entry.Mode&os.ModeDevice != 0 {
			stat := getDeviceStat(path + "/" + entry.Name)
			major, minor := majorMinor(stat.Rdev)
			fmt.Printf("%-10s %*s %-*s %-*s %*d, %*d %s %s%s\033[0m",
				modeStr, linkLen, linkCount, userLen, entry.User, groupLen, entry.Group,
				majorLen, major, minorLen, minor, timeStr, color, entry.Name)
		} else {
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

func formatTime(t time.Time) string {
	now := time.Now()
	sixMonthsAgo := now.AddDate(0, -6, 0)

	if t.After(sixMonthsAgo) {
		return t.Format("Jan _2 15:04")
	} else {
		return t.Format("Jan _2  2006")
	}
}
