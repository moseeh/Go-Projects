package utils

import "strings"

// GetBannerFile takes a slice of strings and returns a modified slice and a banner file path
func GetBannerFile(args []string) ([]string, string) {
	if len(args) == 0 {
		return nil, ""
	}

	switch len(args) {
	case 1:
		return args, "banners/standard.txt"
	case 2:
		banner := strings.ToLower(args[1])
		if banner == "standard" || banner == "thinkertoy" || banner == "shadow" {
			return []string{args[0]}, "banners/" + banner + ".txt"
		}
		return args, "banners/standard.txt"
	case 3:
		return args[:2], "banners/" + strings.ToLower(args[2]) + ".txt"
	default:
		return args, "banners/standard.txt"
	}
}
