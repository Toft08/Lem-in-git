package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Ant struct {
	ID         int
	Name       int
	Position   string // Keeping track of the current room
	ReachedEnd bool
}

// Simulating ant movements by minimal amount of turns for the given path combination (CHANGE)
func simulateAntMovement(paths [][]string, numAnts int, start, end string) []string {
	ants := make([]Ant, numAnts)
	for i := 0; i < numAnts; i++ {
		ants[i] = Ant{ID: i + 1, Position: start}
	}

	// Assigning ants to paths
	assignedPath := assignAntsToPaths(paths, numAnts)

	// Slice to store movements
	movements := []string{}
	nextAntName := 1

	// Simulation loop recording all the ant movements
	for {
 		allFinished := true
		tunnelInUse := make(map[string]bool) // Reset tunnel usage for each turn

		turnMovements := []string{} 

		for i := range ants {
			ant := &ants[i]

			// Skip if the ant has already reached the end
			if ant.ReachedEnd {
				continue
			}

			// Get the path assigned to this ant
			path := assignedPath[ant.ID]
			currentRoomIdx := indexOf(path, ant.Position)
			if currentRoomIdx == -1 {
				return nil
			}

			// Determine the next room
			if currentRoomIdx+1 < len(path) {
				nextRoom := path[currentRoomIdx+1]
				tunnel := fmt.Sprintf("%s->%s", ant.Position, nextRoom)

				// Move only if the tunnel is not in use then mark it for this turn
				if !tunnelInUse[tunnel] {
					tunnelInUse[tunnel] = true

					// Move the ant
					ant.Position = nextRoom

					// Giving the name for the ant in the order of moving (CHANGE)
					if ant.Name == 0 {
						ant.Name = nextAntName
						nextAntName++
					}

					// Record the movement
					turnMovements = append(turnMovements, fmt.Sprintf("L%d-%s", ant.Name, ant.Position))

					if nextRoom == end {
						ant.ReachedEnd = true
					}
				}
			}

			// Check if all ants are finished
			if !ant.ReachedEnd {
				allFinished = false
			}
		}

		// Add current turn movements to all movements struct
		if len(turnMovements) > 0 {
			turnMovements = ascendingOrder(turnMovements)
			movements = append(movements, strings.Join(turnMovements, " "))
		}

		if allFinished {
			break
		}
	}

	return movements
}

// Assigning paths to all ants based on the queue length and path structure
func assignAntsToPaths(paths [][]string, numAnts int) map[int][]string {
	// Initialize the assignedPath map
	assignedPath := make(map[int][]string) 

	// Track the number of ants assigned to each path
	pathAntCounts := make([]int, len(paths))

	// Assign ants to paths
	for antID := 1; antID <= numAnts; antID++ {
		bestPath := -1           // Giving value outside of possible path index (assuming invalid scenario)
		minLength := math.MaxInt 

		for i, path := range paths {
			length := len(path) + pathAntCounts[i]
			if length < minLength {
				bestPath = i
				minLength = length
			}
		}

		// Assign the ant to the best path
		pathAntCounts[bestPath]++
		assignedPath[antID] = paths[bestPath]
	}

	return assignedPath
}

// Finding the index for the room on given path
func indexOf(path []string, room string) int {
	for i, r := range path {
		if r == room {
			return i
		}
	}
	return -1 
}

// Changing the movements into ascending order based on the ant name
func ascendingOrder(movements []string) []string {
	if len(movements) > 1 {
		for i := 0; i < len(movements)-1; i++ {
			for j := 0; j < len(movements)-i-1; j++ {
				ant1 := getAntName(movements[j])
				ant2 := getAntName(movements[j+1])
				if ant1 > ant2 {
					movements[j], movements[j+1] = movements[j+1], movements[j]
				}
			}
		}
	}
	return movements
}

// Getting the ant name from the movement (L1-3 -> ant name 1)
func getAntName(s string) int {
	antmovement := strings.Split(s, "-")
	ant := strings.TrimPrefix(antmovement[0], "L")
	antInt, _ := strconv.Atoi(ant)
	return antInt
}
