package utils

import (
	"fmt"
	"strconv"
)

// converting a slice of string to a slice of int
func ConvToInt(sArr []string) []float64 {
	if len(sArr) == 0 {
		return nil
	}
	nArr := []float64{}
	for _, v := range sArr {
		num, err := strconv.ParseFloat(v, 32)
		if err != nil {
			fmt.Println("error converting number", v)
		} else {
			nArr = append(nArr, num)
		}
	}
	return nArr
}
