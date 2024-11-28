package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

// FilterParams holds the filter criteria
type FilterParams struct {
	CreationDateStart int
	CreationDateEnd   int
	FirstAlbumStart   int
	FirstAlbumEnd     int
	Location          string
	SelectedMembers   []int
}

// ParseFilterParams extracts and validates filter parameters from the request
func ParseFilterParams(r *http.Request) (*FilterParams, error) {
	if err := r.ParseForm(); err != nil {
		return nil, fmt.Errorf("failed to parse form: %w", err)
	}

	params := &FilterParams{}

	// Parse date ranges
	var err error
	params.CreationDateStart, err = strconv.Atoi(r.FormValue("creationDateStart"))
	if err != nil {
		return nil, fmt.Errorf("invalid creation date start: %w", err)
	}

	params.CreationDateEnd, err = strconv.Atoi(r.FormValue("creationDateEnd"))
	if err != nil {
		return nil, fmt.Errorf("invalid creation date end: %w", err)
	}

	params.FirstAlbumStart, err = strconv.Atoi(r.FormValue("firstAlbumStart"))
	if err != nil {
		return nil, fmt.Errorf("invalid first album start: %w", err)
	}

	params.FirstAlbumEnd, err = strconv.Atoi(r.FormValue("firstAlbumEnd"))
	if err != nil {
		return nil, fmt.Errorf("invalid first album end: %w", err)
	}

	// Clean location string
	params.Location = cleanString(r.FormValue("location"))

	// Parse member checkboxes
	params.SelectedMembers = parseSelectedMembers(r)

	return params, nil
}

// cleanString normalizes a string by trimming spaces and replacing special characters
func cleanString(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, ",", "")
	return strings.ReplaceAll(s, "_", " ")
}

// parseSelectedMembers extracts selected member counts from checkboxes
func parseSelectedMembers(r *http.Request) []int {
	var selected []int
	for i := 1; i <= 8; i++ {
		if r.FormValue(fmt.Sprintf("members%d", i)) == "on" {
			selected = append(selected, i)
		}
	}
	return selected
}

// HandleFilters processes the filter request and returns matching artists
func HandleFilters(w http.ResponseWriter, r *http.Request) {
	params, err := ParseFilterParams(r)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	filteredBands := filterArtists(ModArtists, params)

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		fmt.Printf("template error: %v\n", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, filteredBands); err != nil {
		fmt.Printf("template execution error: %v\n", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

// filterArtists applies the filter criteria to the artist list
func filterArtists(artists []Artist, params *FilterParams) []Artist {
	var filtered []Artist

	for _, band := range artists {
		if !meetsDateCriteria(band, params) {
			continue
		}

		if !meetsMembersCriteria(band, params.SelectedMembers) {
			continue
		}

		if !meetsLocationCriteria(band, params.Location) {
			continue
		}

		filtered = append(filtered, band)
	}

	return filtered
}

// meetsDateCriteria checks if the band meets the date range criteria
func meetsDateCriteria(band Artist, params *FilterParams) bool {
	firstAlbumYear, err := strconv.Atoi(band.FirstAlbum[6:])
	if err != nil {
		return false
	}

	return band.CreationDate >= params.CreationDateStart &&
		band.CreationDate <= params.CreationDateEnd &&
		firstAlbumYear >= params.FirstAlbumStart &&
		firstAlbumYear <= params.FirstAlbumEnd
}

// meetsMembersCriteria checks if the band meets the member count criteria
func meetsMembersCriteria(band Artist, selected []int) bool {
	if len(selected) == 0 {
		return true
	}

	for _, n := range selected {
		if len(band.Members) == n {
			return true
		}
	}
	return false
}

// meetsLocationCriteria checks if the band meets the location criteria
func meetsLocationCriteria(band Artist, location string) bool {
	for _, place := range band.Locations {
		place = cleanString(place)
		if strings.Contains(place, location) || strings.Contains(location, place) {
			return true
		}
	}
	return false
}
