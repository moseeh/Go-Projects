package handlers

import (
	"groupie-tracker/utils"
	"net/http"
	"strings"
	"text/template"
)

// SearchHandler handles search requests.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Search query is required", http.StatusBadRequest)
		return
	}

	// Fetch the data
	apiIndex := utils.GetApiIndex()
	artists, err := utils.GetArtists(apiIndex.Artists)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	// Filter artists based on the search query
	var filteredArtists []utils.Artists // Define Artist struct in utils
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	// Parse and execute the template
	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, filteredArtists)
}
