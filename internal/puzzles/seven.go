package puzzles

import (
	"sort"
	"strconv"
	"strings"
)

var cardScores = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"*": 1,
}

type hand struct {
	cards string
	bid   int
	score int64
}

// Just look away
func SolveSevenOne(input []string) int {
	hands := []hand{}
	for _, line := range input {
		var splitLine = strings.Split(line, " ")
		bid, _ := strconv.Atoi(splitLine[1])
		score := getHandScore(splitLine[0])
		hand := hand{cards: splitLine[0], bid: bid, score: score}
		hands = append(hands, hand)
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].score < hands[j].score
	})

	sum := 0
	for rank, hand := range hands {
		sum += ((rank + 1) * hand.bid)
	}

	return sum
}

// Why are you still looking??
func SolveSevenTwo(input []string) int {
	hands := []hand{}
	for _, line := range input {
		// J is now worth less, so in order to keep the 7-1 working I'm replacing the letter J with * here (* is worth 1)
		line = strings.Replace(line, "J", "*", -1)
		var splitLine = strings.Split(line, " ")
		bid, _ := strconv.Atoi(splitLine[1])
		score := getHandScoreFromWish(splitLine[0])
		hand := hand{cards: splitLine[0], bid: bid, score: score}
		hands = append(hands, hand)
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].score < hands[j].score
	})

	sum := 0
	for rank, hand := range hands {
		sum += ((rank + 1) * hand.bid)
	}

	return sum
}

// I wish I picked C# for this years' AoC. I can't figure out how to do this in a normal way without using 5 billion lines of code to do something this simple
func getHandScore(cards string) int64 {
	score := int64(0)
	secondaryScore := int64(0)
	cardsPerType := make(map[string]int)
	for _, card := range cards {
		cardsPerType[string(card)] += 1
		secondaryScore *= 100
		secondaryScore += int64(cardScores[string(card)])
	}

	if len(cardsPerType) == 5 {
		// high card, lowest score
		score = 10000000000
	} else if len(cardsPerType) == 1 {
		// five of a kind, max score
		score = 70000000000
	} else {
		highest := 0
		secondHighest := 0
		for _, numberOfType := range cardsPerType {
			if numberOfType > highest {
				secondHighest = highest
				highest = numberOfType
			} else if numberOfType > secondHighest {
				secondHighest = numberOfType
			}
		}
		if highest == 4 {
			// four of a kind, second highest score
			score = 60000000000
		} else if highest == 3 && secondHighest == 2 {
			// full house, third highest score
			score = 50000000000
		} else if highest == 3 {
			// three of a kind, fourth highest score
			score = 40000000000
		} else if highest == 2 && secondHighest == 2 {
			// two pair, third lowest score
			score = 30000000000
		} else if highest == 2 {
			// one pair, second lowest score
			score = 20000000000
		}
	}

	return score + secondaryScore
}

// ...
func getHandScoreFromWish(cards string) int64 {
	score := int64(0)
	secondaryScore := int64(0)
	cardsPerType := make(map[string]int)
	hasJoker := false
	for _, card := range cards {
		if string(card) == "*" {
			hasJoker = true
		}
		cardsPerType[string(card)] += 1
		secondaryScore *= 100
		secondaryScore += int64(cardScores[string(card)])
	}
	if len(cardsPerType) == 5 && !hasJoker {
		// high card, lowest score
		score = 10000000000
	} else if len(cardsPerType) == 1 || (len(cardsPerType) == 2 && hasJoker) {
		// five of a kind, max score
		score = 70000000000
	} else {
		highest := 0
		secondHighest := 0
		jokerAmount := 0
		for card, numberOfType := range cardsPerType {
			if card == "*" {
				jokerAmount = numberOfType
			} else if numberOfType > highest {
				secondHighest = highest
				highest = numberOfType
			} else if numberOfType > secondHighest {
				secondHighest = numberOfType
			}
		}
		highest += jokerAmount
		if highest == 4 {
			// four of a kind, second highest score
			score = 60000000000
		} else if highest == 3 && secondHighest == 2 {
			// full house, third highest score
			score = 50000000000
		} else if highest == 3 {
			// three of a kind, fourth highest score
			score = 40000000000
		} else if highest == 2 && secondHighest == 2 {
			// two pair, third lowest score
			score = 30000000000
		} else if highest == 2 {
			// one pair, second lowest score
			score = 20000000000
		}
	}

	return score + secondaryScore
}
