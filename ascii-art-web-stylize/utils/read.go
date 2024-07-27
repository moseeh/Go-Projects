package utils

import (
	"fmt"
	"os"
)

// ReadsFile reads the file and returns its content as a string along with any error encountered
func ReadsFile(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		// Return an empty string and the error
		return "", fmt.Errorf("failed to read file %s: %v", file, err)
	}
	if !checksum(file) {
		return "", fmt.Errorf("error: The %s file has been tampered with", file)
	}
	return string(content), nil
}
