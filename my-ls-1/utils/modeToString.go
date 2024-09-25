package utils

import "os"

func modeToString(mode os.FileMode) string {
	var str string

	// Check file type
	if mode.IsDir() {
		str += "d"
	} else if mode.IsRegular() {
		str += "-"
	} else if mode&os.ModeSymlink != 0 {
		str += "l"
	} else if mode&os.ModeDevice != 0 {
		if mode&os.ModeCharDevice != 0 {
			str += "c" // Character device
		} else {
			str += "b" // Block device
		}
	} else {
		str += "?"
	}

	// Append permissions (rwx)
	const rwx = "rwxrwxrwx"
	for i, c := range rwx {
		if mode&(1<<uint(9-1-i)) != 0 {
			str += string(c)
		} else {
			str += "-"
		}
	}

	return str
}
