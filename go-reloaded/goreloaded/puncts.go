package goreloaded

import "strings"

func PunctuateSentence(s string) string {
	wordsInS := strings.Fields(s)
	punctuations := []string{".", ",", "!", "?", ":", ";"}
	puncts := ""

	for i, v := range wordsInS {
		for _, punct := range punctuations {
			if v == punct {
				wordsInS[i-1] += punct
				wordsInS[i] = ""
			}
			if string(v[0]) == punct && len(v) > 1 {
				for _, char := range wordsInS[i] {
					for _, punct2 := range punctuations {
						if string(char) == punct2 {
							puncts += punct2
							wordsInS[i-1] += puncts
							wordsInS[i] = strings.TrimLeft(wordsInS[i], punct)
							puncts = ""
						}
					}
				}
			}
		}
	}
	newS := strings.Join(wordsInS, " ")
	newS1 := strings.Join(strings.Fields(newS), " ")

	return newS1
}
