package utils

import (
    "fmt"
    "strings"
    "lem-in/models"
)

func PrintTurns(paths []models.Path, distribution []int, numOfAnts int) {
    type antPosition struct {
        pathIdx int
        step    int    // starts at 1 to skip first room
    }
    antPositions := make(map[int]antPosition)
    
    // Initialize starting positions for all ants, starting at step 1 to skip first room
    currentAnt := 1
    for pathIdx, antCount := range distribution {
        for i := 0; i < antCount; i++ {
            if currentAnt <= numOfAnts {
                antPositions[currentAnt] = antPosition{
                    pathIdx: pathIdx,
                    step: 1, // Start at 1 to skip first room
                }
                currentAnt++
            }
        }
    }

    // Continue until all ants reach the end
    finished := false
    for !finished {
        finished = true
        var roundMoves []string
        occupiedRooms := make(map[string]bool)
        endRoomPathUsed := make(map[int]bool) // Track which paths have used the end room this turn
        movedAnts := make(map[int]bool)

        // Process each ant's movement
        for ant := 1; ant <= numOfAnts; ant++ {
            pos, exists := antPositions[ant]
            if !exists || movedAnts[ant] {
                continue
            }

            if pos.step < len(paths[pos.pathIdx].Rooms) {
                currentRoom := paths[pos.pathIdx].Rooms[pos.step]
                isEndRoom := pos.step == len(paths[pos.pathIdx].Rooms)-1
                
                // Check if room is available
                if occupiedRooms[currentRoom] {
                    if isEndRoom {
                        // For end room, check if this path has already been used
                        if endRoomPathUsed[pos.pathIdx] {
                            continue
                        }
                    } else {
                        continue
                    }
                }

                finished = false
                move := fmt.Sprintf("L%d-%s", ant, currentRoom)
                roundMoves = append(roundMoves, move)
                
                // Mark room and ant as used for this turn
                occupiedRooms[currentRoom] = true
                if isEndRoom {
                    endRoomPathUsed[pos.pathIdx] = true
                }
                movedAnts[ant] = true
                
                // Move ant to next position
                antPositions[ant] = antPosition{
                    pathIdx: pos.pathIdx,
                    step: pos.step + 1,
                }
            }
        }

        // Print the round's moves
        if len(roundMoves) > 0 {
            fmt.Println(strings.Join(roundMoves, " "))
        }
    }
}