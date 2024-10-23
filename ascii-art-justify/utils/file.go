package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func CheckFileSize(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}
	switch filename {
	case "banners/standard.txt":
		if HashChecker(filename) != "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf" {
			return false
		}

	case "banners/shadow.txt":

		if HashChecker(filename) != "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73" {
			return false
		}
	case "banners/thinkertoy.txt":

		if HashChecker(filename) != "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3" {
			return false
		}

	}
	return true
}

// Checks the hash of a file and returns the value in string format
func HashChecker(str string) string {
	file, err := os.Open(str)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Create a SHA256 hash object
	h := sha256.New()

	// Copy the file content to the hash
	if _, err := io.Copy(h, file); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Get the hash bytes
	hashBytes := h.Sum(nil)

	// Convert the bytes to a string
	return fmt.Sprintf("%x", hashBytes)
}
