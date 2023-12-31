package puzzles

import (
	"strconv"
	"strings"
)

func SolveTwoOne(input []string, maxAllowed map[string]int) int {
	var sum = 0

	for _, item := range input {
		gameID, hands := splitInput(item)
		isPossible := true
		for _, hand := range hands {
			colorMap := getColorMap(hand)
			// Loop through the map with maximum number of cubes per color.
			// If the hand exceeds any of them, set isPossible to false and break the loop
			for key, value := range maxAllowed {
				if colorMap[key] > value {
					isPossible = false
					break
				}
			}
			// No need to check the next hand if the previous one was impossible
			if !isPossible {
				break
			}
		}
		if isPossible {
			sum += gameID
		}
	}

	return sum
}

func SolveTwoTwo(input []string) int {
	var sum = 0

	for _, item := range input {
		_, hands := splitInput(item)
		// a map to keep track of the highest amount of cubes we have of a color in 1 round
		maxPerColor := map[string]int{}
		for _, hand := range hands {
			colorMap := getColorMap(hand)
			// Replace the number in maxPerColor if the current amount of cubes for the specific color is higher.
			for key, value := range colorMap {
				if value > maxPerColor[key] {
					maxPerColor[key] = value
				}
			}
		}
		// handPower sounds cool. Keep multiplying with the next color
		handPower := 1
		for _, value := range maxPerColor {
			handPower *= value
		}
		sum += handPower
	}

	return sum
}

func splitInput(input string) (int, []string) {
	// First split the input on the colon so we seperate the ID and actual input
	inputSplit := strings.Split(input, ":")
	// If the length of the slice isn't exactly 2, the data is bwoken, IT'S BWOKEN!
	if len(inputSplit) != 2 {
		return 0, []string{}
	}

	// We can use a regex like yesterday, split on whitespace or remove the word 'game', but using a substring is more fun and has magic numbers
	gameID, _ := strconv.Atoi(inputSplit[0][5:len(inputSplit[0])])
	// Next, split the other part, the actual input, on the semicolon.
	// Each of the resulting items is one handful of cubes (or whatever they are called in the description)
	hands := strings.Split(inputSplit[1], ";")

	return gameID, hands
}

func getColorMap(hand string) map[string]int {
	// More splitting, now on the comma to get the different colors and quantities
	colors := strings.Split(hand, ",")
	// We are going to put the colors and number of cubes in a map.
	var colorMap map[string]int
	colorMap = make(map[string]int)
	for _, color := range colors {
		// Moar splitting! Split on whitespace to get the color and number of cubes. Trim leading and trailing whitespace first
		colorAndQuantity := strings.Split(strings.TrimSpace(color), " ")
		colorMap[colorAndQuantity[1]], _ = strconv.Atoi(colorAndQuantity[0])
	}

	return colorMap
}
