package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X    int
	Y    int
}

type ParsedData struct {
	NumAnts   int
	StartRoom string
	EndRoom   string
	Rooms     map[string]Room
	Tunnels   map[string][]string
}

// Parses global variable fileContent as per structs and returns the data
func parseInput(fileContent []string) (*ParsedData, error) {
	// creating dynamic data
	parsedData := &ParsedData{
		Rooms:   make(map[string]Room),
		Tunnels: make(map[string][]string),
	}

	numAnts, err := strconv.Atoi(fileContent[0])
	if err != nil || numAnts < 1 {
		return nil, fmt.Errorf("invalid number of ants '%v'", fileContent[0])
	}
	parsedData.NumAnts = numAnts

	// Parsing rooms and tunnels (CHANGE)flagging that the next rooms is a starting room, flagging that the next rooms is a starting room
	var isStart, isEnd bool
	for _, line := range fileContent[1:] {
		line = strings.TrimSpace(line)
		if line == "##start" {
			isStart = true
			continue
		}
		if line == "##end" {
			isEnd = true
			continue
		}

		// Checking if the line defines a room (length of the slice is 3) (CHANGE)
		parts := strings.Fields(line)
		if len(parts) == 3 {
			name := parts[0]
			if name[0] == 'L' || name[0] == '#' { // Name cannot start with a L or # (CHANGE)
				return nil, fmt.Errorf("invalid room name")
			}
			x, err1 := strconv.Atoi(parts[1])
			y, err2 := strconv.Atoi(parts[2])
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("invalid room coordinates")
			}
			room := Room{Name: name, X: x, Y: y} // creating the room (CHANGE)
			for _, existingRoom := range parsedData.Rooms {
				if existingRoom.X == x && existingRoom.Y == y {
					return nil, fmt.Errorf("Duplicate room coordinates")
				} else if existingRoom.Name == name {
					return nil, fmt.Errorf("Duplicate room name")
				}
			}
			parsedData.Rooms[name] = room // adding the room to the map (CHANGE)

			// Checking if the room is a start or end room
			if isStart {
				if parsedData.StartRoom != "" {
					return nil, fmt.Errorf("several start rooms defined")
				}
				parsedData.StartRoom = name
				isStart = false
			} else if isEnd {
				if parsedData.EndRoom != "" {
					return nil, fmt.Errorf("several end rooms defined")
				}
				parsedData.EndRoom = name
				isEnd = false
			}
			continue
		}

		// Checking if the line defines a connection (includes "-") (CHANGE)looking for conns
		if strings.Contains(line, "-") {
			connParts := strings.Split(line, "-")
			if len(connParts) != 2 {
				return nil, fmt.Errorf("invalid connection: %v", line)
			}
			room1, room2 := connParts[0], connParts[1]

			if !roomExists(room1, parsedData.Rooms) || !roomExists(room2, parsedData.Rooms) {
				continue
			}

			parsedData.Tunnels[room1] = append(parsedData.Tunnels[room1], room2) // adding the connection on both tunnel maps
			parsedData.Tunnels[room2] = append(parsedData.Tunnels[room2], room1)
		}
	}

	// Checking if the star or end room is missing
	if parsedData.StartRoom == "" {
		return nil, fmt.Errorf("start room not defined")
	}
	if parsedData.EndRoom == "" {
		return nil, fmt.Errorf("end room not defined")
	}

	return parsedData, nil
}

// Checking if the room exists in the given map
func roomExists(room string, allRooms map[string]Room) bool {

	for r := range allRooms {
		if r == room {
			return true
		}
	}
	return false
}
