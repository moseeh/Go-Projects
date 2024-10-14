package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, errorCode int) {
	var templateFile string
	switch errorCode {
	case http.StatusBadRequest:
		templateFile = "./templates/404.html"
	case http.StatusNotFound:
		templateFile = "./templates/404.html"
	case http.StatusMethodNotAllowed:
		templateFile = "./templates/405.html"
	case http.StatusInternalServerError:
		templateFile = "./templates/500.html"
	default:
		templateFile = "./templates/500.html"
	}

	tmpl := template.Must(template.ParseFiles(templateFile))
	w.WriteHeader(errorCode)
	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("Error rendering error template: %v", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
