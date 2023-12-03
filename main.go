package main

import (
	"fmt"
	"lutrinae/internal/infra"
	"lutrinae/internal/puzzles"
)

func main() {
	printDayThreeResults()
}

func printDayOneResults() {
	var input = infra.ReadInputFile("./input/1.aoc")
	fmt.Printf("Result of 1-1 = %d \n", puzzles.SolveOneOne(input))
	fmt.Printf("Result of 1-2 = %d \n", puzzles.SolveOneTwo(input))
}

func printDayTwoResults() {
	var input = infra.ReadInputFile("./input/2.aoc")
	maxAllowed := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	fmt.Printf("Result of 2-1 = %d \n", puzzles.SolveTwoOne(input, maxAllowed))
	fmt.Printf("Result of 2-2 = %d \n", puzzles.SolveTwoTwo(input))
}

func printDayThreeResults() {
	var input = infra.ReadInputFile("./input/3.aoc")
	fmt.Printf("Result of 3-1 = %d \n", puzzles.SolveThreeOne(input))
	fmt.Printf("Result of 3-2 = %d \n", puzzles.SolveThreeTwo(input))
}
