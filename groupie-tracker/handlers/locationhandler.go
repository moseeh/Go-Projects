package handlers

import (
    "fmt"
    "net/http"
    "strconv"
    "text/template"

    "groupie-tracker/utils"
)

func LocationHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.URL.Path[len("/locations/"):])
    locations, err := utils.GetLocations(utils.GetApiIndex().Locations + "/" + strconv.Itoa(id))
    if err != nil {
        fmt.Println(err)
    }
    
    // Retrieve artist information for the given ID
    Artists := utils.GetArtists(utils.GetApiIndex().Artists)

    data := struct {
        Artist    utils.Artist
        Locations utils.Location
    }{
        Artist:    Artists[id-1],
        Locations: locations,
    }

    tmpl, err := template.ParseFiles("templates/locations.html")
    if err != nil {
        fmt.Printf("error: %v", err)
        http.Error(w, "Internal Server error", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, data)
}
