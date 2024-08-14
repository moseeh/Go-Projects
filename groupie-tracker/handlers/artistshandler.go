package handlers

import (
	"fmt"
	"groupie-tracker/utils"
	"net/http"
	"text/template"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artists := utils.GetArtists(utils.GetApiIndex().Artists)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Printf("error: %v", err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artists)
}
