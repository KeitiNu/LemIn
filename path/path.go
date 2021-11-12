package main

import (
	"fmt"
)

var (
	anthill = make(map[string][]string)
	paths   [][]string
)

// Path takes in an anthill layout gives back the most optimal routh to the end
// func Path() [][]string {
func main() {
	// 0 - 1/2/3 - 4 - 5
	anthill["0"] = []string{"start", "1", "2", "3"}
	anthill["1"] = []string{"middle", "0", "4"}
	anthill["2"] = []string{"middle", "0", "4"}
	anthill["3"] = []string{"middle", "0", "4"}
	anthill["4"] = []string{"middle", "1", "2", "3", "5"}
	anthill["5"] = []string{"end", "4"}

	for start, options := range anthill {
		if options[0] == "start" {
			findWay(start, []string{})
			break
		}
	}

	fmt.Println(filter())
}

// findWay takes an anthill layout and gives us all the possible paths to the end
func findWay(room string, way []string) {
	options := anthill[room]
	way = append(way, room)

	// if we're at the end then add this path to the path list
	if options[0] == "end" {
		paths = append(paths, way)
		return
	}

	// skipping over the room type, we try to find a room we haven't travled yet
	// if we find such a room then we call findWay with the room info and traveled path
	for i := 1; i < len(options); i++ {
		for j, oldRoom := range way {
			if oldRoom == options[i] {
				break
			} else if j == len(way)-1 {
				findWay(options[i], way)
			}
		}
	}
}

// takes two roads, eachs if each room is unique
func filter() [][]string {
	var rightWay [][]string

	// if we only found one path, that's that
	if len(paths) == 1 {
		return paths
	}

	// sort the list, the shortest path at the front
	for i := 1; i < len(paths); i++ {
		if i < 1 {
			continue
		} else if len(paths[i-1]) > len(paths[i]) {
			paths[i], paths[i-1] = paths[i-1], paths[i]
			i = i - 2
		}
	}

	// searching for the right way...
	for _, short := range paths {
		// if we've ran out of short roads, the program will end
		if len(short) != len(paths[0]) {
			break
		}

		// way is the path that will store the current path
		way := [][]string{short}
		way = findBranchingPaths(short, way)

		if len(rightWay) == 0 {
			rightWay = way
		} else {
			// 1. tuleb vaadata mitu branchi kummaski on
			// 2. Kui palju sipelgaid on 
		}
	}

	return rightWay
}
// if number of rouths on < sipelgat arv - kasuta kõiki võimalikke teid ja kui sipelgate arv on sama või väiksem kui path siis kasutage kõige lühemat

// 2 	 = 2 liigutust - 1 = 5 liigutust - 4 = 6 liigutust - 5 = 7 liigutust - 6
// 5 - 3 = 5 liigutust - 3 = 5 liigutust - 3 = 6 liigutust - 5 = 7 liigutust - 7

func findBranchingPaths(short []string, way [][]string) [][]string {
	// in both middles, start and end get cut off for all roads share that
	middle1 := [][]string{short[1 : len(short)-2]}

	for _, long := range paths {
		middle2 := long[1 : len(long)-2]
		var breaker bool

		for _, path := range middle1 {
			// compared the paths
			for _, room1 := range path {
				for _, room2 := range middle2 {
					// if the roads cross then we break and a new long get compared
					if room2 == room1 { // breaking out of for ... middle2
						breaker = true
						break
					}
				}
				if breaker { // breaking out of for ... path
					break
				}
			}
			if breaker { // breaking out of for ... middle1
				breaker = false
				break
			}

			// If we made it all the way through without there being any matching rooms, then
			way = append(way, long)
			middle1 = append(middle1, middle2) // middle2 get appended so future roads will not cross it's path
		}
	}

	return way
}
