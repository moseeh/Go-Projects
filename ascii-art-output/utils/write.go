package utils

import (
	"fmt"
	"os"
)

func WriteFile(s string, banner string) {
	file, err := os.Create(banner)
	if err != nil {
		fmt.Println("os error in file ceration")
		os.Exit(1)
	}
	defer file.Close()
	_, er := file.WriteString(s)
	if er != nil {
		fmt.Println("err in writing output content")
		os.Exit(1)
	}
}
