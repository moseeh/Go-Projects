package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"groupie-tracker/utils"
)

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := utils.GetArtists(utils.GetApiIndex().Artists)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	locations, err := utils.GetLocations(utils.GetApiIndex().Locations)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	dates, err := utils.GetDates(utils.GetApiIndex().Dates)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	var modArtists []Artist
	for i, v := range artists {
		var modArtist Artist
		modArtist.Id = v.Id
		modArtist.Image = v.Image
		modArtist.Name = v.Name
		modArtist.Members = v.Members
		modArtist.CreationDate = v.CreationDate
		modArtist.FirstAlbum = v.FirstAlbum
		modArtist.Locations = locations[i].Locations
		modArtist.ConcertDates = dates[i].Dates
		modArtists = append(modArtists, modArtist)
	}
	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		fmt.Printf("error: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, modArtists)
}
