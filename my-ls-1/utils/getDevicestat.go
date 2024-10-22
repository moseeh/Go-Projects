package utils

import (
	"fmt"
	"syscall"
)

// getDeviceStat retrieves the device statistics for the specified file path.
// It returns a pointer to a `syscall.Stat_t` structure containing the file's metadata.
func getDeviceStat(filePath string) *syscall.Stat_t {
	stat := &syscall.Stat_t{}
	err := syscall.Lstat(filePath, stat)
	if err != nil {
		fmt.Printf("Error getting device stat: %v\n", err)
	}
	return stat
}

// majorMinor extracts the major and minor device numbers from the raw device ID (rdev).
// It returns the major and minor numbers separately for easier processing and display.
func majorMinor(rdev uint64) (uint64, uint64) {
	major := (rdev >> 8) & 0xfff
	minor := (rdev & 0xff) | ((rdev >> 12) & 0xfff00)
	return major, minor
}
