package main

import (
	"log"
	"net/http"

	"forum/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	log.Println("The server is running on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
