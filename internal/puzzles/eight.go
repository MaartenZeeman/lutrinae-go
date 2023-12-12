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

func readDirections(directions string) {
	for _, direction := range directions {
		route = append(route, string(direction))
	}
}

func readCrossroads(input []string) {
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
	currentLocation := location
	for _, direction := range route {
		steps++
		currentLocation = crossroads[currentLocation][direction]

		if currentLocation == END {
			return steps
		}
	}
	return calculateAmountOfSteps(currentLocation, steps)
}
