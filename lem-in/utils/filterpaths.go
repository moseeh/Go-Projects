package utils

import (
	"lem-in/models"
)

func FilterPaths(paths []models.Path, colony *models.AntColony) ([]models.Path, []int) {
	antNo := colony.NumberOfAnts
	filteredpaths := []models.Path{}
	antdis := []int{}
	turns := 999999999999999999
	for _, path := range paths {
		combinations := getCombinations(path, paths)
		combturns, antDistribution := calcTurns(combinations, antNo)
		if combturns < turns {
			turns = combturns
			filteredpaths = combinations
			antdis = antDistribution
		}
	}
	return filteredpaths, antdis
}
