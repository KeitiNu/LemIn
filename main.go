package main

import (
	"errors"
	"fmt"
	"lem-in/path"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Checks, if the arguments have been given correctly
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [INPUTFILE]")
		fmt.Println("EX: go run . example00.txt")
	}

	//Reads in the file from arguments and splits in into strings
	fileName := "./examples/" + os.Args[1]
	input, err := os.ReadFile(fileName)
	errorCheck(err)
	data := strings.Split(string(input), "\n")

	//Finds the number of ants
	numAnts, roomData := numberOfAnts(data)

	//Sorts the data into coordinate data and relation data
	locationData, relationData := sortData(roomData)
	coordinatesMap := mapRoomCoordinates(locationData)
	originalMap := mapRoomConnections(relationData, coordinatesMap) //unsorted map of rooms and relations
	croppedMap := removeDeadEnds(originalMap)


	//FINAL PRINTING

	//Info about the paths and distribution from path.go
	path, distribution := path.Path(croppedMap, numAnts)

	fmt.Println(string(input) + "\n")

	//func print(path [][]string, distribution []int, numAnts int) {
	var ants [][]int
	antNr := 1
	//Creating a slice into var ants[][] that contains integers of how many ants there are
	for i := 0; i < len(distribution); i++ {
		ants = append(ants, []int{})
		for j := 0; j < distribution[i]; j++ {
			ants[i] = append(ants[i], antNr)
			antNr++
		}
	}

	biggest := 0
	//Assigning the largest value from var ants to var biggest
	for i := 0; i < len(ants); i++ {
		if len(ants[i]) > biggest {
			biggest = len(ants[i])
		}
	}

	connected := false
	remaining := numAnts
	//Creating a infite loop that prints the required output and ends when var remaining is equal to 0
	for i := 1; i > 0; i++ {
		whichANT := 0
		for k := 0; k < biggest; k++ {
			for j := 0; j < len(ants); j++ {
				if len(ants[j]) > whichANT && len(path[j]) > i-whichANT && i-whichANT > 0 {
					fmt.Print("L")
					fmt.Print(ants[j][whichANT])
					fmt.Print("-")
					fmt.Print(path[j][i-whichANT])
					fmt.Print(" ")
					if i-whichANT == len(path[j])-1 {
						remaining--
					}
					if len(path[j]) == 2 {
						connected = true
					} else {
						connected = false
					}
				}

			}
			whichANT++
		}
		if !connected {
			fmt.Println()
		}

		if remaining == 0 {
			if connected {
				fmt.Println()
			}

			break
		}

	}
}

//removes dead ends from function
func removeDeadEnds(rawMap map[string][]string) map[string][]string {
	var err error

	for key, value := range rawMap {
		if len(value) == 1 {
			if value[0] == "start" || value[0] == "end" {
				err = errors.New("missing a route to end or start")
				errorCheck(err)
			} else {
				delete(rawMap, key)
			}
		} else if len(value) == 2 && value[0] != "start" && value[0] != "end" {
			for key1, values := range rawMap {
				for i, v := range values {
					if v == key {
						values = remove(values, i)
						rawMap[key1] = values
					}
				}
			}
			delete(rawMap, key)
			removeDeadEnds(rawMap)
		}
	}

	return rawMap
}

//removes a named value from string
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

//Breaks the remaining arguments to room info and relation info
func sortData(roomData []string) ([]string, []string) {

	var locationData []string

	//finds data related to coordinates of the room
	for i := 0; i < len(roomData); i++ {
		if len(strings.Split(roomData[i], " ")) == 3 {
			locationData = append(locationData, roomData[i])
		} else if roomData[i][0] == '#' {
			if roomData[i][1] == '#' {
				locationData = append(locationData, roomData[i])
			}
		} else {
			roomData = roomData[i:]
			break
		}
	}

	//finds room relations data
	var realationData []string

	for _, v := range roomData {
		if len(strings.Split(v, "-")) == 2 {
			realationData = append(realationData, v)
		} else if v[0] == '#' && v[1] != '#' {

		} else {
			var err error = nil
			err = errors.New("data in wrong format")
			errorCheck(err)
		}
	}

	return locationData, realationData
}

//Finds the number of ants
func numberOfAnts(data []string) (int, []string) {
	var nrOfAnts int
	var roomData []string
	var err error

	for i := 0; i < len(data); i++ {
		if data[i][0] != '#' {
			nrOfAnts, err = strconv.Atoi(data[i])
			if nrOfAnts < 1 {
				err = errors.New("not enough ants")
			}
			errorCheck(err)
			roomData = data[i+1:]
			break
		}
	}

	return nrOfAnts, roomData
}

//Returns an error if err!=nil
func errorCheck(e error) {
	if e != nil {
		fmt.Println("ERROR: invalid data format")
		log.Fatal(e)
	}
}

//Maps rooms and their coordinates
func mapRoomCoordinates(arguments []string) map[string][]int {
	roomMap := make(map[string][]int)
	var err error

	//checks for ##start and ##end location
	var startPoint int
	var endPoint int
	var startPointer *int
	var endPointer *int

	for i, v := range arguments {
		if v == "##start" {
			startPoint = i + 1
			startPointer = &startPoint
		} else if v == "##end" {
			endPoint = i + 1
			endPointer = &endPoint
		}
	}

	//If there´s a start or end point missing, send an error
	if startPointer == nil || endPointer == nil {
		err = errors.New("missing start or end")
		errorCheck(err)
	}

	//Puts each location and it´s coordinates to map
	//1: position(see next line), 2: x, 3: y
	//0:start, 1: middle, 2:end
	for i := 0; i < len(arguments); i++ {
		room := strings.Split(arguments[i], " ")
		if len(room) == 3 {
			roomX, err := strconv.Atoi(room[1])
			errorCheck(err)
			roomY, err := strconv.Atoi(room[2])
			errorCheck(err)
			if i == startPoint {
				roomMap[room[0]] = append(roomMap[room[0]], 0, roomX, roomY)
			} else if i == endPoint {
				roomMap[room[0]] = append(roomMap[room[0]], 2, roomX, roomY)
			} else if arguments[i][:2] != "##" {
				roomMap[room[0]] = append(roomMap[room[0]], 1, roomX, roomY)
			}
		}
	}

	//Checks for coordinate duplicates
	for key1, v1 := range roomMap {
		for key2, v2 := range roomMap {
			if key1 == key2 {
				break
			} else if v1[1] == v2[1] && v1[2] == v2[2] {
				err := errors.New("wrong coordinates in rooms: " + key1 + " and " + key2)
				errorCheck(err)
			}
		}
	}

	return roomMap
}

//maps connections between rooms
func mapRoomConnections(rawData []string, coordinatesMap map[string][]int) map[string][]string {
	var err error
	originalMap := make(map[string][]string)

	//adds room type as 1st string in slice
	for key, v := range coordinatesMap {
		switch v[0] {
		case 0:
			originalMap[key] = append(originalMap[key], "start")
		case 2:
			originalMap[key] = append(originalMap[key], "end")
		default:
			originalMap[key] = append(originalMap[key], "middle")
		}
	}

	//adds relations to the map
	for _, v := range rawData {
		connection := strings.Split(v, "-")
		if originalMap[connection[0]] != nil && originalMap[connection[1]] != nil {
			originalMap[connection[0]] = append(originalMap[connection[0]], connection[1])
			originalMap[connection[1]] = append(originalMap[connection[1]], connection[0])
		} else {
			err = errors.New("missing location")
			errorCheck(err)
		}
	}
	return originalMap
}
