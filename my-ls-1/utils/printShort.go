package utils

import (
	"fmt"

	"my-ls-1/models"
)

func printShortFormat(entries []models.FileInfo, indent string) {
	for _, entry := range entries {
		fmt.Printf("%s%s  ", indent, entry.Name)
	}
	fmt.Println()
}
