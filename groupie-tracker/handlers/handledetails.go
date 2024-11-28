package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"groupie-tracker/utils"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL and validate
	idStr := strings.TrimPrefix(r.URL.Path, "/details/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 || id > 52 {
		// http.Error(w, "Invalid ID", http.StatusBadRequest)
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Fetch relations
	relations, err := utils.GetRelations(utils.GetApiIndex().Relation + "/" + strconv.Itoa(id))
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch artists
	artists, err := utils.GetArtists(utils.GetApiIndex().Artists)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	if id > len(artists) {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	// Format locations
	formattedLocations := make(map[string][]string)
	for location, datesList := range relations.DatesLocations {
		formattedLocation := formatLocation(location)
		formattedLocations[formattedLocation] = datesList
	}

	data := struct {
		Artist    utils.Artists
		Relations utils.Relations
	}{
		Artist:    artists[id-1],
		Relations: utils.Relations{DatesLocations: formattedLocations},
	}

	// Parse template
	tmpl, err := template.New("details.html").Funcs(template.FuncMap{
		"hasPrefix":  strings.HasPrefix,
		"trimPrefix": strings.TrimPrefix,
	}).ParseFiles("templates/details.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Execute template
	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func formatLocation(location string) string {
	// Replace "-" and "_" with spaces
	location = strings.ReplaceAll(location, "-", ", ")
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
