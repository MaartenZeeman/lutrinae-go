package puzzles

import (
	"fmt"
	"testing"
)

func TestSolveEightOne(t *testing.T) {
	input := []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	const shouldBe = 6
	var result = SolveEightOne(input)

	if shouldBe != result {
		t.Error(fmt.Sprintf("Expected %d, found %d", shouldBe, result))
	}
}
