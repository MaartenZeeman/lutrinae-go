package puzzles

import (
	"strconv"
	"strings"
)

func SolveTwoOne(input []string, maxAllowed map[string]int) int {
	var sum = 0

	for _, item := range input {
		// First split the input on the colon so we seperate the ID and actual input
		inputSplit := strings.Split(item, ":")
		// If the length of the slice isn't exactly 2, the data is bwoken, IT'S BWOKEN!
		if len(inputSplit) != 2 {
			continue
		}
		// We can use a regex like yesterday or remove the word 'game', but using a substring is more fun
		gameID, _ := strconv.Atoi(inputSplit[0][5:len(inputSplit[0])])
		// Next, split the other part, the actual input, on the semicolon.
		// Each of the resulting items is one handful of balls (or whatever they are called in the description)
		draws := strings.Split(inputSplit[1], ";")

		isPossible := true
		for _, draw := range draws {
			// More splitting, now on the comma to get the different colors and quantities
			colors := strings.Split(draw, ",")
			// We are going to put the colors and number of balls in a map.
			// The easier and probably faster way is to split on the keys in the maxAllowed input and get the number in front of it,
			// but I want to mess with maps a bit more
			var colorMap map[string]int
			colorMap = make(map[string]int)
			for _, color := range colors {
				// Moar splitting! Split on whitespace to get the color and number of balls. Trim leading and trailing whitespace first
				colorAndQuantity := strings.Split(strings.TrimSpace(color), " ")
				colorMap[colorAndQuantity[1]], _ = strconv.Atoi(colorAndQuantity[0])
			}
			// Loop through the map with maximum number of balls per color.
			// If the draw exceeds any of them, set isPossible to false and break the loop
			for key, value := range maxAllowed {
				if colorMap[key] > value {
					isPossible = false
					break
				}
			}
			// No need to check the next draw if the previous one was impossible
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
