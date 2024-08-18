package handlers

import (
    "groupie-tracker/utils"
    "net/http"
    "text/template"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
    // Fetch artists
    artists, err := utils.GetArtists(utils.GetApiIndex().Artists)
    if err != nil {
        http.Error(w, "Error fetching artists", http.StatusInternalServerError)
        return
    }

    // Check if artists slice is empty
    if len(artists) == 0 {
        http.Error(w, "No artists found", http.StatusNotFound)
        return
    }

    // Parse template
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Error parsing template", http.StatusInternalServerError)
        return
    }

    // Execute template
    err = tmpl.Execute(w, artists)
    if err != nil {
        http.Error(w, "Error executing template", http.StatusInternalServerError)
        return
    }
}