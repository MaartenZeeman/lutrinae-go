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

func TestSolveEightTwo(t *testing.T) {
	input := []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}

	const shouldBe = 6
	var result = SolveEightTwo(input)

	if shouldBe != result {
		t.Error(fmt.Sprintf("Expected %d, found %d", shouldBe, result))
	}
}
