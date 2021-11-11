package path

import (
	"fmt"
)

var (
	anthill = make(map[string][]string)
 	paths 	[][]string
)

// Path takes in an anthill layout gives back the most optimal routh to the end
func Path() []string {
	// 0 - 1/2/3 - 4 - 5 to the end
	anthill["0"] = []string{"start", "1", "2", "3"}
	anthill["1"] = []string{"middle", "0", "4"}
	anthill["2"] = []string{"middle", "0", "4"}
	anthill["3"] = []string{"middle", "0", "4"}
	anthill["4"] = []string{"middle", "1", "2", "3", "5"}
	anthill["5"] = []string{"end", "4"}

	for start, rooms := range anthill {
		if rooms[0] == "start" {
			findWay(start, []string{})
		}
	}
	fmt.Println(paths)
	
	return filter()
}

func findWay(room string, way []string) {
	options := anthill[room]
	way = append(way, room)
	
	// if we're at the end then add this path to the path list
	if options[0] == "end" {
		paths = append(paths, way)
		return
	}

	// skipping over the room type, we try to find a room we haven't travled yet
	// if we find such a room then we call findWay with the room info and extended path
	for i := 1; i < len(options); i++ {
		for j, oldRoom := range way {
			if oldRoom == options[i] {
				break
			} else if j == len(way) - 1 {
				findWay(options[i], way)
			}
		}
	}
}

func filter() []string {
	var str []string

	// sort the list, the shortest path at the front

	return str
}