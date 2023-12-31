package main

import (
	"fmt"
	"lutrinae/internal/infra"
	"lutrinae/internal/puzzles"
)

func main() {
	printDayEightResults()
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

func printDayFourResults() {
	var input = infra.ReadInputFile("./input/4.aoc")
	fmt.Printf("Result of 4-1 = %d \n", puzzles.SolveFourOne(input))
	fmt.Printf("Result of 4-2 = %d \n", puzzles.SolveFourTwo(input))
}

func printDayFiveResults() {
	var input = infra.ReadInputFile("./input/5.aoc")
	fmt.Printf("Result of 5-1 = %d \n", puzzles.SolveFiveOne(input))
	fmt.Printf("Result of 5-2 = %d \n", puzzles.SolveFiveTwo(input))
}

func printDaySixResults() {
	var input = infra.ReadInputFile("./input/6.aoc")
	fmt.Printf("Result of 6-1 = %d \n", puzzles.SolveSixOne(input))
	fmt.Printf("Result of 6-2 = %d \n", puzzles.SolveSixTwo(input))
}

func printDaySevenResults() {
	var input = infra.ReadInputFile("./input/7.aoc")
	fmt.Printf("Result of 7-1 = %d \n", puzzles.SolveSevenOne(input))
	fmt.Printf("Result of 7-2 = %d \n", puzzles.SolveSevenTwo(input))
}

func printDayEightResults() {
	var input = infra.ReadInputFile("./input/8.aoc")
	fmt.Printf("Result of 8-1 = %d \n", puzzles.SolveEightOne(input))
	fmt.Printf("Result of 8-2 = %d \n", puzzles.SolveEightTwo(input))
}
