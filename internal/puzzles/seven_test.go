package puzzles

import (
	"fmt"
	"testing"
)

func TestSolveSeveneOne(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	const shouldBe = 6440
	var result = SolveSevenOne(input)

	if shouldBe != result {
		t.Error(fmt.Sprintf("Expected %d, found %d", shouldBe, result))
	}
}
