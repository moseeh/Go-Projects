package main

import (
	"fmt"
	"os"
	"strings"

	goreloaded "go-reloaded/goreloaded"
)

func main() {
	toRead := os.Args[1]
	toWrite := os.Args[2]
	myTextByte, _ := os.ReadFile(toRead)
	theText := string(myTextByte)
	artText := goreloaded.Articles(theText)
	punctText := goreloaded.PunctuateSentence(artText)
	qoutText := goreloaded.QuotationMarks(punctText)
	modSentence := goreloaded.Modifications(qoutText)

	punctContent := goreloaded.PunctuateSentence(modSentence)
	contentToWrite2 := strings.Join(strings.Fields(punctContent), " ")
	err := os.WriteFile(toWrite, []byte(contentToWrite2), 0o644)
	if err != nil {
		fmt.Println("there's an error")
	}
}
