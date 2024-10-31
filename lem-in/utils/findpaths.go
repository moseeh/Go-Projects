package utils

import (
	"lem-in/models"
)

// FindPaths finds all possible paths from start to end using BFS
func FindPaths(colony *models.AntColony) []models.Path {
	var allPaths []models.Path
	var queue [][]string
	queue = append(queue, []string{colony.Start})

	for i := range colony.Rooms {
		colony.Rooms[i].IsVisited = false
	}

	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]
		currentRoom := currentPath[len(currentPath)-1]
		if currentRoom == colony.End {
			allPaths = append(allPaths, models.Path{Rooms: currentPath})
			continue
		}
		for _, nextRoom := range colony.Links[currentRoom] {
			if !containsRoom(currentPath, nextRoom) {
				// Create a new path with the next room added
				newPath := make([]string, len(currentPath))
				copy(newPath, currentPath)
				newPath = append(newPath, nextRoom)
				queue = append(queue, newPath)
			}
		}
	}

	return allPaths
}

// Helper function to check if a room is in a path
func containsRoom(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}
