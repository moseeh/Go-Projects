package main

import (
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.ArtistHandler)
	http.HandleFunc("/locations/", handlers.LocationHandler)
	http.ListenAndServe(":8080", nil)
}
