package utils

import (
	"fmt"
	"syscall"
)

func getDeviceStat(filePath string) *syscall.Stat_t {
	stat := &syscall.Stat_t{}
	err := syscall.Lstat(filePath, stat)
	if err != nil {
		fmt.Printf("Error getting device stat: %v\n", err)
	}
	return stat
}

func majorMinor(rdev uint64) (uint64, uint64) {
	major := (rdev >> 8) & 0xfff
	minor := (rdev & 0xff) | ((rdev >> 12) & 0xfff00)
	return major, minor
}
