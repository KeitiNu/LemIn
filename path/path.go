package path

var (
	anthill = make(map[string][]string)
	paths   [][]string
)

// Path takes in an anthill layout gives back the most optimal route to the end
// func Path()
func Path(data map[string][]string, ants int) ([][]string, []int) {
	anthill = data

	for start, options := range anthill {
		if options[0] == "start" {
			findWay([]string{start})
			break
		}
	}

	var way [][]string
	var distribution []int

	for end, options := range anthill {
		if options[0] == "end" {
			way, distribution = filter(end, ants)
		}
	}

	return way, distribution
}

// findWay takes an anthill layout and gives us all the possible paths to the end.
// Creadit to Zane who helped fix the info overwritting problem.
func findWay(way []string) {
	options := anthill[way[len(way)-1]]

	// if we're at the end then add this path to the path list
	if options[0] == "end" {
		paths = append(paths, way)
	} else {
		// skipping over the room type, we try to find a room we haven't travled yet
		// if we find such a room then we call findWay with the room info and traveled path
		// Options
	loop:
		for i := 1; i < len(options); i++ {
			// The path it took you to get here
			for _, oldRoom := range way {
				if oldRoom == options[i] {
					continue loop
				}
			}
			newPath := append(way, options[i])
			test := make([]string, len(newPath))
			for i := 0; i < len(newPath); i++ {
				test[i] = newPath[i]
			}
			findWay(test)

		}
	}
}

// takes two roads, eachs if each room is unique
func filter(endroom string, ants int) ([][]string, []int) {
	var (
		rightWay          [][]string
		rightDistribution []int
		rightMoves        int
	)

	// if we only found one path, that's that
	if len(paths) == 1 {
		return paths, []int{ants}
	}

	for i := 1; i < len(paths); i++ {
		if i < 1 {
			continue
		} else if len(paths[i]) < len(paths[i-1]) {
			paths[i], paths[i-1] = paths[i-1], paths[i]
			i -= 2
		}
	}

	var tempPath [][]string
	tempPath = append(tempPath, paths...)

	// searching for the right way...
loop:
	for _, short := range paths {

		// if we find a road that's only 2 long,
		// then we have a path that connects straight to the end
		if len(short) == 2 {
			return [][]string{short}, []int{ants}
		}

		// way is the path that will store the current path
		// way := findBranchingPaths(short[1:len(short)-1], [][]string{short}, [][][]string{})
		way := findBranchingPaths(short[1:len(short)-1], [][]string{short}, tempPath)
		for _, arr := range way {
			if arr[len(arr)-1] != endroom {
				continue loop
			}
		}

		if rightMoves < 1 {
			rightWay, rightDistribution, rightMoves = formula(endroom, way, ants)
		} else {
			newWay, newDistribution, newMoves := formula(endroom, way, ants)
			if newMoves < rightMoves {
				rightDistribution, rightMoves, rightWay = newDistribution, newMoves, newWay
			}
		}
	}

	return rightWay, rightDistribution
}

func findBranchingPaths(middle1 []string, way [][]string, tempPath [][]string) [][]string {
	var breaker bool
	for _, long := range tempPath {
		middle2 := long[1 : len(long)-1]

		for i, room1 := range middle1 {
			// compared the paths
			for _, room2 := range middle2 {
				// if the roads cross then we break and a new long get compared
				if room2 == room1 { // breaking out of for ... middle2
					breaker = true
					break
				}
			}
			if breaker {
				breaker = false
				break
			}

			if i == len(middle1)-1 {
				// If we made it all the way through without there being any matching rooms, then
				way = append(way, long)
				// middle1 = append(middle1, middle2...)
				mid := make([]string, len(middle1)+len(middle2))
				i := 0
				for i < len(middle1) {
					mid[i] = middle1[i]
					i++
				}

				for j := 0; j < len(middle2); j++ {
					mid[i] = middle2[j]
					i++
				}

				middle1 = mid
			}
		}
	}
	return way
}

// a formula to send the right amount of ants down each path, return the distribution and the move count
func formula(endroom string, option [][]string, ants int) ([][]string, []int, int) {

	finished, distribution := moveAnts(option)
	moves := len(option[len(option)-1]) - 1

	for _, arr := range option {
		if len(arr)-1 > moves {
			moves = len(arr) - 1
		}
	}

	// if the way we distributed the ants is greater than the amount of ants we have...
	// ...then we're fucked...
	// if it's the same then we got off easy.
	if finished > ants {
		return subtraction(option, ants, finished, distribution, moves)
	} else if finished == ants {
		return option, distribution, moves
	}

	// start : we send out the beginning path of unevenly distributed ants
	ants = ants - finished

	// middle/end : now that the uneven part is done then we
	base := make([]int, len(distribution))
	copy(base, distribution)

	if len(distribution) == 1 {
		distribution[0] += ants
		moves += ants
	} else {
		for i := 0; i < len(distribution); i++ {
			if i == 0 {
				moves++
			}
			distribution[i]++
			ants--
			if ants == 0 {
				break
			}

			if ants > 0 && i == len(distribution)-1 {
				i = -1
			} else if ants == 0 {
				break
			}
		}
	}

	return option, distribution, moves
}

// moveAnts takes in a given way
// and return the base amout of ants that finish
// and how those ant are divided to each road.
func moveAnts(way [][]string) (int, []int) {
	if len(way) == 1 {
		return 1, []int{1}
	}

	var distribution []int              // the number of ants that are sent down each path
	var antsfinished int                // ants that finish by x number of moves
	longestPath := len(way[len(way)-1]) // length of the longest path

	for _, arr := range way {
		if len(arr) > longestPath {
			longestPath = len(arr)
		}
	}

	for i := range way {
		antsfinished += longestPath - len(way[i]) + 1
		distribution = append(distribution, longestPath-len(way[i])+1)
	}

	return antsfinished, distribution
}

// if the way we distributed the ants is greater than the amount of ants we have
// then we start subtracting them from te roads
func subtraction(option [][]string, ants int, finished int, distribution []int, moves int) ([][]string, []int, int) {
	for i := len(distribution) - 1; i > -1; i-- {
		distribution[i] -= 1
		finished--

		// if we've taken all the ants off one road, then we remove it
		if distribution[i] == 0 {
			tempDis, tempWay := distribution[:i], option[:i]
			tempDis, tempWay = append(tempDis, distribution[i+1:]...), append(tempWay, option[i+1:]...)
			distribution, option = tempDis, tempWay
			i--
		}

		// if finally we have the right amount of ants, we end the program
		if finished == ants {
			if i == 0 {
				moves--
			}
			return option, distribution, moves
		}

		// if we took 1 level off the ants and there's still too much
		// we remove a move and start the process over again
		if i == 0 {
			moves--
			i = len(distribution) - 1
		}
	}
	return option, distribution, moves

}
