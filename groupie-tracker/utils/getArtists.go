package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetArtists(url string) []Artist {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", response.StatusCode)
		return nil
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	var artists []Artist
	err = json.Unmarshal(bodyBytes, &artists)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil
	}

	return artists
}
