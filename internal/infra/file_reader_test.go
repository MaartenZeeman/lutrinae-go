package infra

import (
	"fmt"
	"testing"
)

func TestReadInputFileSub(t *testing.T) {
	var input = ReadInputFile("../../input/test.txt")
	var shouldBe = []string{"regel 1", "regel 2", "regel 3"}

	t.Run("Input length", func(t *testing.T) {
		if len(input) != len(shouldBe) {
			t.Error("Incorrent input length")
		}
	})
	t.Run("Input contents", func(t *testing.T) {
		for index, item := range input {
			if item != shouldBe[index] {
				t.Error(fmt.Sprintf("Incorrect input value. Expected '%s', found '%s'", shouldBe[index], item))
			}
		}
	})
}
