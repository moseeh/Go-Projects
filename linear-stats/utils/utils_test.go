package utils_test

import (
	"math"
	"testing"

	"linear-stats/utils"
)

func TestIsTxt(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{
			name:     "valid txt file",
			filename: "document.txt",
			want:     true,
		},
		{
			name:     "uppercase extension",
			filename: "document.TXT",
			want:     false,
		},
		{
			name:     "no extension",
			filename: "document",
			want:     false,
		},
		{
			name:     "different extension",
			filename: "document.pdf",
			want:     false,
		},
		{
			name:     "multiple dots",
			filename: "my.file.txt",
			want:     true,
		},
		{
			name:     "empty string",
			filename: "",
			want:     false,
		},
		{
			name:     "only extension",
			filename: ".txt",
			want:     true,
		},
		{
			name:     "path with txt file",
			filename: "/path/to/file.txt",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsTxt(tt.filename); got != tt.want {
				t.Errorf("IsTxt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatePearson(t *testing.T) {
	// Helper function to check if two float64 values are approximately equal
	equalWithinTolerance := func(a, b float64) bool {
		epsilon := 1e-10
		diff := math.Abs(a - b)
		return diff < epsilon
	}

	tests := []struct {
		name string
		data []float64
		want float64
	}{
		{
			name: "perfect positive correlation",
			data: []float64{1, 2, 3, 4, 5},
			want: 1.0,
		},
		{
			name: "perfect negative correlation",
			data: []float64{5, 4, 3, 2, 1},
			want: -1.0,
		},
		{
			name: "no correlation",
			data: []float64{1, 1, 1, 1, 1},
			want: math.NaN(), // Should return NaN for zero variance
		},
		{
			name: "moderate positive correlation",
			data: []float64{1, 2, 2, 3, 4},
			want: 0.970725343394151,
		},
		{
			name: "moderate negative correlation",
			data: []float64{4, 3, 2, 2, 1},
			want: -0.970725343394151,
		},
		{
			name: "real world example",
			data: []float64{68.7, 71.2, 74.3, 77.1, 78.9},
			want: 0.9966698341336411,
		},
		{
			name: "single element",
			data: []float64{1},
			want: math.NaN(), // Cannot calculate correlation for single point
		},
		{
			name: "two elements",
			data: []float64{1, 2},
			want: 1.0,
		},
		{
			name: "floating point precision test",
			data: []float64{1.23456, 2.34567, 3.45678, 4.56789, 5.67890},
			want: 0.9999999998379941,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.CalculatePearson(tt.data)

			// Special handling for NaN cases
			if math.IsNaN(tt.want) {
				if !math.IsNaN(got) {
					t.Errorf("CalculatePearson() = %v, want NaN", got)
				}
				return
			}

			if !equalWithinTolerance(got, tt.want) {
				t.Errorf("CalculatePearson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateLinearRegression(t *testing.T) {
	// Helper function to check if two float64 values are approximately equal
	equalWithinTolerance := func(a, b float64) bool {
		epsilon := 1e-10
		return math.Abs(a-b) < epsilon
	}

	tests := []struct {
		name          string
		data          []float64
		wantSlope     float64
		wantIntercept float64
	}{
		{
			name:          "perfect positive correlation",
			data:          []float64{1, 2, 3, 4, 5},
			wantSlope:     1.0,
			wantIntercept: 1.0,
		},
		{
			name:          "perfect negative correlation",
			data:          []float64{5, 4, 3, 2, 1},
			wantSlope:     -1.0,
			wantIntercept: 5.0,
		},
		{
			name:          "horizontal line",
			data:          []float64{2, 2, 2, 2, 2},
			wantSlope:     0.0,
			wantIntercept: 2.0,
		},
		{
			name:          "two points",
			data:          []float64{1, 2},
			wantSlope:     1.0,
			wantIntercept: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSlope, gotIntercept := utils.CalculateLinearRegression(tt.data)

			// Check if the line fits the points
			for i := range tt.data {
				x := float64(i)
				expectedY := tt.wantSlope*x + tt.wantIntercept
				actualY := gotSlope*x + gotIntercept
				if !equalWithinTolerance(expectedY, actualY) {
					t.Errorf("Point %d: got y=%v, want y=%v", i, actualY, expectedY)
				}
			}

			// For specific cases, also check the exact slope and intercept
			if math.Abs(tt.wantSlope) > 0.0001 { // Only check when we have specific expectations
				if !equalWithinTolerance(gotSlope, tt.wantSlope) {
					t.Errorf("Slope = %v, want %v", gotSlope, tt.wantSlope)
				}
				if !equalWithinTolerance(gotIntercept, tt.wantIntercept) {
					t.Errorf("Intercept = %v, want %v", gotIntercept, tt.wantIntercept)
				}
			}
		})
	}
}
