package utils

func distributeAnts(pathLengths []int, numAnts int) []int {
	n := len(pathLengths)
	if n == 0 || numAnts == 0 {
		return make([]int, n)
	}

	// Initialize distribution array
	distribution := make([]int, n)
	remaining := numAnts

	for remaining > 0 {
		// Find minimum total time across all paths
		minTotal := pathLengths[0] + distribution[0]
		for i := 1; i < n; i++ {
			total := pathLengths[i] + distribution[i]
			if total < minTotal {
				minTotal = total
			}
		}

		// Find shortest path among those with minimum total time
		shortestPathIndex := 0
		for i := 0; i < n; i++ {
			total := pathLengths[i] + distribution[i]
			if total == minTotal {
				shortestPathIndex = i
				break // Since paths are sorted, first match is shortest
			}
		}

		// Add ant to the chosen path
		distribution[shortestPathIndex]++
		remaining--
	}

	return distribution
}
