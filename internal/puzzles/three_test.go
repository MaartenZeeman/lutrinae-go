package puzzles

import (
	"fmt"
	"testing"
)

func TestSolveThreeOne(t *testing.T) {
	input := []string{
		"*467..114.",
		"..........",
		"..35...633",
		".*.....#..",
		"-617.......",
		".....+.58.",
		"..592.....",
		"...4-755..",
		"$.........",
		".660..%598",
		"..*.......",
		".1.1......",
		".......1!1",
		"1........1",
		"1......^1#",
		"1@........",
	}

	const shouldBe = 4369
	var result = SolveThreeOne(input)

	if shouldBe != result {
		t.Error(fmt.Sprintf("Expected %d, found %d", shouldBe, result))
	}
}

func TestSolveThreeTwo(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	const shouldBe = 467835
	var result = SolveThreeTwo(input)

	if shouldBe != result {
		t.Error(fmt.Sprintf("Expected %d, found %d", shouldBe, result))
	}
}
