package main

import (
	"fmt"
	"groupie-tracker/utils"
)

func main() {
	apiIndex := utils.GetApiIndex()

	Artists := utils.GetArtists(apiIndex.Artists)
	Locations, err := utils.GetLocations(apiIndex.Locations)
	if err != nil {
		fmt.Println(err)
		return
	}
	Dates, err1 := utils.GetDates(apiIndex.Dates)
	if err1 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Artists[0].CreationDate)
	fmt.Println(Locations[0])
	fmt.Println(Dates[0])
}
