package utils

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"

	"my-ls-1/models"
)

// Winsize struct to store terminal window size
type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// GetTerminalWidth returns the width of the terminal in characters
func GetTerminalWidth() (int, error) {
	ws := &Winsize{}
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
	if err != 0 {
		return 0, err
	}
	return int(ws.Col), nil
}

// Function to determine if a file is executable
func isExecutable(mode os.FileMode) bool {
	return mode&0o111 != 0
}

// Function to check if a file is an image based on extension
func isImage(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg"}
	for _, imgExt := range imageExtensions {
		if ext == imgExt {
			return true
		}
	}
	return false
}

func printShortFormat(entries []models.FileInfo, indent string) {
	// Get the terminal width
	termWidth, err := GetTerminalWidth()
	if err != nil {
		termWidth = 80
	}

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

	// Recalculate final layout with optimal number of columns
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
				if entry.IsDir {
					// Blue for directories
					fmt.Printf("%s\033[38;2;0;120;255m%-*s\033[0m", indent, colWidths[col]+2, entry.Name)
				} else if isImage(entry.Name) {
					// Magenta for image files
					fmt.Printf("%s\033[38;2;255;0;255m%-*s\033[0m", indent, colWidths[col]+2, entry.Name)
				} else if isExecutable(entry.Mode) {
					// Green for executable files
					fmt.Printf("%s\033[38;2;0;255;0m%-*s\033[0m", indent, colWidths[col]+2, entry.Name)
				} else {
					// Default color for other files
					fmt.Printf("%s%-*s", indent, colWidths[col]+2, entry.Name)
				}
			}
		}
		fmt.Println()
	}
}
