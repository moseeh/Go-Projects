package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem-in/models"
)

func ParseFile(filename string) (*models.AntColony, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	colony := &models.AntColony{
		Links: make(map[string][]string),
		Rooms: make([]models.Room, 0),
	}
	room := &models.Room{}

	lineNum := 0
	var currentCommand string

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if lineNum == 1 {
			colony.NumberOfAnts, err = strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("invalid number of ants on line %d: %v", lineNum, err)
			}
			continue
		}

		if strings.HasPrefix(line, "#") {
			if line == "##start" || line == "##end" {
				currentCommand = line
			}
			continue
		}

		if strings.Contains(line, " ") {
			// Room definition
			parts := strings.Fields(line)
			if len(parts) != 3 {
				return nil, fmt.Errorf("invalid room definition on line %d", lineNum)
			}

			x, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid x-coordinate on line %d: %v", lineNum, err)
			}

			y, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, fmt.Errorf("invalid y-coordinate on line %d: %v", lineNum, err)
			}
			room.Name = parts[0]
			room.Coord_X = x
			room.Coord_Y = y

			colony.Rooms = append(colony.Rooms, *room)

			if currentCommand == "##start" {
				colony.Start = parts[0]
				currentCommand = ""
			} else if currentCommand == "##end" {
				colony.End = parts[0]
				currentCommand = ""
			}
		} else if strings.Contains(line, "-") {
			// Link definition
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid link definition on line %d", lineNum)
			}
			room1, room2 := parts[0], parts[1]
			colony.Links[room1] = append(colony.Links[room1], room2)
			colony.Links[room2] = append(colony.Links[room2], room1)
		} else {
			return nil, fmt.Errorf("invalid line format on line %d", lineNum)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	if colony.Start == "" || colony.End == "" {
		return nil, fmt.Errorf("start or end room not defined")
	}

	return colony, nil
}
