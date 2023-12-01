package infra

import (
	"bufio"
	"log"
	"os"
)

// Read a file and append each row to a slice. Return that slice. Found the code online, looks fancy.
func ReadInputFile(filename string) []string {
	var output = []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}
