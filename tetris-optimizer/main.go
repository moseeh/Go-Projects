package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"tetris/utils"
)

func main() {
	timeStarted := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <tetromino_file>")
		return
	}

	tetrominoes, err := utils.ReadParseTetromino(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	if !utils.ValidateTetrominoSize(tetrominoes) || !utils.IsTetrominoValidShape(tetrominoes) {
		fmt.Println("ERROR")
		return
	}

	tetrominoes = utils.TrimTetromino(tetrominoes)

	// Initialize the smallest board to start with
	boardSize := int(math.Ceil(math.Sqrt(float64(len(tetrominoes) * 4))))
	var solvedBoard [][]string
	for {
		board := utils.CreateInitialBoard(boardSize)
		solvedBoard = utils.PlaceTetromino(board, tetrominoes)
		if solvedBoard != nil {
			break
		}
		boardSize++
	}
	fmt.Println("\n⌛ Time took to solve: ", time.Since(timeStarted))

	utils.PrintBoard(solvedBoard)
}
