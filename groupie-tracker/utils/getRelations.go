package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRelations(url string) (Relations, error) {
	response, err := http.Get(url)
	if err != nil {
		return Relations{}, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Relations{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return Relations{}, fmt.Errorf("error reading response body: %w", err)
	}

	var Relation Relations
	err = json.Unmarshal(bodyBytes, &Relation)
	if err != nil {
		return Relations{}, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return Relation, nil
}
