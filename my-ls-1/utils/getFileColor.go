package utils

import (
	"os"
	"strings"
)

var lsColors map[string]string

func init() {
	lsColors = parseLSColors(os.Getenv("LS_COLORS"))
}

func parseLSColors(lsColorsEnv string) map[string]string {
	colors := make(map[string]string)
	pairs := strings.Split(lsColorsEnv, ":")
	for _, pair := range pairs {
		parts := strings.Split(pair, "=")
		if len(parts) == 2 {
			colors[parts[0]] = "\033[" + parts[1] + "m"
		}
	}
	return colors
}

func getFileColor(mode os.FileMode, fileName string) string {
	switch {
	case mode&os.ModeDir != 0:
		if mode&0o002 != 0 && mode&0o010 != 0 {
			return lsColors["tw"] // Directory, writable by others, with sticky bit
		}
		if mode&0o002 != 0 {
			return lsColors["ow"] // Directory, writable by others
		}
		return lsColors["di"] // Directory
	case mode&os.ModeSymlink != 0:
		return lsColors["ln"] // Symlink
	case mode&os.ModeNamedPipe != 0:
		return lsColors["pi"] // Named pipe
	case mode&os.ModeSocket != 0:
		return lsColors["so"] // Socket
	case mode&os.ModeDevice != 0:
		return lsColors["bd"] // Block device
	case mode&os.ModeCharDevice != 0:
		return lsColors["cd"] // Character device
	case mode&os.ModeSetuid != 0:
		return lsColors["su"] // Setuid
	case mode&os.ModeSetgid != 0:
		return lsColors["sg"] // Setgid
	case mode&0o111 != 0:
		return lsColors["ex"] // Executable
	default:
		return getColorByExtension(strings.ToLower(getFileExtension(fileName)))
	}
}

func getColorByExtension(ext string) string {
	if color, ok := lsColors["*."+ext]; ok {
		return color
	}
	return lsColors["rs"] // Default color
}

func getFileExtension(name string) string {
	parts := strings.Split(name, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}
