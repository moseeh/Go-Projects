package models

type AntColony struct {
	NumberOfAnts int
	Rooms        []Room
	Links        map[string][]string
	Start        string
	End          string
}

type Room struct {
	Name             string
	IsVisited        bool
	Coord_X, Coord_Y int
}

type Path struct {
	Rooms []string
}
