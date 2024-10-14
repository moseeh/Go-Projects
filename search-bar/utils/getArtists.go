package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetArtists(artistsURL string) ([]Artists, error) {
	resp, err := http.Get(artistsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch artists: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var artists []Artists
	if err := json.Unmarshal(data, &artists); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return artists, nil
}
