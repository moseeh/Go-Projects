package main

import (
	"fmt"
	"math"
	utils "mathSkills/utils"
	"os"
	"strings"
)

func main() {
	// taking the arguments passed
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		fmt.Println("please provide file to read")
		return
	}
	// reading the file
	if utils.IsTxt(arguments[0]) {
		dataByte, err := os.ReadFile(arguments[0])
		if err != nil {
			fmt.Println("there's an error reading the file")
			return
		}
		// converting the data to bytes
		dataString := string(dataByte)
		dataSlice := strings.Split(dataString, "\n")
		dataNum := utils.ConvToInt(dataSlice)
		fmt.Printf("Average: %.0f\n", math.Round(utils.CalcAverage(dataNum)))
		fmt.Printf("Median: %.0f\n", math.Round(utils.CalcMedian(dataNum)))
		fmt.Printf("Variance: %.0f\n", math.Round(utils.CalcVariance(dataNum)))
		fmt.Printf("Standard Deviation: %.0f\n", math.Round(utils.CalcStandardDev(utils.CalcVariance(dataNum))))

	} else {
		fmt.Println("Please provide a .txt file")
	}

}
