package handlers

import (
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
