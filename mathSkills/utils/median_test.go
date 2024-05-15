package utils

import "testing"

func TestMedian(t *testing.T) {
	testCases := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"Empty Slice", []float64{}, 0},
		{"Single value", []float64{5}, 5},
		{"Positive values", []float64{1, 2, 3, 4, 5}, 3},
		{"Negatuve values", []float64{-1, -2, -3, -4, -5}, -3},
		{"Mixed values", []float64{-1, 0, 1, 2, 3}, 1},
		{"an array of even length", []float64{2, 4, 6, 8}, 5},
	}
	// compare function values and expected value

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := CalcMedian(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected: %f, Got: %f", testCase.expected, result)
			}

		})
	}
}
