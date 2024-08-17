package handlers

import (
	"fmt"
	"groupie-tracker/utils"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/locations/"):])
	locations, err := utils.GetLocations(utils.GetApiIndex().Locations + "/" + strconv.Itoa(id))
	if err != nil {
		fmt.Println(err)
	}
	relations, err := utils.GetRelations(utils.GetApiIndex().Relation + "/" + strconv.Itoa((id)))
	if err != nil {
		fmt.Println(err)
	}
	dates, err1 := utils.GetDates(locations.Dates)
	if err1 != nil {
		fmt.Println(err1)
	}
	Artists := utils.GetArtists(utils.GetApiIndex().Artists)

	// Format locations
	formattedLocations := make(map[string][]string)
	for location, datesList := range relations.DatesLocations {
		formattedLocation := formatLocation(location)
		formattedLocations[formattedLocation] = datesList
	}

	data := struct {
		Artist    utils.Artist
		Locations utils.Location
		Dates     utils.Date
		Relations utils.Relations
	}{
		Artist:    Artists[id-1],
		Locations: locations,
		Dates:     dates,
		Relations: utils.Relations{DatesLocations: formattedLocations},
	}

	tmpl, err := template.New("locations.html").Funcs(template.FuncMap{
		"hasPrefix":  strings.HasPrefix,
		"trimPrefix": strings.TrimPrefix,
	}).ParseFiles("templates/locations.html")
	if err != nil {
		fmt.Printf("error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func formatLocation(location string) string {
	// Replace "-" and "_" with spaces
	location = strings.ReplaceAll(location, "-", " ")
	location = strings.ReplaceAll(location, "_", " ")

	// Split the string into words
	words := strings.Fields(location)

	// Capitalize each word
	for i, word := range words {
		words[i] = strings.Title(strings.ToLower(word))
	}

	// Join the words back together
	return strings.Join(words, " ")
}
