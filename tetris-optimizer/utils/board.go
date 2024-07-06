package utils

import (
	"fmt"
)

// CreateInitialBoard creates an empty board of the given size
func CreateInitialBoard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

// PrintBoard prints the board to the console
func PrintBoard(board [][]string) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Println()
}
