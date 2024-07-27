package utils

// takes second argument and returns its filepath as a string
func GetFile(banner string) string {
	// Handle cases where there are three or more arguments
	bannerFile := ""
	switch banner {
	case "standard":
		bannerFile = "banners/standard.txt"
	case "thinkertoy":
		bannerFile = "banners/thinkertoy.txt"
	case "shadow":
		bannerFile = "banners/shadow.txt"
	default:
		bannerFile = "Invalid bannerfile name"
	}

	return bannerFile
}
