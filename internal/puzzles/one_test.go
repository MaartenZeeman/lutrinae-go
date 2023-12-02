package puzzles

import (
	"fmt"
	"testing"
)

func TestSolveOneOne(t *testing.T) {
	var input = []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	var shouldBe = 142
	var result = SolveOneOne(input)

	if result != shouldBe {
		t.Error(fmt.Sprintf("Wrong answer. Expected %d, got %d", shouldBe, result))
	}
}

func TestSolveOneTwo(t *testing.T) {
	var input = []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	var shouldBe = 281
	var result = SolveOneTwo(input)

	if result != shouldBe {
		t.Error(fmt.Sprintf("Wrong answer. Expected %d, got %d", shouldBe, result))
	}
}
