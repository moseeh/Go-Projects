package utils

import (
	"fmt"
	"os"
	"strings"
)

// This function takes a filepath and  variable holding a string literal as arguments
// It then creates the file using os package(with respect to the filepath) and writes the string into the file
// Upon encountering any problems in execution of th two operations, it returns an error
func CreateFile(FilePath string, content string) error {
	_, errr := os.Stat(FilePath)
	if os.IsNotExist(errr) {
		if strings.HasSuffix(FilePath, ".txt") {
			file, err := os.Create(FilePath)
			if err != nil {
				return fmt.Errorf("error in file creation: %w", err)
			}
			defer file.Close()
			_, err = file.WriteString(content)
			if err != nil {
				return fmt.Errorf("error in writing content: %w", err)
			}
		} else {
			return fmt.Errorf("wrong file extension for output")
		}
	}
	file, er := os.OpenFile(FilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if er != nil {
		fmt.Println("1")
		return fmt.Errorf("error opening file")

	}
	defer file.Close()
	_, er = file.WriteString(content)
	if er != nil {
		fmt.Println("2")
		return fmt.Errorf("error in writing content: %w", er)
	}
	return nil
}
