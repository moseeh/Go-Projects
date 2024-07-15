package utils

func IsEnglish(str string) bool {
	for _, char := range str {
		// Check if the character's ASCII value is within the range of English characters
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
