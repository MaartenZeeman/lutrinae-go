package puzzles

import (
	"fmt"
	"testing"
)

func TestSolveSixeOne(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	const shouldBe = 288
	var result = SolveSixOne(input)

	if shouldBe != result {
		t.Error(fmt.Sprintf("Expected %d, found %d", shouldBe, result))
	}
}

func TestSolveSixeTwo(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	const shouldBe = 71503
	var result = SolveSixTwo(input)

	if shouldBe != result {
		t.Error(fmt.Sprintf("Expected %d, found %d", shouldBe, result))
	}
}
