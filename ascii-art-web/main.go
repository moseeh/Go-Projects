package main

import (
	"fmt"
	"net/http"
	"web/utils"
)

func main() {
	http.HandleFunc("/", utils.HandleIndex)
	http.HandleFunc("/generate", utils.HandleGenerate)

	// Handle 404 errors
	http.HandleFunc("/404", utils.Handle404)
	http.HandleFunc("/favicon.ico", utils.Handle404)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", http.HandlerFunc(utils.RouteHandler))
}
