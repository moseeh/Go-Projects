package utils

import (
	"slices"

	"lem-in/models"
)

func getCombinations(path models.Path, paths []models.Path) []models.Path {
	combinations := []models.Path{path}
	if len(path.Rooms) == 2 {
		return combinations
	}

outer:
	for _, v := range paths {
		if hasCommonElements(v.Rooms[1:len(v.Rooms)-1], path.Rooms[1:len(path.Rooms)-1]) {
			continue
		}
		for _, comb := range combinations {
			if hasCommonElements(comb.Rooms[1:len(comb.Rooms)-1], v.Rooms[1:len(v.Rooms)-1]) {
				continue outer
			}
		}
		combinations = append(combinations, v)
	}
	return combinations
}

func hasCommonElements(slice1, slice2 []string) bool {
	for _, v := range slice2 {
		if slices.Contains(slice1, v) {
			return true
		}
	}
	return false
}
