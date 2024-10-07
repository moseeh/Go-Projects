package main

import (
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.ArtistHandler)
	http.HandleFunc("/locations/", handlers.LocationHandler)
	http.HandleFunc("/search", handlers.SearchHandler)
	log.Println("The server is running on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
