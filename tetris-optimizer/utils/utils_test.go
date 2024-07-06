package utils_test

import (
	"reflect"
	"testing"
	"tetris/utils"
)

// Test function for IsTetrominoValidShape from utils package
func TestIsTetrominoValidShape(t *testing.T) {
	t.Run("Should return false when tetromino is not valid", func(t *testing.T) {
		tetromino := []string{
			"....",
			"....",
			"....",
			"....",
		}
		tetrominos := [][]string{tetromino}

		result := utils.IsTetrominoValidShape(tetrominos)

		if result != false {
			t.Errorf("Expected false, got %v", result)
		}
	})

	t.Run("Should return true when tetromino is valid", func(t *testing.T) {
		tetromino := []string{
			"A...",
			"AA..",
			"A...",
			"....",
		}
		tetrominos := [][]string{tetromino}

		result := utils.IsTetrominoValidShape(tetrominos)

		if result != true {
			t.Errorf("Expected true, got %v", result)
		}
	})
}

// Test function for TrimTetromino from utils package
func TestTrimTetromino(t *testing.T) {
	t.Run("Should return tetromino without unused lines", func(t *testing.T) {
		tetromino := []string{
			"A...",
			"AA..",
			"A...",
			"....",
		}
		tetrominos := [][]string{tetromino}

		result := utils.TrimTetromino(tetrominos)

		expectedResult := []string{
			"A.",
			"AA",
			"A.",
		}

		if !reflect.DeepEqual(result[0], expectedResult) {
			t.Errorf("Expected %v, got %v", expectedResult, result[0])
		}
	})
}

// Test function for the Resolve logic
func TestPlaceTetromino(t *testing.T) {
	t.Run("Should return solved board", func(t *testing.T) {
		tetromino := []string{
			"A...",
			"AA..",
			"A...",
			"....",
		}
		tetrominos := [][]string{tetromino}
		boardSize := 4

		board := make([][]string, boardSize)
		for i := range board {
			board[i] = make([]string, boardSize)
			for j := range board[i] {
				board[i][j] = "."
			}
		}

		result := utils.PlaceTetromino(board, tetrominos)

		expectedResult := [][]string{
			{"A", ".", ".", "."},
			{"A", "A", ".", "."},
			{"A", ".", ".", "."},
			{".", ".", ".", "."},
		}

		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("Expected %v, got %v", expectedResult, result)
		}
	})
}

func TestRemoveCharAtIndex(t *testing.T) {
	tests := []struct {
		input    string
		index    int
		expected string
	}{
		{"hello", 2, "helo"},
		{"world", 0, "orld"},
		{"go", 1, "g"},
		{"abc", 3, "abc"},  // index out of range
		{"abc", -1, "abc"}, // index out of range
		{"", 0, ""},        // empty string
		{"a", 0, ""},       // single character string
	}

	for _, test := range tests {
		result := utils.RemoveCharAtIndex(test.input, test.index)
		if result != test.expected {
			t.Errorf("removeCharAtIndex(%q, %d) = %q; expected %q", test.input, test.index, result, test.expected)
		}
	}
}
