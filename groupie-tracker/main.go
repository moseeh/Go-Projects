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
	Relations, err2 := utils.GetRelations(apiIndex.Relation)
	fmt.Println(apiIndex.Relation)
	if err2 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Artists[0].CreationDate)
	fmt.Println(Locations[0])
	fmt.Println(Dates[0])
	// Print out relations data
	for _, relation := range Relations {
		fmt.Printf("ID: %d\n", relation.ID)
		for location, dates := range relation.DatesLocations {
			fmt.Printf("  Location: %s\n", location)
			fmt.Printf("  Dates: %v\n", dates)
		}
		fmt.Println("------------------------")
	}
}
