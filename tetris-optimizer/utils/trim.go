package utils

import "strings"

// TrimTetromino removes unused lines from tetrominoes
func TrimTetromino(tetrominos [][]string) [][]string {
	var trimmedTetrominos [][]string

	for _, tetromino := range tetrominos {
		trimmed := trimHorizontal(tetromino)
		trimmed = trimVertical(trimmed)
		trimmedTetrominos = append(trimmedTetrominos, trimmed)
	}

	return trimmedTetrominos
}

//removing horizontal lines that don't contain any letters
func trimHorizontal(tetromino []string) []string {
	var result []string
	for _, line := range tetromino {
		if strings.ContainsAny(line, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			result = append(result, line)
		}
	}
	return result
}

//removing vertical lines that don't contain any letters
func trimVertical(tetromino []string) []string {
	for i := len(tetromino[0]) - 1; i >= 0; i-- {
		hasLetter := false
		for _, line := range tetromino {
			if line[i] >= 'A' && line[i] <= 'Z' {
				hasLetter = true
				break
			}
		}
		if !hasLetter {
			for j := range tetromino {
				tetromino[j] = RemoveCharAtIndex(tetromino[j], i)
			}
		}
	}
	return tetromino
}

// removing characters from specified index
func RemoveCharAtIndex(input string, index int) string {
	if index < 0 || index >= len(input) {
		return input
	}
	return input[:index] + input[index+1:]
}
