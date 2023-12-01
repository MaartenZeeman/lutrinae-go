package puzzles

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SolveFirst(input []string) int {
	var sum = 0
	// Loop through input slice and add the result of each row (item) to 'sum'
	for _, item := range input {
		sum += getFirstAndLastDigits(item)
	}

	return sum
}

func SolveSecond(input []string) int {
	var sum = 0
	// SolveFirst with extra steps
	for _, item := range input {
		// Replace spelled out numbers with actual numbers. Because we are replacing the letters, we need to replace the 'typos' like "oneight" first to make sure we get both the 1 and 8.
		r := strings.NewReplacer("oneight", "18", "twone", "21", "eightwo", "82", "one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
		var fixedItem = r.Replace(item)
		sum += getFirstAndLastDigits(fixedItem)
	}

	return sum
}

func getFirstAndLastDigits(input string) int {
	var result = 0
	// Put each digit as an item in a slice, grab the first and last, put them in a string and then convert it to an integer
	re := regexp.MustCompile("[\\d]{1}")
	var digits = re.FindAllString(input, -1)
	if len(digits) > 0 {
		result, _ = strconv.Atoi(fmt.Sprintf("%s%s", digits[0], digits[len(digits)-1]))
	}
	return result
}
