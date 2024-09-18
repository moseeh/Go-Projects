package utils

import (
	"fmt"

	"my-ls-1/models"
)

func printShortFormat(entries []models.FileInfo, indent string) {
	for _, entry := range entries {
		if entry.IsDir {
			fmt.Printf("%s\033[38;2;0;120;255m%s\033[0m  ", indent, entry.Name)
		} else {
			fmt.Printf("%s%s  ", indent, entry.Name)
		}
	}
	fmt.Println()
}