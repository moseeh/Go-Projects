package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func RouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/generate" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	http.DefaultServeMux.ServeHTTP(w, r)
}
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func HandleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Read the banner file
	content, err := ReadBannerFile(banner)
	if err != nil {
		http.Error(w, "Invalid banner file", http.StatusBadRequest)
		return
	}

	contentLines := SplitFile(string(content))
	if len(contentLines) != 856 {
		http.Error(w, "Invalid banner file", http.StatusBadRequest)
		return
	}

	result, err := DisplayText(text, contentLines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("result.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Result string
	}{
		Result: result,
	}

	tmpl.Execute(w, data)
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 - Page Not Found")
}
