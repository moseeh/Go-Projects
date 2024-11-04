package utils

import (
	"sort"

	"lem-in/models"
)

func calcTurns(paths []models.Path, antnumber int) (int, []int) {
	lengths := []int{}
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i].Rooms) < len(paths[j].Rooms)
	})
	for _, v := range paths {
		lengths = append(lengths, len(v.Rooms[1:]))
	}
	turns := 0
	antDistribution := distributeAnts(lengths, antnumber)

	for i := 0; i <= len(antDistribution)-1; i++ {
		n := antDistribution[i] + lengths[i] - 1
		if n > turns {
			turns = n
		}
	}

	return turns, antDistribution
}
