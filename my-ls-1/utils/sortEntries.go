package utils

import (
	"strings"

	"my-ls-1/models"
)

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

// lsLess mimics the ls command's lexicographical ordering
func lsLess(a, b string) bool {
	aDot := strings.HasPrefix(a, ".")
	bDot := strings.HasPrefix(b, ".")
	aLower, bLower := "", ""

	if aDot {
		aLower = strings.ToLower(a[1:])
	} else {
		aLower = strings.ToLower(a)
	}

	if bDot {
		bLower = strings.ToLower(b[1:])
	} else {
		bLower = strings.ToLower(b)
	}
	aComp, bComp := "", ""
	for _, v := range aLower {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			aComp += string(v)
		}
	}
	for _, v := range bLower {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			bComp += string(v)
		}
	}
	if aComp != bComp {
		return aComp > bComp
	}
	// If both names are equal lexicographically, compare the original case-sensitive names
	return a > b
}
