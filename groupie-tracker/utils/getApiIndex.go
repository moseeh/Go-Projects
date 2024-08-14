package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetApiIndex() ApiIndex {
	url := "https://groupietrackers.herokuapp.com/api"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return ApiIndex{}
	}
	defer response.Body.Close()
	var apiIndex ApiIndex

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return ApiIndex{}
		}
		err = json.Unmarshal(bodyBytes, &apiIndex)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return ApiIndex{}
		}
	}
	return apiIndex
}
