package utils

import (
	"fmt"
	"os"
	"strings"
)

func ParseFile(filename string) []float64 {
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("there's an error reading the file")
		os.Exit(0)
	}
	dataString := string(dataByte)
	dataSlice := strings.Split(dataString, "\n")
	dataNum := ConvToFloat(dataSlice)
	return dataNum
}
