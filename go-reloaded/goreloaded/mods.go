package goreloaded

import (
	"strconv"
	"strings"
)

func Modifications(s string) string {
	wordsInText := strings.Fields(s)
	for i, v := range wordsInText {
		switch {
		case strings.Contains(v, "(hex)"):
			if i > 0 {
				hexToDec, _ := strconv.ParseInt(wordsInText[i-1], 16, 64)
				wordsInText[i-1] = strconv.Itoa(int(hexToDec))
				wordsInText[i] = ""
			}
		case strings.Contains(v, "(bin)"):
			if i > 0 {
				hexToDec, _ := strconv.ParseInt(wordsInText[i-1], 2, 64)
				wordsInText[i-1] = strconv.Itoa(int(hexToDec))
				wordsInText[i] = ""
			}
		case strings.Contains(v, "(up)"):
			wordsInText[i-1] = strings.ToUpper(wordsInText[i-1])
			wordsInText[i] = strings.Trim(wordsInText[i], "(up)")
		case strings.Contains(v, "(low)"):
			wordsInText[i-1] = strings.ToLower(wordsInText[i-1])
			wordsInText[i] = strings.Trim(wordsInText[i], "(low))")
		case strings.Contains(v, "(cap)"):
			wordsInText[i-1] = strings.Title(wordsInText[i-1])
			wordsInText[i] = strings.Trim(wordsInText[i], "(cap)")
		case strings.Contains(v, "(up,"):
			newStr := wordsInText[i+1]
			for index, value := range newStr {
				if value == ')' {
					number := string(newStr[0:index])
					convNUmber, _ := strconv.Atoi(number)
					for j := 1; j <= convNUmber; j++ {
						wordsInText[i-j] = strings.ToUpper(wordsInText[i-j])
						wordsInText[i] = ""
					}
					if index+1 == len(newStr) {
						wordsInText[i+1] = ""
					} else {
						wordsInText[i+1] = wordsInText[i+1][index+1:]
					}
				}
			}
		case strings.Contains(v, "(low,"):
			newStr := wordsInText[i+1]
			for index, value := range newStr {
				if value == ')' {
					number := string(newStr[0:index])
					convNUmber, _ := strconv.Atoi(number)
					for j := 1; j <= convNUmber; j++ {
						wordsInText[i-j] = strings.ToLower(wordsInText[i-j])
						wordsInText[i] = ""
					}
					if index+1 == len(newStr) {
						wordsInText[i+1] = ""
					} else {
						wordsInText[i+1] = wordsInText[i+1][index+1:]
					}
				}
			}
		case strings.Contains(v, "(cap,"):
			newStr := wordsInText[i+1]
			for index, value := range newStr {
				if value == ')' {
					number := string(newStr[0:index])
					convNUmber, _ := strconv.Atoi(number)
					for j := 1; j <= convNUmber; j++ {
						wordsInText[i-j] = strings.Title(wordsInText[i-j])
						wordsInText[i] = ""
					}
					if index+1 == len(newStr) {
						wordsInText[i+1] = ""
					} else {
						wordsInText[i+1] = wordsInText[i+1][index+1:]
					}
				}
			}
		}
	}
	newS := strings.Join(wordsInText, " ")
	return newS
}
