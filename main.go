package main

import (
	"fmt"
	"lutrinae/internal/infra"
	"lutrinae/internal/puzzles"
)

func main() {
	var input = infra.ReadInputFile("./input/2.aoc")
	maxAllowed := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var result = puzzles.SolveTwoOne(input, maxAllowed)
	fmt.Println(result)
}
