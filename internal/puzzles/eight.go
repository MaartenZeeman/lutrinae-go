package puzzles

var route = []string{}
var crossroads = make(map[string]map[string]string)

const START = "AAA"
const END = "ZZZ"

func SolveEightOne(input []string) int {
	readDirections(input[0])
	readCrossroads(input[2:len(input)])
	return calculateAmountOfSteps(START, 0)
}

func SolveEightTwo(input []string) int64 {
	readDirections(input[0])
	readCrossroads(input[2:len(input)])
	return calculateAmountOfStepsForGhosts(getGhostsStartingLocations(), 0)
}

func readDirections(directions string) {
	route = []string{}
	for _, direction := range directions {
		route = append(route, string(direction))
	}
}

func readCrossroads(input []string) {
	crossroads = make(map[string]map[string]string)
	for _, line := range input {
		location := line[0:3]
		left := line[7:10]
		right := line[12:15]
		crossroads[location] = make(map[string]string)
		crossroads[location]["L"] = left
		crossroads[location]["R"] = right
	}
}

func calculateAmountOfSteps(location string, steps int) int {
	for _, direction := range route {
		steps++
		location = crossroads[location][direction]

		if location == END {
			return steps
		}
	}
	return calculateAmountOfSteps(location, steps)
}

func getGhostsStartingLocations() []string {
	locations := []string{}
	for location, _ := range crossroads {
		if string(location[2]) == "A" {
			locations = append(locations, location)
		}
	}
	return locations
}

func calculateAmountOfStepsForGhosts(locations []string, steps int64) int64 {
	for _, direction := range route {
		allFinished := true
		steps++
		for locIndex, location := range locations {
			locations[locIndex] = crossroads[location][direction]
			if string(locations[locIndex][2]) != "Z" {
				allFinished = false
			}
		}

		if allFinished {
			return steps
		}
	}
	return calculateAmountOfStepsForGhosts(locations, steps)
}
