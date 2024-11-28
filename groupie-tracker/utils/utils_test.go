package utils_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"groupie-tracker/utils"
)

var url = "https://groupietrackers.herokuapp.com/api"

func TestGetApiIndex(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the correct endpoint is being called
		if r.URL.String() != "/api" {
			t.Errorf("Expected to request '/api', got: %s", r.URL.String())
		}

		// Create a mock ApiIndex
		mockApiIndex := utils.ApiIndex{
			Artists:   "https://groupietrackers.herokuapp.com/api/artists",
			Locations: "https://groupietrackers.herokuapp.com/api/locations",
			Dates:     "https://groupietrackers.herokuapp.com/api/dates",
			Relation:  "https://groupietrackers.herokuapp.com/api/relation",
		}

		// Write the mock data as response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockApiIndex)
	}))
	defer server.Close()

	// Save the original URL and replace it with our test server URL
	originalURL := url
	url = server.URL + "/api"
	defer func() { url = originalURL }()

	// Call the function we're testing
	result := utils.GetApiIndex()

	// Define the expected result
	expected := utils.ApiIndex{
		Artists:   "https://groupietrackers.herokuapp.com/api/artists",
		Locations: "https://groupietrackers.herokuapp.com/api/locations",
		Dates:     "https://groupietrackers.herokuapp.com/api/dates",
		Relation:  "https://groupietrackers.herokuapp.com/api/relation",
	}

	// Compare the result with the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetApiIndex() = %v, want %v", result, expected)
	}
}

func TestGetArtists(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the correct endpoint is being called
		if r.URL.String() != "/api/artists" {
			t.Errorf("Expected to request '/api/artists', got: %s", r.URL.String())
		}

		// Create mock artists data based on the provided sample
		mockArtists := []utils.Artists{
			{
				Id:           1,
				Image:        "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
				Name:         "Queen",
				Members:      []string{"Freddie Mercury", "Brian May", "John Daecon", "Roger Meddows-Taylor", "Mike Grose", "Barry Mitchell", "Doug Fogie"},
				CreationDate: 1970,
				FirstAlbum:   "14-12-1973",
				Locations:    "https://groupietrackers.herokuapp.com/api/locations/1",
				ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/1",
				Relations:    "https://groupietrackers.herokuapp.com/api/relation/1",
			},
			{
				Id:           2,
				Image:        "https://groupietrackers.herokuapp.com/api/images/soja.jpeg",
				Name:         "SOJA",
				Members:      []string{"Jacob Hemphill", "Bob Jefferson", "Ryan \"Byrd\" Berty", "Ken Brownell", "Patrick O'Shea", "Hellman Escorcia", "Rafael Rodriguez", "Trevor Young"},
				CreationDate: 1997,
				FirstAlbum:   "05-06-2002",
				Locations:    "https://groupietrackers.herokuapp.com/api/locations/2",
				ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/2",
				Relations:    "https://groupietrackers.herokuapp.com/api/relation/2",
			},
		}

		// Write the mock data as response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockArtists)
	}))
	defer server.Close()

	// Call the function we're testing
	result, err := utils.GetArtists(server.URL + "/api/artists")
	// Check for errors
	if err != nil {
		t.Fatalf("GetArtists() returned an error: %v", err)
	}

	// Define the expected result
	expected := []utils.Artists{
		{
			Id:           1,
			Image:        "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
			Name:         "Queen",
			Members:      []string{"Freddie Mercury", "Brian May", "John Daecon", "Roger Meddows-Taylor", "Mike Grose", "Barry Mitchell", "Doug Fogie"},
			CreationDate: 1970,
			FirstAlbum:   "14-12-1973",
			Locations:    "https://groupietrackers.herokuapp.com/api/locations/1",
			ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/1",
			Relations:    "https://groupietrackers.herokuapp.com/api/relation/1",
		},
		{
			Id:           2,
			Image:        "https://groupietrackers.herokuapp.com/api/images/soja.jpeg",
			Name:         "SOJA",
			Members:      []string{"Jacob Hemphill", "Bob Jefferson", "Ryan \"Byrd\" Berty", "Ken Brownell", "Patrick O'Shea", "Hellman Escorcia", "Rafael Rodriguez", "Trevor Young"},
			CreationDate: 1997,
			FirstAlbum:   "05-06-2002",
			Locations:    "https://groupietrackers.herokuapp.com/api/locations/2",
			ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/2",
			Relations:    "https://groupietrackers.herokuapp.com/api/relation/2",
		},
	}

	// Compare the result with the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetArtists() = %v, want %v", result, expected)
	}
}

// ... TestGetArtistsError function remains the same ...
func TestGetArtistsError(t *testing.T) {
	// Test with invalid URL
	_, err := utils.GetArtists("http://invalid-url")
	if err == nil {
		t.Error("Expected error with invalid URL, got nil")
	}

	// Test with server error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err = utils.GetArtists(server.URL)
	if err == nil {
		t.Error("Expected error with server error, got nil")
	}

	// Test with invalid JSON
	serverInvalidJSON := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer serverInvalidJSON.Close()

	_, err = utils.GetArtists(serverInvalidJSON.URL)
	if err == nil {
		t.Error("Expected error with invalid JSON, got nil")
	}
}

func TestGetDates(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the correct endpoint is being called
		if r.URL.String() != "/api/dates" {
			t.Errorf("Expected to request '/api/dates', got: %s", r.URL.String())
		}

		// Create mock date data based on the provided sample
		mockDateIndex := utils.DateIndex{
			Index: []utils.Date{
				{
					Id: 1,
					Dates: []string{
						"*23-08-2019", "*22-08-2019", "*20-08-2019",
						"*26-01-2020", "*28-01-2020", "*30-01-2019",
						"*07-02-2020", "*10-02-2020",
					},
				},
				{
					Id: 2,
					Dates: []string{
						"*05-12-2019", "06-12-2019", "07-12-2019",
						"08-12-2019", "09-12-2019", "*16-11-2019",
						"*15-11-2019",
					},
				},
				{
					Id: 3,
					Dates: []string{
						"*10-05-2007", "02-07-2005", "29-10-1994",
						"28-10-1994", "27-10-1994", "26-10-1994",
						"23-10-1994", "22-10-1994", "21-10-1994",
						"20-10-1994", "19-10-1994", "17-10-1994",
						"16-10-1994", "15-10-1994", "14-10-1994",
						"13-10-1994", "12-10-1994", "*25-09-1994",
						"*23-09-1994",
					},
				},
			},
		}

		// Write the mock data as response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockDateIndex)
	}))
	defer server.Close()

	// Call the function we're testing
	result, err := utils.GetDates(server.URL + "/api/dates")
	// Check for errors
	if err != nil {
		t.Fatalf("utils.GetDates() returned an error: %v", err)
	}

	// Define the expected result
	expected := []utils.Date{
		{
			Id: 1,
			Dates: []string{
				"*23-08-2019", "*22-08-2019", "*20-08-2019",
				"*26-01-2020", "*28-01-2020", "*30-01-2019",
				"*07-02-2020", "*10-02-2020",
			},
		},
		{
			Id: 2,
			Dates: []string{
				"*05-12-2019", "06-12-2019", "07-12-2019",
				"08-12-2019", "09-12-2019", "*16-11-2019",
				"*15-11-2019",
			},
		},
		{
			Id: 3,
			Dates: []string{
				"*10-05-2007", "02-07-2005", "29-10-1994",
				"28-10-1994", "27-10-1994", "26-10-1994",
				"23-10-1994", "22-10-1994", "21-10-1994",
				"20-10-1994", "19-10-1994", "17-10-1994",
				"16-10-1994", "15-10-1994", "14-10-1994",
				"13-10-1994", "12-10-1994", "*25-09-1994",
				"*23-09-1994",
			},
		},
	}

	// Compare the result with the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("utils.GetDates() = %v, want %v", result, expected)
	}
}

func TestGetDatesError(t *testing.T) {
	// Test with invalid URL
	_, err := utils.GetDates("http://invalid-url")
	if err == nil {
		t.Error("Expected error with invalid URL, got nil")
	}

	// Test with server error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err = utils.GetDates(server.URL)
	if err == nil {
		t.Error("Expected error with server error, got nil")
	}

	// Test with invalid JSON
	serverInvalidJSON := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer serverInvalidJSON.Close()

	_, err = utils.GetDates(serverInvalidJSON.URL)
	if err == nil {
		t.Error("Expected error with invalid JSON, got nil")
	}
}

func TestGetLocations(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the correct endpoint is being called
		if r.URL.String() != "/api/locations" {
			t.Errorf("Expected to request '/api/locations', got: %s", r.URL.String())
		}

		// Create mock location data based on the provided sample
		mockLocationIndex := utils.LocationIndex{
			Index: []utils.Location{
				{
					Id: 1,
					Locations: []string{
						"north_carolina-usa",
						"georgia-usa",
						"los_angeles-usa",
						"saitama-japan",
						"osaka-japan",
						"nagoya-japan",
						"penrose-new_zealand",
						"dunedin-new_zealand",
					},
					Dates: "https://groupietrackers.herokuapp.com/api/dates/1",
				},
				{
					Id: 2,
					Locations: []string{
						"playa_del_carmen-mexico",
						"papeete-french_polynesia",
						"noumea-new_caledonia",
					},
					Dates: "https://groupietrackers.herokuapp.com/api/dates/2",
				},
				{
					Id: 3,
					Locations: []string{
						"london-uk",
						"lausanne-switzerland",
						"lyon-france",
					},
					Dates: "https://groupietrackers.herokuapp.com/api/dates/3",
				},
			},
		}

		// Write the mock data as response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockLocationIndex)
	}))
	defer server.Close()

	// Call the function we're testing
	result, err := utils.GetLocations(server.URL + "/api/locations")
	// Check for errors
	if err != nil {
		t.Fatalf("GetLocations() returned an error: %v", err)
	}

	// Define the expected result
	expected := []utils.Location{
		{
			Id: 1,
			Locations: []string{
				"north_carolina-usa",
				"georgia-usa",
				"los_angeles-usa",
				"saitama-japan",
				"osaka-japan",
				"nagoya-japan",
				"penrose-new_zealand",
				"dunedin-new_zealand",
			},
			Dates: "https://groupietrackers.herokuapp.com/api/dates/1",
		},
		{
			Id: 2,
			Locations: []string{
				"playa_del_carmen-mexico",
				"papeete-french_polynesia",
				"noumea-new_caledonia",
			},
			Dates: "https://groupietrackers.herokuapp.com/api/dates/2",
		},
		{
			Id: 3,
			Locations: []string{
				"london-uk",
				"lausanne-switzerland",
				"lyon-france",
			},
			Dates: "https://groupietrackers.herokuapp.com/api/dates/3",
		},
	}

	// Compare the result with the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetLocations() = %v, want %v", result, expected)
	}
}

func TestGetLocationsError(t *testing.T) {
	// Test with invalid URL
	_, err := utils.GetLocations("http://invalid-url")
	if err == nil {
		t.Error("Expected error with invalid URL, got nil")
	}

	// Test with server error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err = utils.GetLocations(server.URL)
	if err == nil {
		t.Error("Expected error with server error, got nil")
	}

	// Test with invalid JSON
	serverInvalidJSON := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer serverInvalidJSON.Close()

	_, err = utils.GetLocations(serverInvalidJSON.URL)
	if err == nil {
		t.Error("Expected error with invalid JSON, got nil")
	}
}
