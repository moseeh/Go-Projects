package utils

import (
	"strings"

	"my-ls-1/models"
)

// sort entries of a directory using the specified flags
func sortEntries(entries []models.FileInfo, options models.Options) {
	n := len(entries)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if shouldSwap(entries[j], entries[j+1], options) {
				entries[j], entries[j+1] = entries[j+1], entries[j]
			}
		}
	}

	// Reverse the slice if the reverse option is enabled
	if options.Reverse {
		for i := 0; i < n/2; i++ {
			entries[i], entries[n-1-i] = entries[n-1-i], entries[i]
		}
	}
}

func shouldSwap(a, b models.FileInfo, options models.Options) bool {
	// Always keep . and .. at the top, in that order
	if a.Name == "." && b.Name == ".." {
		return false
	}
	if b.Name == ".." && a.Name == "." {
		return true
	}
	// Sort by modification time if the option is enabled
	if options.SortByTime {
		if !a.ModTime.Equal(b.ModTime) {
			return a.ModTime.Before(b.ModTime)
		}
	}

	// Use ls-like lexicographical order
	return lsLess(a.Name, b.Name)
}

func lsLess(a, b string) bool {
	// Helper function to get normalized strings without non-alphanumeric characters.
	normalize := func(s string) string {
		start := 1
		if s[0] != '.' { // Skip the first character only if it's a dot.
			start = 0
		}
		var builder strings.Builder
		for _, v := range s[start:] {
			if ('a' <= v && v <= 'z') || ('0' <= v && v <= '9') {
				builder.WriteRune(v)
			} else if 'A' <= v && v <= 'Z' {
				// Convert uppercase to lowercase
				builder.WriteRune(v + 32)
			}
		}
		return builder.String()
	}

	// Normalize both strings for comparison.
	aComp := normalize(a)
	bComp := normalize(b)

	if aComp != bComp {
		return aComp > bComp
	}

	// If normalized names are equal, fall back to the original case-sensitive comparison.
	return a > b
}
