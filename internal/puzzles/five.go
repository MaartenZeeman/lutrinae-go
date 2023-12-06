package puzzles

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type almanac struct {
	from   int64
	to     int64
	ranges int64
}

// Checks whether any of the given indices borders this partnumber
func (almanac *almanac) contains(number int64) bool {
	return number >= almanac.from && number < (almanac.from+almanac.ranges)
}
func (almanac *almanac) getCorrespondingNumber(number int64) int64 {
	return number - almanac.from + almanac.to
}

var seeds = []int64{}
var seedsToSoil = []almanac{}
var soilToFertilizer = []almanac{}
var fertilizerToWater = []almanac{}
var waterToLight = []almanac{}
var lightToTemperature = []almanac{}
var temperatureToHumidity = []almanac{}
var humidityToLocation = []almanac{}

func SolveFiveOne(input []string) int64 {
	readInput(input)

	return getLowestSeedNumber()
}

func SolveFiveTwo(input []string) int64 {
	readInput(input)

	return getLowestSeedNumberOnSteroids()
}

func readInput(input []string) {
	var currentMap *[]almanac = &seedsToSoil
	for _, line := range input {
		if strings.HasPrefix(line, "seeds:") {
			seeds = []int64{}
			tempSeeds := strings.Split(strings.Replace(strings.TrimSpace(line), "seeds: ", "", -1), " ")

			for _, seed := range tempSeeds {
				seedNumber, _ := strconv.ParseInt(seed, 10, 64)
				seeds = append(seeds, seedNumber)
			}
		} else if len(strings.TrimSpace(line)) == 0 {
			continue
		} else if line[0] >= ASCII_FIRST_DIGIT && line[0] <= ASCII_LAST_DIGIT {
			var to, from, amount int64
			fmt.Sscanf(line, "%d %d %d", &to, &from, &amount)

			*currentMap = append(*currentMap, almanac{from, to, amount})
		} else if strings.HasSuffix(line, "map:") {
			switch line {
			case "seed-to-soil map:":
				currentMap = &seedsToSoil
				break
			case "soil-to-fertilizer map:":
				currentMap = &soilToFertilizer
				break
			case "fertilizer-to-water map:":
				currentMap = &fertilizerToWater
				break
			case "water-to-light map:":
				currentMap = &waterToLight
				break
			case "light-to-temperature map:":
				currentMap = &lightToTemperature
				break
			case "temperature-to-humidity map:":
				currentMap = &temperatureToHumidity
				break
			case "humidity-to-location map:":
				currentMap = &humidityToLocation
				break
			}
			*currentMap = []almanac{}
		}

	}
}

// loop through the seeds and test each one
func getLowestSeedNumber() int64 {
	var lowest int64 = math.MaxInt64
	for _, seed := range seeds {
		soilNumber := testSeed(seed)

		if soilNumber < lowest {
			lowest = soilNumber
		}
	}

	return lowest
}

// Similar, but now loop through everything between the even and odd (i+1) seeds. Perfect to test your CPU fan
func getLowestSeedNumberOnSteroids() int64 {
	var lowest int64 = math.MaxInt64
	for seedIndex, seed := range seeds {
		if seedIndex%2 == 1 {
			continue
		}
		for i := seed; i < seed+seeds[seedIndex+1]; i++ {
			soilNumber := testSeed(i)
			if soilNumber < lowest {
				lowest = soilNumber
			}
		}
	}

	return lowest
}

// loopception. Loop through the loops until we no longer know what the loop is going on
func testSeed(seed int64) int64 {
	nextNumber := seed
	for _, sts := range seedsToSoil {
		if sts.contains(nextNumber) {
			nextNumber = sts.getCorrespondingNumber(nextNumber)
			break
		}
	}

	for _, sts := range soilToFertilizer {
		if sts.contains(nextNumber) {
			nextNumber = sts.getCorrespondingNumber(nextNumber)
			break
		}
	}

	for _, sts := range fertilizerToWater {
		if sts.contains(nextNumber) {
			nextNumber = sts.getCorrespondingNumber(nextNumber)
			break
		}
	}

	for _, sts := range waterToLight {
		if sts.contains(nextNumber) {
			nextNumber = sts.getCorrespondingNumber(nextNumber)
			break
		}
	}

	for _, sts := range lightToTemperature {
		if sts.contains(nextNumber) {
			nextNumber = sts.getCorrespondingNumber(nextNumber)
			break
		}
	}

	for _, sts := range temperatureToHumidity {
		if sts.contains(nextNumber) {
			nextNumber = sts.getCorrespondingNumber(nextNumber)
			break
		}
	}

	for _, sts := range humidityToLocation {
		if sts.contains(nextNumber) {
			nextNumber = sts.getCorrespondingNumber(nextNumber)
			break
		}
	}

	return nextNumber
}
