package utils

// calculating average
func CalcAverage(arr []float64) float64 {
	var sum float64
	if len(arr) == 0 {
		return 0
	}
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}
