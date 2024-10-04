package handlers

import (
	"net/http"
	"text/template"

	"groupie-tracker/utils"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := utils.GetArtists(utils.GetApiIndex().Artists)
	if err != nil {
		http.Error(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}

	if len(artists) == 0 {
		http.Error(w, "No artists found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
