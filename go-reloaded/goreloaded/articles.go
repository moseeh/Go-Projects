package goreloaded

import "strings"

func Articles(s string) string {
	vowels := []string{"a", "e", "i", "o", "u", "h"}
	wordsInS := strings.Fields(s)

	for i, v := range wordsInS {
		for _, vowel := range vowels {
			if v == "a" && string(wordsInS[i+1][0]) == vowel {
				v = "an"
			}
			if v == "A" && string(wordsInS[i+1][0]) == vowel {
				v = "An"
			}
		}
	}
	newS := strings.Join(wordsInS, " ")
	return newS
}
