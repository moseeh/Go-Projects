package handlers

import (
	"groupie-tracker/utils"
	"net/http"
	"strings"
	"text/template"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	artists, err := utils.GetArtists(utils.GetApiIndex().Artists)
	if err != nil {
		http.Error(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	query = strings.ToLower(query)
	var filteredArtists []utils.Artist

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	tmpl.Execute(w, filteredArtists)
}

// func containsIgnoreCase(slice []string, query string) bool {
// 	for _, item := range slice {
// 		if strings.Contains(strings.ToLower(item), query) {
// 			return true
// 		}
// 	}
// 	return false
// }
