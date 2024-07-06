package utils

// PlaceTetromino attempts to place all tetrominoes on the board
// It returns the solved board if successful, or nil if impossible
func PlaceTetromino(board [][]string, tetrominoes [][]string) [][]string {
	if placeTetrominoRecursive(board, tetrominoes, 0) {
		return board
	}
	return nil
}

// placeTetrominoRecursive is a helper function that recursively tries to place tetrominoes
// It returns true if all tetrominoes are successfully placed, false otherwise
func placeTetrominoRecursive(board [][]string, tetrominoes [][]string, index int) bool {
	// Base case: all tetrominoes have been placed
	if index == len(tetrominoes) {
		return true
	}

	tetromino := tetrominoes[index]
	// Try to place the current tetromino at every possible position
	for y := range board {
		for x := range board[y] {
			if canPlace(board, tetromino, x, y) {
				placeTetromino(board, tetromino, x, y)
				// Recursively try to place the next tetromino
				if placeTetrominoRecursive(board, tetrominoes, index+1) {
					return true
				}
				// If unsuccessful, remove the tetromino and try the next position
				removeTetromino(board, tetromino, x, y)
			}
		}
	}
	return false
}

// canPlace checks if a tetromino can be placed at a given position on the board
func canPlace(board [][]string, tetromino []string, x, y int) bool {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				newX, newY := x+dx, y+dy
				// Check if the position is within the board and empty
				if newY >= len(board) || newX >= len(board[0]) || board[newY][newX] != "." {
					return false
				}
			}
		}
	}
	return true
}

// placeTetromino places a tetromino on the board at a given position
func placeTetromino(board [][]string, tetromino []string, x, y int) {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = string(char)
			}
		}
	}
}

// removeTetromino removes a tetromino from the board at a given position
func removeTetromino(board [][]string, tetromino []string, x, y int) {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = "."
			}
		}
	}
}
