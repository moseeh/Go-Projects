package utils

// IsTetrominoValidShape checks if tetrominoes have valid shapes and has 4 cubes
func IsTetrominoValidShape(tetrominos [][]string) bool {
	for _, tetromino := range tetrominos {
		blockCount, connectionCount := 0, 0
		for i, row := range tetromino { // i is the horizontal index
			for j, char := range row { // j is the vertical index
				if char >= 'A' && char <= 'Z' {
					blockCount++
					if i > 0 && tetromino[i-1][j] == byte(char) {
						connectionCount++
					}
					if i < len(tetromino)-1 && tetromino[i+1][j] == byte(char) {
						connectionCount++
					}
					if j > 0 && tetromino[i][j-1] == byte(char) {
						connectionCount++
					}
					if j < len(row)-1 && tetromino[i][j+1] == byte(char) {
						connectionCount++
					}
				}
			}
		}
		if blockCount != 4 || connectionCount < 6 {
			return false
		}
	}
	return true
}

// ValidateTetrominoSize checks if tetrominoes are of valid size (4x4)
func ValidateTetrominoSize(tetrominos [][]string) bool {
	for _, tetromino := range tetrominos {
		if len(tetromino) != 4 {
			return false
		}
		for _, line := range tetromino {
			if len(line) != 4 {
				return false
			}
		}
	}
	return true
}
