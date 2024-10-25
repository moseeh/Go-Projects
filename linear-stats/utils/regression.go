// utils/stats.go
package utils

// CalculateLinearRegression computes the slope and y-intercept of the linear regression line
func CalculateLinearRegression(data []float64) (slope, intercept float64) {
	n := float64(len(data))

	// Create x values (indices starting from 1)
	x := make([]float64, len(data))
	for i := range x {
		x[i] = float64(i)
	}

	// Calculate means
	meanX := 0.0
	meanY := 0.0
	for i := range data {
		meanX += x[i]
		meanY += data[i]
	}
	meanX /= n
	meanY /= n

	// Calculate slope
	numerator := 0.0
	denominator := 0.0
	for i := range data {
		xDiff := x[i] - meanX
		yDiff := data[i] - meanY
		numerator += xDiff * yDiff
		denominator += xDiff * xDiff
	}

	slope = numerator / denominator
	// Calculate y-intercept using point-slope form
	intercept = meanY - slope*meanX

	return slope, intercept
}
