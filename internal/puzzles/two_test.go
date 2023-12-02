package puzzles

import (
	"fmt"
	"testing"
)

func TestSolveTwoOne(t *testing.T) {
	var input = []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	//12 red cubes, 13 green cubes, and 14 blue
	maxAllowed := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var shouldBe = 8
	var result = SolveTwoOne(input, maxAllowed)

	if result != shouldBe {
		t.Error(fmt.Sprintf("Wrong answer. Expected %d, got %d", shouldBe, result))
	}
}
