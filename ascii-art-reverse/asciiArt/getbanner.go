package asciiArt

import "os"

func BannerFile() string {
	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "standard":
			return "banners/standard.txt"
		case "shadow":
			return "banners/shadow.txt"
		case "thinkertoy":
			return "banners/thinkertoy.txt"
		default:
			return ""
		}
	}
	if len(os.Args) == 2 {
		return "banners/standard.txt"
	}
	return ""
}
