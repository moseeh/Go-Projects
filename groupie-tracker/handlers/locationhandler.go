package handlers

import (
	"groupie-tracker/utils"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL and validate
	idStr := strings.TrimPrefix(r.URL.Path, "/locations/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 || id > 52 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Fetch locations
	locations, err := utils.GetLocations(utils.GetApiIndex().Locations + "/" + strconv.Itoa(id))
	if err != nil {
		http.Error(w, "Error fetching locations", http.StatusInternalServerError)
		return
	}

	// Fetch relations
	relations, err := utils.GetRelations(utils.GetApiIndex().Relation + "/" + strconv.Itoa(id))
	if err != nil {
		http.Error(w, "Error fetching relations", http.StatusInternalServerError)
		return
	}

	// Fetch dates
	dates, err := utils.GetDates(locations.Dates)
	if err != nil {
		http.Error(w, "Error fetching dates", http.StatusInternalServerError)
		return
	}

	// Fetch artists
	artists, err := utils.GetArtists(utils.GetApiIndex().Artists)
	if err != nil {
		http.Error(w, "Error fetching artists", http.StatusInternalServerError)
	}
	if id > len(artists) {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

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
		Artist:    artists[id-1],
		Locations: locations,
		Dates:     dates,
		Relations: utils.Relations{DatesLocations: formattedLocations},
	}

	// Parse template
	tmpl, err := template.New("locations.html").Funcs(template.FuncMap{
		"hasPrefix":  strings.HasPrefix,
		"trimPrefix": strings.TrimPrefix,
	}).ParseFiles("templates/locations.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Execute template
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
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
