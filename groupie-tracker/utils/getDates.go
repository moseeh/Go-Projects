package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetDates(url string) (Date, error) {
	response, err := http.Get(url)
	if err != nil {
		return Date{}, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Date{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return Date{}, fmt.Errorf("error reading response body: %w", err)
	}

	var dates Date
	err = json.Unmarshal(bodyBytes, &dates)
	if err != nil {
		return Date{}, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return dates, nil
}
