package main;

import (
    "log"
    "bufio"
    "os"
    "strconv"
    "strings"
	"reflect"
	"sort"
)

type HandData struct {
	cards []string
	bet int
}

func main() {
	filename := "input.txt"

    partOneResult := partOne(filename)
    log.Printf("Part One: %v", partOneResult)

    partTwoResult := partTwo(filename)
    log.Printf("Part Two: %v", partTwoResult)
}

func partOne(filename string) int {
	cardsOrder := map[string]int {
		"A": 12,
		"K": 11,
		"Q": 10,
		"J": 9,
		"T": 8,
		"9": 7,
		"8": 6,
		"7": 5,
		"6": 4,
		"5": 3,
		"4": 2,
		"3": 1,
		"2": 0,
	}

	handTypes := []map[int]int {
		// cardCounter: appearance
		{ // five of a kind
			5: 1,
		},
		{ // four of a kind
			4: 1,
			1: 1,
		},
		{ // full house
			3: 1,
			2: 1,
		},
		{ // three of a kind
			3: 1,
			1: 2,
		},
		{ // two pair
			2: 2,
			1: 1,
		},
		{ // one pair
			2: 1,
			1: 3,
		},
		{ // high card
			1: 5,
		},
	}

	file, err := os.Open(filename)

    if err != nil {
        log.Fatal("Couldn't read file")
    }

    defer file.Close()

	handsDataByType := make([][]HandData, len(handTypes))

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
		handData := parseHandData(line)
		handTypeIndex := getHandType(handTypes, handData, "")
		handsDataByType[handTypeIndex] = append(handsDataByType[handTypeIndex], handData)
	}

	for handTypeIndex, handDataList := range handsDataByType {
		handsDataByType[handTypeIndex] = sortHandsByCardSequence(cardsOrder, handDataList)
	}

	totalRanks := 0
	for _, handsData := range handsDataByType {
		totalRanks += len(handsData)
	}

	totalWiningBet := 0
	for _, handsData := range handsDataByType {
		for _, hand := range handsData {
			totalWiningBet += hand.bet * totalRanks
			totalRanks--
		}
	}

	return totalWiningBet
}

func partTwo(filename string) int {
	cardsOrder := map[string]int {
		"A": 12,
		"K": 11,
		"Q": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
		"J": 0,
	}

	handTypes := []map[int]int {
		// cardCounter: appearance
		{ // five of a kind
			5: 1,
		},
		{ // four of a kind
			4: 1,
			1: 1,
		},
		{ // full house
			3: 1,
			2: 1,
		},
		{ // three of a kind
			3: 1,
			1: 2,
		},
		{ // two pair
			2: 2,
			1: 1,
		},
		{ // one pair
			2: 1,
			1: 3,
		},
		{ // high card
			1: 5,
		},
	}

	file, err := os.Open(filename)

    if err != nil {
        log.Fatal("Couldn't read file")
    }

    defer file.Close()

	handsDataByType := make([][]HandData, len(handTypes))

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
		handData := parseHandData(line)
		handTypeIndex := getHandType(handTypes, handData, "J")
		handsDataByType[handTypeIndex] = append(handsDataByType[handTypeIndex], handData)
	}

	for handTypeIndex, handDataList := range handsDataByType {
		handsDataByType[handTypeIndex] = sortHandsByCardSequence(cardsOrder, handDataList)
	}

	totalRanks := 0
	for _, handsData := range handsDataByType {
		totalRanks += len(handsData)
	}

	totalWiningBet := 0
	for _, handsData := range handsDataByType {
		for _, hand := range handsData {
			totalWiningBet += hand.bet * totalRanks
			totalRanks--
		}
	}

	return totalWiningBet
}


func parseHandData(handStr string) HandData {

	handDataList := strings.Split(handStr, " ")
	cards := strings.Split(handDataList[0], "")
	bet, _ := strconv.Atoi(handDataList[1])

	handData := HandData {
		cards: cards,
		bet: bet,
	}

	return handData
}

func sortHandsByCardSequence(cardsOrder map[string]int, handsData []HandData) []HandData {
	sort.SliceStable(handsData, func(i, j int) bool {
		for k := 0; k < 5; k++ {
			firstCardWeight := cardsOrder[handsData[i].cards[k]]
			secondCardWeight := cardsOrder[handsData[j].cards[k]]
			if firstCardWeight > secondCardWeight {
				return true
			} else if firstCardWeight < secondCardWeight {
				return false
			}
		}

		return true
	})

	return handsData
}

func getHandType(handTypeConfig []map[int]int, hand HandData, wildCard string) int {
	cardsCount := make(map[string]int)
	for _, card := range hand.cards {
		_, ok := cardsCount[card]
		if !ok {
			cardsCount[card] = 1
			continue
		}
		cardsCount[card]++
	}

	handType := make(map[int]int)
	for card, count := range cardsCount {
		if card == wildCard {
			continue
		}
		_, ok := handType[count]
		if !ok {
			handType[count] = 1
			continue
		}
		handType[count]++
	}

	wildCardCount, wildCardExists := cardsCount[wildCard]
	if wildCardExists {
		largest := 0
		for htype, _ := range handType {
			if largest < htype {
				largest = htype
			}
		}
		// we take one set from our largest combination
		handType[largest + wildCardCount] = 1
		handType[largest] -= 1
		if handType[largest] <= 0 {
			delete(handType, largest)
		}
	}

	handTypeIndex := 0
	for typeIndex, _ := range handTypeConfig {
		if reflect.DeepEqual(handTypeConfig[typeIndex], handType) {
			handTypeIndex = typeIndex
			break
		}
	}

	return handTypeIndex
}
