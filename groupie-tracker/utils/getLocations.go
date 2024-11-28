package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocations(url string) ([]Location, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var locationIndex LocationIndex
	err = json.Unmarshal(bodyBytes, &locationIndex)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return locationIndex.Index, nil
}
