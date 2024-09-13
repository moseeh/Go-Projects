package utils

import (
	"sort"

	"my-ls-1/models"
)

func sortEntries(entries []models.FileInfo, options models.Options) {
	sort.Slice(entries, func(i, j int) bool {
		if options.SortByTime {
			if entries[i].ModTime.Equal(entries[j].ModTime) {
				return entries[i].Name < entries[j].Name
			}
			return entries[i].ModTime.After(entries[j].ModTime)
		}
		return entries[i].Name < entries[j].Name
	})

	if options.Reverse {
		for i := len(entries)/2 - 1; i >= 0; i-- {
			opp := len(entries) - 1 - i
			entries[i], entries[opp] = entries[opp], entries[i]
		}
	}
}
