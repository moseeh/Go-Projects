package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadParseTetromino reads and parses tetrominoes from a file
func ReadParseTetromino(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var tetrominoes [][]string
	var currentTetromino []string
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if len(currentTetromino) > 0 {
				tetrominoes = append(tetrominoes, currentTetromino)
				currentTetromino = nil
			}
		} else {
			if !isValidLine(line) {
				return nil, fmt.Errorf("invalid character in tetromino at line %d", lineNum)
			}
			currentTetromino = append(currentTetromino, line)
		}
	}

	if len(currentTetromino) > 0 {
		tetrominoes = append(tetrominoes, currentTetromino)
	}

	for i, tetromino := range tetrominoes {
		letter := 'A' + rune(i)
		for j, line := range tetromino {
			tetromino[j] = strings.ReplaceAll(line, "#", string(letter))
		}
	}

	return tetrominoes, nil
}

func isValidLine(line string) bool {
	for _, char := range line {
		if char != '#' && char != '.' {
			return false
		}
	}
	return true
}
