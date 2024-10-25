package utils

import "math"

// CalculatePearson computes the Pearson correlation coefficient
func CalculatePearson(data []float64) float64 {
	n := float64(len(data))

	// Create x values (indices starting from 1)
	x := make([]float64, len(data))
	for i := range x {
		x[i] = float64(i)
	}

	// Calculate sums
	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumXSquare := 0.0
	sumYSquare := 0.0

	for i := range data {
		sumX += x[i]
		sumY += data[i]
		sumXY += x[i] * data[i]
		sumXSquare += x[i] * x[i]
		sumYSquare += data[i] * data[i]
	}

	// Calculate Pearson correlation coefficient
	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt(n*sumXSquare-sumX*sumX) * math.Sqrt(n*sumYSquare-sumY*sumY)

	return numerator / denominator
}
