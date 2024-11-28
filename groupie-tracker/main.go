package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	// Serve static files from the "static" directory at the "/static/" URL path.
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.HandleUrls)
	// The server runs asynchronously on port 8080 using a goroutine.
	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Println(err)
		}
	}()
	fmt.Println("Server running on http://localhost:8000")
	select {} // keep the main function running indefinitely.
}
