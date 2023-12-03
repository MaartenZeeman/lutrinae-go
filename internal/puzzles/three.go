package puzzles

const ASCII_FIRST_DIGIT = 48
const ASCII_LAST_DIGIT = 57
const ASCII_PERIOD = 46
const ASCII_ASTERIX = 42

type partNumber struct {
	number     int
	startIndex int
	endIndex   int
	lineIndex  int
}

type gear struct {
	lineIndex int
	charIndex int
}

// Checks whether any of the given indices borders this partnumber
func (partNumber *partNumber) bordersSymbol(symbolIncices []int) bool {
	for _, symbolIndex := range symbolIncices {
		if symbolIndex >= (partNumber.startIndex-1) && symbolIndex <= (partNumber.endIndex+1) {
			return true
		}
	}

	return false
}

// This code isn't great. Somehow i thought it was a good idea to set myself some rules and challenges before knowing what the assigment would be
func SolveThreeOne(input []string) int {
	var sum = 0

	// Slices of sadness
	previousLinePartNumbers := []partNumber{}
	previousSymbolIndices := []int{}

	for _, line := range input {
		// the 'fun' begins. After using regex on day 1 and splits and substrings on day 2, I didn't want to use either now. Big mistake
		// first we need some slices to store the partnumbers and symbols we've seen
		currentLinePartNumbers := []partNumber{}
		currentSymbolIndices := []int{}
		partialPartNumber := partNumber{}

		// Loop through all the characters in the current line
		for charIndex, character := range line {
			// If the character is a dot/period and the one(s) before digit(s), then we need to complete the partnumber and add it to the list
			if character == ASCII_PERIOD && partialPartNumber.number > 0 {
				// check if there are any symbols in the previous row (the one above) or in this one that border the part number.
				if partialPartNumber.bordersSymbol(previousSymbolIndices) || partialPartNumber.bordersSymbol(currentSymbolIndices) {
					sum += partialPartNumber.number
				} else {
					// if that's not the case we need to add it to the list of part numbers so we can check it when we come across symbols on the next row
					currentLinePartNumbers = append(currentLinePartNumbers, partialPartNumber)
				}
				partialPartNumber = partNumber{}
			} else if character >= ASCII_FIRST_DIGIT && character <= ASCII_LAST_DIGIT && partialPartNumber.number > 0 {
				// the previous char was a digit and so is this one, so we need to add it to the part number.
				// Multiply the current one by 10 and add it (so 3 and 5 will be 35 rather than 8)
				partialPartNumber.number = partialPartNumber.number*10 + int(character-'0')
				partialPartNumber.endIndex = charIndex
			} else if character >= ASCII_FIRST_DIGIT && character <= ASCII_LAST_DIGIT {
				// previous character wasn't a digit and this one is, so we need to set the initial number and indices
				partialPartNumber.number = int(character - '0')
				partialPartNumber.startIndex = charIndex
				partialPartNumber.endIndex = charIndex
			} else if character != ASCII_PERIOD && partialPartNumber.number > 0 {
				// symbol that follows a number. Add the part number to the sum and check if the symbol also borders a number in the previous row
				sum += partialPartNumber.number
				sum += bordersNumber(&previousLinePartNumbers, charIndex)
				partialPartNumber = partNumber{}
				// add symbol to the list of symbols so we can check if it borders numbers in the next line as well (forgot this one initially)
				currentSymbolIndices = append(currentSymbolIndices, charIndex)
			} else if character != ASCII_PERIOD {
				// symbol but it doesn't follow a number. Just check if it borders any in the previous line and add to the list
				sum += bordersNumber(&previousLinePartNumbers, charIndex)
				currentSymbolIndices = append(currentSymbolIndices, charIndex)
			}
		}
		// if the line ends with a part number, we need to check if it borders anything and if it does add it to the sum
		if partialPartNumber.number > 0 {
			if partialPartNumber.bordersSymbol(previousSymbolIndices) || partialPartNumber.bordersSymbol(currentSymbolIndices) {
				sum += partialPartNumber.number
			} else {
				// doesn't border anything yet, so add to list and check later
				currentLinePartNumbers = append(currentLinePartNumbers, partialPartNumber)
			}
		}
		// overwrite the slices of previous values because we're going to the next line!
		previousLinePartNumbers = currentLinePartNumbers
		previousSymbolIndices = currentSymbolIndices
	}

	return sum
}

// like SolveThreeOne but without the sums and with the * (Gear) check
func SolveThreeTwo(input []string) int {
	partNumbers := []partNumber{}
	gears := []gear{}

	for lineIndex, line := range input {
		partialPartNumber := partNumber{}
		for charIndex, character := range line {
			if character >= ASCII_FIRST_DIGIT && character <= ASCII_LAST_DIGIT && partialPartNumber.number > 0 {
				partialPartNumber.number = partialPartNumber.number*10 + int(character-'0')
				partialPartNumber.endIndex = charIndex
				continue
			} else if character >= ASCII_FIRST_DIGIT && character <= ASCII_LAST_DIGIT {
				partialPartNumber.number = int(character - '0')
				partialPartNumber.startIndex = charIndex
				partialPartNumber.endIndex = charIndex
				// this is new. we need to keep track of the line now because we need to loop through this mess later
				partialPartNumber.lineIndex = lineIndex
				continue
			} else if partialPartNumber.number > 0 {
				// we only need to keep track of the gears, but symbols are still
				partNumbers = append(partNumbers, partialPartNumber)
				partialPartNumber = partNumber{}
			}
			// if it's an asterix we need to add it to the gear slice
			if character == ASCII_ASTERIX {
				gears = append(gears, gear{lineIndex, charIndex})
			}
		}
		if partialPartNumber.number > 0 {
			partNumbers = append(partNumbers, partialPartNumber)
		}
	}

	return getGearRatios(partNumbers, gears)
}

// checks if the given (character) index borders a partnumber
// using structs this way hurt me and the performance here, but the input is so small that the entire assiment still only takes 2ms anyway
func bordersNumber(partNumbers *[]partNumber, index int) int {
	var sum = 0
	// we need to get rid of a part number once we've added it to the sum,
	// otherwise it would be added multiple times if multiple symbols border it
	unusedPartNumbers := []partNumber{}
	for _, partNumber := range *partNumbers {
		// index -1 and +1 here because it can also border the symbol diagonally.
		if index >= (partNumber.startIndex-1) && index <= (partNumber.endIndex+1) {
			sum += partNumber.number
		} else {
			// if it doesn't border the symbol we should add it to the new slice so we can try it again with the next symbol (if any)
			unusedPartNumbers = append(unusedPartNumbers, partNumber)
		}
	}
	// overwrite the list of part numbers with the new list that doesn't include part numbers that we just used (added)
	*partNumbers = unusedPartNumbers
	return sum
}

// Loops through the gears and part numbers (really inefficiently) and returns the sum of gear ratios
// It loops through *all* the part numbers for every gear. I should have used a map or something to filter the lines
func getGearRatios(partNumbers []partNumber, gears []gear) int {
	var sum = 0

	for _, gear := range gears {
		gearSum, gearAmount := 1, 0
		for _, partNumber := range partNumbers {
			// break if we've past the relevant lines (gear line + 1) or we've seen too many gears already
			// continue if we aren't on the first relevant line (gear line -1) yet
			if partNumber.lineIndex > gear.lineIndex+1 || gearAmount > 2 {
				break
			} else if partNumber.lineIndex < gear.lineIndex-1 {
				continue
			} else if gear.charIndex >= (partNumber.startIndex-1) && gear.charIndex <= (partNumber.endIndex+1) {
				gearSum *= partNumber.number
				gearAmount++
			}
		}
		if gearAmount == 2 {
			sum += gearSum
		}
	}

	return sum
}
