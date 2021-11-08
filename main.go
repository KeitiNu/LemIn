package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Returns an error if err!=nil
func errorCheck(e error) {
	if e != nil {
		fmt.Println("ERROR: invalid data format")
		log.Fatal(e)
	}
}

//Maps rooms and their connections
func mapRoomCoordinates(arguments []string) {
	var startPoint int
	var endPoint int
	roomMap := make(map[string][]int)

	for i := 0; i < len(arguments); i++ {
		if arguments[i] == "##start" {
			startPoint = i + 1
		} else if arguments[i] == "##end" {
			endPoint = i + 1
			break
		}
	}

	for _, value := range arguments[startPoint : endPoint+1] {
		if value[0] != '#' {
			room := strings.Split(value, " ")
			if len(room) == 3 {
				roomX, err := strconv.Atoi(room[1])
				errorCheck(err)
				roomY, err := strconv.Atoi(room[2])
				errorCheck(err)
				roomMap[room[0]] = append(roomMap[room[0]], roomX, roomY)
			}else{
				err := errors.New("wrong number of arguments in a room")
				errorCheck(err)
			}
		}
	}

	//JUST A TEST PRINT
	for key, v := range roomMap {
		fmt.Println(key, v)
	}

}

func main() {

	//Checks, if the arguments have been given correctly
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [INPUTFILE]")
		fmt.Println("EX: go run . example00.txt")
	}

	//Reads in the file from arguments
	fileName := "./examples/" + os.Args[1]
	input, err := os.ReadFile(fileName)
	errorCheck(err)

	arguments := strings.Split(string(input), "\n")

	//Finds the number of ants
	var nrOfAnts int
	for i := 0; i < len(arguments); i++ {
		if arguments[i][0] != '#' {
			nrOfAnts, err = strconv.Atoi(arguments[i])
			errorCheck(err)
			break
		}
	}

	fmt.Println("Number of ants: ", nrOfAnts) //TEST PRINT

	//Maps the rooms and their coordinates
	mapRoomCoordinates(arguments)
}
