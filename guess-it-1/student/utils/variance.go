package utils

// calculating the variance
func CalcVariance(arr []float64) float64 {
	average := CalcAverage(arr)
	var sum float64
	if len(arr) == 0 {
		return 0
	}
	for _, v := range arr {
		vfloat := float64(v)
		vfloat = (vfloat - average)
		vfloat = vfloat * vfloat
		sum += vfloat
	}
	return sum / float64(len(arr))
}
