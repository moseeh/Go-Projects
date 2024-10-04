package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"groupie-tracker/utils"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := utils.GetArtists(utils.GetApiIndex().Artists)
	if err != nil {
		http.Error(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	query = strings.ToLower(query)
	var filteredArtists []utils.Artist

	for _, artist := range artists {
		if strings.HasPrefix(strings.ToLower(artist.Name), query) {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	tmpl.Execute(w, filteredArtists)
}

func ApiSearchHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := utils.GetArtists(utils.GetApiIndex().Artists)
	if err != nil {
		http.Error(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}

	// Get search query
	query := strings.ToLower(r.URL.Query().Get("query"))

	// Get pagination params (page and limit)
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	// Default values for pagination
	page := 1
	limit := 10

	if pageParam != "" {
		page, err = strconv.Atoi(pageParam)
		if err != nil || page <= 0 {
			page = 1
		}
	}

	if limitParam != "" {
		limit, err = strconv.Atoi(limitParam)
		if err != nil || limit <= 0 {
			limit = 10
		}
	}

	// Filter the artists by the search query
	var filteredArtists []utils.Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	// Calculate total pages
	total := len(filteredArtists)
	totalPages := (total + limit - 1) / limit

	// Ensure current page doesn't exceed total pages
	if page > totalPages {
		page = totalPages
	}

	// Paginate the filtered artists
	start := (page - 1) * limit
	end := start + limit
	if end > total {
		end = total
	}

	paginatedArtists := filteredArtists[start:end]

	// Return paginated data with page info
	response := map[string]interface{}{
		"artists":     paginatedArtists,
		"totalPages":  totalPages,
		"currentPage": page,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
