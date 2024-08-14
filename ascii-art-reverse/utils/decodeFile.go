package utils

// decodes file passed to create the string equivalent to the art
func DecodeFile(lines []string, artToChar map[string]string) string {
	start := 0
	s := ""
	for start < len(lines) {
		if lines[start] != "" {
			if start+7 < len(lines) {
				s += decoder(lines[start:start+8], artToChar)
				if s == "" {
					return ""
				} else {
					s += "\n"
				}
			}
			start += 8
		} else {
			s += "\n"
			start++
		}
	}
	return s
}
