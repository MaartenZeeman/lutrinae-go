package puzzles

import (
	"math"
	"strings"
)

func SolveFourOne(input []string) int {
	var sum float64 = 0
	for _, line := range input {
		//de-rubbish-fy the line by spliting on the :. [0] is rubbish, [1] contains the numbers
		splitLine := strings.Split(line, ":")
		//split into winning numbers and our numbers
		splitCard := strings.Split(splitLine[1], "|")

		// split both parts (numbers) on whitespace
		winningNumbers := strings.Fields(splitCard[0])
		ourNumbers := strings.Fields(splitCard[1])

		// go doesn't have a 'contains' function, so we're adding all of our numbers in a map to make it easy to check for winning numbers later
		// the ticket probably doesn't have duplicate numbers, but this supports it anyway
		ourNumbersMap := make(map[string]int)
		for _, number := range ourNumbers {
			ourNumbersMap[number] += 1
		}

		amountOfWinningNumbers := 0
		for _, number := range winningNumbers {
			amountOfWinningNumbers += ourNumbersMap[number]
		}
		// first number is 1 point and it doubles for every additional match. So 1, 2, 4, 8, etc. So we do 2 ^ (amount of matches - 1)
		if amountOfWinningNumbers > 0 {
			sum += math.Pow(2, float64(amountOfWinningNumbers-1))
		}
	}

	return int(sum)
}

func SolveFourTwo(input []string) int {
	var sum = 0
	// a map that contains the number of cards we have for each 'card'/id whatever. It's late
	amountOfCards := make(map[int]int)

	// I don't like it, but the default value should be 1 (because we start with 1 of each card).
	// this can't be done in the next loop because we need to set it for 'future' cards as well in case we win them before we had a chance to set the value
	for lineIndex, _ := range input {
		amountOfCards[lineIndex] = 1
	}

	for lineIndex, line := range input {
		if amountOfCards[lineIndex] == 0 {
			amountOfCards[lineIndex] = 1
		}
		sum += amountOfCards[lineIndex]
		//de-rubbish-fy the line by spliting on the :. [0] is rubbish, [1] contains the numbers
		splitLine := strings.Split(line, ":")
		//split into winning numbers and our numbers
		splitCard := strings.Split(splitLine[1], "|")

		// split both parts (numbers) on whitespace
		winningNumbers := strings.Fields(splitCard[0])
		ourNumbers := strings.Fields(splitCard[1])

		// go doesn't have a 'contains' function, so we're adding all of our numbers in a map to make it easy to check for winning numbers later
		// the ticket probably doesn't have duplicate numbers, but this supports it anyway
		ourNumbersMap := make(map[string]int)
		for _, number := range ourNumbers {
			ourNumbersMap[number] += 1
		}

		amountOfWinningNumbers := 0
		for _, number := range winningNumbers {
			amountOfWinningNumbers += ourNumbersMap[number]
		}
		// add the amount of cards we have of this one (because thats how many) to the next amountOfWinningNumbers cards.
		if amountOfWinningNumbers > 0 {
			for i := lineIndex + 1; i <= lineIndex+int(amountOfWinningNumbers); i++ {
				amountOfCards[i] += amountOfCards[lineIndex]
			}
		}
	}

	return sum
}
