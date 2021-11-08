package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func errorCheck(e error) {
	if e != nil {
		fmt.Println("ERROR: invalid data format")
		log.Fatal(e)
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

	fmt.Println(nrOfAnts)//Just to use the variable

	//Maps the rooms and their coordinates

}
