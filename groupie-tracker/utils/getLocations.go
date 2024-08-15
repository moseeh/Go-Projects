package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocations(url string) (Location, error) {
	response, err := http.Get(url)
	if err != nil {
		return Location{}, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return Location{}, fmt.Errorf("error reading response body: %w", err)
	}

	var location Location
	err = json.Unmarshal(bodyBytes, &location)
	if err != nil {
		return Location{}, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	// Extract locations from locationIndex, assuming LocationIndex has a field Locations of type []Location
	return location, nil
}
