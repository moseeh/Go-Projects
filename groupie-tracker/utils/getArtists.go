package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetArtists(url string) ([]Artist, error) {
	// Make HTTP request
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer response.Body.Close()

	// Check status code
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	// Read response body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Unmarshal JSON
	var artists []Artist
	err = json.Unmarshal(bodyBytes, &artists)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return artists, nil
}
