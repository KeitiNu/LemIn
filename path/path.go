package path

import "fmt"

var (
	anthill = make(map[string][]string)
	paths   [][]string
)

// Path takes in an anthill layout gives back the most optimal route to the end
// func Path()
func Path(data map[string][]string, ants int) ([][]string, []int) {
	/*
		// 0 - 1/2/3 - 4 - 5
		ants := 10
		anthill["0"] = []string{"start", "1", "2", "3"}
		anthill["1"] = []string{"middle", "0", "4"}
		anthill["2"] = []string{"middle", "0", "4"}
		anthill["3"] = []string{"middle", "0", "4"}
		anthill["4"] = []string{"middle", "1", "2", "3", "5"}
		anthill["5"] = []string{"end", "4"}
	*/

	anthill = data

	fmt.Println(anthill)

	for start, options := range anthill {
		if options[0] == "start" {
			findWay([]string{start})
			break
		}
	}

	way, distribution := filter(ants)

	return way, distribution
}

// findWay takes an anthill layout and gives us all the possible paths to the end.
// Creadit to Zane who helped fix the info overwritting problem.
func findWay(way []string) {
	options := anthill[way[len(way)-1]]

	// if we're at the end then add this path to the path list
	if options[0] == "end" {
		paths = append(paths, way)

		// fmt.Printf("Found %v\n", paths)
		// fmt.Println()
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
			// fmt.Printf("Third %v\n", paths)
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
func filter(ants int) ([][]string, []int) {
	var (
		rightWay          [][]string
		rightDistribution []int
		rightMoves        int
	)

	// if we only found one path, that's that
	if len(paths) == 1 {
		return paths, []int{ants}
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

		// if we find a road that's only 2 long,
		// then we have a path that connects straight to the end
		if len(short) == 2 {
			return [][]string{short}, []int{ants}
		}

		// way is the path that will store the current path
		way := [][]string{short}
		way = findBranchingPaths(short, way)

		if len(rightWay) == 0 {
			rightWay, rightDistribution, rightMoves = formula(way, ants)
		} else {
			newWay, newDistribution, newMoves := formula(way, ants)
			if newMoves > rightMoves {
				rightDistribution, rightMoves, rightWay = newDistribution, newMoves, newWay
			}
		}
	}

	return rightWay, rightDistribution
}

func findBranchingPaths(short []string, way [][]string) [][]string {
	// in both middles, start and end get cut off for all roads share that
	middle1 := [][]string{short[1 : len(short)-1]}

	for _, long := range paths {
		middle2 := long[1 : len(long)-1]
		var breaker bool

		for i, path := range middle1 {
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

			if i == len(middle1)-1 {
				// If we made it all the way through without there being any matching rooms, then
				way = append(way, long)
				middle1 = append(middle1, middle2) // middle2 get appended so future roads will not cross it's path
			}
		}
	}
	return way
}

// a formula to send the right amount of ants down each path, return the distribution and the move count
func formula(option [][]string, ants int) ([][]string, []int, int) {
	finished, distribution := moveAnts(option)
	roadCount := len(option)

	fmt.Printf("raw distribution: %v\n", distribution)

	// if the way we distributed the ants is greater than the amount of ants we have...
	// ...then we're fucked...
	// if it's the same then we got off easy.
	if finished > ants {
		// fmt.Printf("%v > %v\n", finished, ants)

		// moves := len(option[len(option)-1])
		// return subtraction(option, ants, finished, distribution, moves)
	} else if finished == ants {
		return option, distribution, len(option[len(option)-1])
	}

	// start : we send out the beginning path of unevenly distributed ants
	moves := len(option[len(option)-1])
	ants = ants - finished

	// middle/end : now that the uneven part is done then we
	moves += ants / roadCount
	ants = ants % roadCount
	for i := range distribution {
		distribution[i] += ants / roadCount
	}

	// end : if theres still some ants lingering then we do one extra move for them
	if ants > 0 {
		moves++
		for i := range distribution {
			distribution[i] += 1
			ants--
			if ants == 0 {
				break
			}
		}
	}

	// var sum int
	// for _, v := range distribution {
	// 	sum += v
	// }
	fmt.Printf("way : %v\n", option)
	fmt.Printf("distribution: %v\n", distribution)

	return option, distribution, moves
}

// moveAnts takes in a given way and the number of ants
// and return how many ants finished base on how long
// it'll take for the ant on the longest path to get home
// and how those ant are divided to each road.
func moveAnts(way [][]string) (int, []int) {
	if len(way) == 1 {
		return len(way[0]), []int{1}
	}

	var distribution []int              // the number of ants that are sent down each path
	var antsfinished int                // ants that finish by x number of moves
	longestPath := len(way[len(way)-1]) // length of the longest path
	antsfinished++

	for i := 0; i < len(way)-2; i++ {
		antsfinished += longestPath - len(way[i]) + 1
		distribution = append(distribution, longestPath-len(way[i])+1)
	}

	distribution = append(distribution, 1)

	return antsfinished, distribution
}

/*
// if the way we distributed the ants is greater than the amount of ants we have
// then we start subtracting them from te roads
func subtraction(option [][]string, ants int, finished int, distribution []int, moves int) ([][]string, []int, int) {
	for i := 0; i < len(distribution); i++ {
		distribution[i] -= 1
		finished--

		// need to make it soo that the option also loses it's value in that case...
		if distribution[i] == 0 {
			tempDis, tempWay := distribution[:i], option[:i]
			tempDis, tempWay = append(tempDis, distribution[i+1:]...), append(tempWay, option[i+1:]...)
			distribution, option = tempDis, tempWay
			i--
		}

		if finished == ants {
			if i == len(distribution)-1 {
				moves--
			}
			return option, distribution, moves
		} else if finished < ants {
			fmt.Println("error in subtraction")
			os.Exit(0)
		}
	}
	moves--
	return option, distribution, moves
	// return subtraction(option, ants, finished, distribution, moves)
}
*/

/*
	ants = 100 - 3 = 97
	moves = 0 + 5 = 5

	[97 / 2 = 48]
	moves = 5 + 48 = 53
	ants = 97 % 2 = 2

	if ants(2) > 0 {
		moves++
	}

	to find the final amount of moves that's needed to finish.

	first we need to find how many ants will finish by the time the longest ant takes to finish
	we save these as 2 moves and 3 finished.

	firstfinish = 3, longest road length = 5
	ants = 100 moves = 0

	100 - 3 = 97
	moves = 0 + 5 = 5

	ants = 97 moves = 5

	now we have the base number of ants we will start working with. After the first iteration of ants is home,
	then we just need to send one after the other down each path, meaning, if we divide the base number of ants
	we got with how many roads we have, we have the middle part of the moving process done.

	97 / 2 = 48
	moves = 5 + 48 = 53
	ants = 97 - 48 * 2 = 97-96 = 2
*/

// if number of rouths on < sipelgat arv - kasuta kõiki võimalikke teid ja kui sipelgate arv on sama või väiksem kui path siis kasutage kõige lühemat

/*

if 10
2 	  = 2 liigutust - 1 = 5 liigutust - 4 = 6 liigutust - 5 = 7 liigutust - 6 = 8 liigutust - 7
5 - 3 = 5 liigutust - 3 = 5 liigutust - 3 = 6 liigutust - 5 = 7 liigutust - 7 = 8 liigutust - 9

100

numberOfAnts - baseFinished
resultOfPrevious / numpaths

100 - 3 = 97
97 / 2 = 48
if 97 % 2 > 0 {
	97 % 2 = 1
	48 + 1 % 2 + 1 / 2 = 49
}
49 + 5 = 54

100 - 1 = 99
99 / 1 = 99
if...
99 + 2 = 101

*/
