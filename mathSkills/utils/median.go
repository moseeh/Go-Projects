package utils

import "sort"

// finding the median value
func CalcMedian(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	sort.Float64s(arr)
	n := len(arr)
	if n%2 == 0 {
		mid := n / 2
		return float64(arr[mid-1] + arr[mid]) / 2.0
	}
	return float64(arr[n/2])
}
