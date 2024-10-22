package utils

import (
	"fmt"
	"math"

	"my-ls-1/models"
)

// printShortFormat prints the entries in a short format with columns
func printShortFormat(entries []models.FileInfo) {
	// Use a fixed terminal width of 80 characters

	const termWidth = 205

	// Function to calculate total width given column widths
	totalWidth := func(widths []int) int {
		sum := 0
		for _, w := range widths {
			sum += w + 2 // Add 2 for spacing
		}
		return sum
	}

	// Binary search to find the optimal number of columns
	left, right := 1, len(entries)
	optimalCols := 1
	for left <= right {
		mid := (left + right) / 2
		numRows := int(math.Ceil(float64(len(entries)) / float64(mid)))

		// Calculate column widths for this configuration
		colWidths := make([]int, mid)
		for i, entry := range entries {
			col := i / numRows
			colWidths[col] = int(math.Max(float64(colWidths[col]), float64(len(entry.Name))))
		}

		if totalWidth(colWidths) <= termWidth {
			optimalCols = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// Recalculate final layout with the optimal number of columns
	numRows := int(math.Ceil(float64(len(entries)) / float64(optimalCols)))
	colWidths := make([]int, optimalCols)
	for i, entry := range entries {
		col := i / numRows
		colWidths[col] = int(math.Max(float64(colWidths[col]), float64(len(entry.Name))))
	}

	// Print entries in vertical columns with proper alignment
	for row := 0; row < numRows; row++ {
		for col := 0; col < optimalCols; col++ {
			index := col*numRows + row
			if index < len(entries) {
				entry := entries[index]

				// Get the color for the current file
				color := getFileColor(entry.Mode, entry.Name)

				// Print the file name with the corresponding color
				fmt.Printf("%s%-*s\033[0m", color, colWidths[col], entry.Name)

				// Add spacing between columns
				fmt.Printf("  ") // Two spaces between columns for better readability
			}
		}
		fmt.Println() // Move to the next row after printing a full row
	}
}
