package utils

import (
	"os"
	"strings"
)

func getFileColor(mode os.FileMode, fileName string) string {
	switch {
	case mode&os.ModeDir != 0:
		if mode&0o002 != 0 && mode&0o010 != 0 {
			return "\033[42;30m" // Black on Green background (directory, writable by others, with sticky bit)
		}
		if mode&0o002 != 0 {
			return "\033[43;30m" // Black on Yellow background (directory, writable by others)
		}
		return "\033[1;34m" // Bold Blue (directory)
	case mode&os.ModeSymlink != 0:
		return "\033[1;36m" // Bold Cyan (symlink)
	case mode&os.ModeNamedPipe != 0:
		return "\033[40;33m" // Yellow on Black background (named pipe)
	case mode&os.ModeSocket != 0:
		return "\033[1;35m" // Bold Magenta (socket)
	case mode&os.ModeDevice != 0:
		return "\033[1;33;40m" // Bold Yellow on Black background (device file)
	case mode&os.ModeCharDevice != 0:
		return "\033[1;33;40m" // Bold Yellow on Black background (character device)
	case mode&os.ModeSetuid != 0:
		return "\033[37;41m" // White on Red background (setuid)
	case mode&os.ModeSetgid != 0:
		return "\033[30;46m" // Black on Cyan background (setgid)
	case mode&0o111 != 0:
		return "\033[1;32m" // Bold Green (executable)
	default:
		// Check for common file extensions
		return getColorByExtension(strings.ToLower(getFileExtension(fileName)))
	}
}

func getColorByExtension(ext string) string {
	switch ext {
	case "gz", "tar", "zip", "rar", "7z":
		return "\033[1;31m" // Bold Red (compressed file)
	case "jpg", "jpeg", "gif", "bmp", "png", "svg":
		return "\033[1;35m" // Bold Magenta (image file)
	case "mp3", "wav", "flac":
		return "\033[1;36m" // Bold Cyan (audio file)
	case "mp4", "mkv", "avi", "mov":
		return "\033[1;33m" // Bold Yellow (video file)
	case "pdf", "epub", "mobi":
		return "\033[0;33m" // Yellow (document file)
	case "sh", "bash", "zsh", "fish":
		return "\033[0;32m" // Green (shell script)
	case "py", "rb", "pl", "php":
		return "\033[0;34m" // Blue (interpreted language script)
	case "cpp", "c", "h", "go", "rs":
		return "\033[0;95m" // Light Magenta (compiled language source)
	default:
		return "\033[0m" // Default color
	}
}

func getFileExtension(name string) string {
	parts := strings.Split(name, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}
