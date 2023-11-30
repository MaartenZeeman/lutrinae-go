package main

import (
	"fmt"
	"lutrinae/internal/infra"
)

func main() {
	var textInput = infra.ReadInputFile("./input/test.txt")
	fmt.Println(textInput)
}
