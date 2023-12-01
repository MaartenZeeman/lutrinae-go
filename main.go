package main

import (
	"fmt"
	"lutrinae/internal/infra"
	"lutrinae/internal/puzzles"
)

func main() {
	var input = infra.ReadInputFile("./input/1.aoc")
	var result = puzzles.SolveSecond(input)
	fmt.Println(result)
}
