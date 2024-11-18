package utils

import (
	"fmt"
	"os"
)

func ReadLogo() []byte {
	file, err := os.ReadFile("linuxlogo.txt")
	if err != nil {
		fmt.Println("file not found", err)
	}
	file = append(file, 10)
	return file
}
