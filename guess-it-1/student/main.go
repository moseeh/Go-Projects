package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"guess-it-1/utils"
)

func main() {
	var previousValues []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number")
			continue
		}
		previousValues = append(previousValues, number)
		if len(previousValues) > 2 {
			previousValues = previousValues[1:]
		}
		if len(previousValues) == 2 {
			mean := utils.CalcAverage(previousValues)
			stdDev := utils.CalcStandardDev(utils.CalcVariance(previousValues))
			lowerLimit := mean - 2*stdDev
			upperLimit := mean + 2*stdDev
			fmt.Printf("%.0f %.0f\n", lowerLimit, upperLimit)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
