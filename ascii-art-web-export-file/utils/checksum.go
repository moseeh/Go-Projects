package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func checksum(fileName string) bool {
	mapfile := map[string]string{
		"banners/standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
		"banners/shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
		"banners/thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error: failed to read file %s: %v\n", fileName, err)
		return false
	}

	check := sha256.Sum256(file)
	encodedfile := hex.EncodeToString(check[:])
	return encodedfile == mapfile[fileName]
}
