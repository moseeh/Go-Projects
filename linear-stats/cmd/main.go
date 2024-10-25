package main

import (
	"fmt"
	"os"

	"linear-stats/utils"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		fmt.Println("please provide a file to read")
		return
	}
	if !utils.IsTxt(arguments[0]) {
		fmt.Println("please provide a *.txt file")
		return
	}
	dataFloat := utils.ParseFile(arguments[0])
	slope, intercept := utils.CalculateLinearRegression(dataFloat)
	pearson := utils.CalculatePearson(dataFloat)

	// Output results in required format
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pearson)
}
