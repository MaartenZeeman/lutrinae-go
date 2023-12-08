package puzzles

import (
	"strconv"
	"strings"
)

func SolveSixOne(input []string) int {
	var sum = 1

	times := strings.Fields(input[0])
	distances := strings.Fields(input[1])

	for timeIndex, time := range times {
		if timeIndex == 0 {
			continue
		}
		timeInt, _ := strconv.Atoi(time)
		distanceInt, _ := strconv.Atoi(distances[timeIndex])

		distance := getPossibleStrategies(timeInt, distanceInt)
		if distance > 0 {
			sum *= distance
		}
	}
	return sum
}

func SolveSixTwo(input []string) int {
	time := strings.Split(input[0], ":")
	distance := strings.Split(input[1], ":")

	timeInt, _ := strconv.Atoi(strings.Replace(time[1], " ", "", -1))
	distanceInt, _ := strconv.Atoi(strings.Replace(distance[1], " ", "", -1))

	sum := getPossibleStrategies(timeInt, distanceInt)

	return sum
}

// inefficient but still only takes milliseconds
func getPossibleStrategies(time int, distance int) int {
	sum := 0
	speed := 0
	for i := 0; i < time; i++ {
		if (time-i)*speed > distance {
			sum++
		} else if sum > 0 {
			break
		}
		speed++
	}
	return sum
}
