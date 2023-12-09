package main;

import (
    "log"
    "bufio"
    "os"
	"regexp"
    "strconv"
)

func main() {
	filename := "input.txt"

    partOneResult := partOne(filename)
    log.Printf("Part One: %v", partOneResult)

    partTwoResult := partTwo(filename)
    log.Printf("Part Two: %v", partTwoResult)
}

func partOne(filename string) int {
	oasisAndSandInstabilityData := parseOasisAndSandInstability(filename)

	nextValuePredictionList := []int{}
	for _, oasisAndSandInstabilityHistory := range oasisAndSandInstabilityData {
		historyPredictionSequenceTable := getHistoryPredictionSequenceTable(oasisAndSandInstabilityHistory)
		nextValuePredictionList = append(nextValuePredictionList, getNextValuePredictionFromHistorySequence(historyPredictionSequenceTable))
	}

	nextValuePredictionSum := 0
	for _, nextValuePrediction := range nextValuePredictionList {
		nextValuePredictionSum += nextValuePrediction
	}
	return nextValuePredictionSum
}

func partTwo(filename string) int {
	oasisAndSandInstabilityData := parseOasisAndSandInstability(filename)

	prevValuePredictionList := []int{}
	for _, oasisAndSandInstabilityHistory := range oasisAndSandInstabilityData {
		historyPredictionSequenceTable := getHistoryPredictionSequenceTable(oasisAndSandInstabilityHistory)
		prevValuePredictionList = append(prevValuePredictionList, getPrevValuePredictionFromHistorySequence(historyPredictionSequenceTable))
	}

	prevValuePredictionSum := 0
	for _, prevValuePrediction := range prevValuePredictionList {
		prevValuePredictionSum += prevValuePrediction
	}
	return prevValuePredictionSum
}

func parseOasisAndSandInstability(oasisAndSandInstabilityDataFilename string) [][]int {
	file, err := os.Open(oasisAndSandInstabilityDataFilename)
    if err != nil {
        log.Fatal("Couldn't read file")
    }
    defer file.Close()

	oasisAndSandInstabilityData := [][]int{}

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		re := regexp.MustCompile("-?[0-9]+")
		valueHistoryStrList := re.FindAllString(line, -1)

		valueHistoryList := []int{}
		for _, history := range valueHistoryStrList {
			historyInt, err := strconv.Atoi(history)
			if err != nil {continue}
			valueHistoryList = append(valueHistoryList, historyInt)
		}
		oasisAndSandInstabilityData = append(oasisAndSandInstabilityData, valueHistoryList)
	}

	return oasisAndSandInstabilityData
}

func getNextValuePredictionFromHistorySequence(historyPredictionSequenceTable [][]int) int {
	historyPredictionSequenceTable[len(historyPredictionSequenceTable) - 1] = append(historyPredictionSequenceTable[len(historyPredictionSequenceTable) - 1], 0)

	for index := len(historyPredictionSequenceTable) - 1; index >= 1; index-- {
		nextValue := historyPredictionSequenceTable[index - 1][len(historyPredictionSequenceTable[index - 1]) - 1] + historyPredictionSequenceTable[index][len(historyPredictionSequenceTable[index]) - 1]
		historyPredictionSequenceTable[index - 1] = append(historyPredictionSequenceTable[index - 1], nextValue)
	}

	return historyPredictionSequenceTable[0][len(historyPredictionSequenceTable[0]) - 1]
}

func getPrevValuePredictionFromHistorySequence(historyPredictionSequenceTable [][]int) int {
	historyPredictionSequenceTable[len(historyPredictionSequenceTable) - 1] = append([]int{0}, historyPredictionSequenceTable[len(historyPredictionSequenceTable) - 1]...)

	for index := len(historyPredictionSequenceTable) - 1; index >= 1; index-- {
		prevValue := historyPredictionSequenceTable[index - 1][0] - historyPredictionSequenceTable[index][0]
		historyPredictionSequenceTable[index - 1] = append([]int{prevValue}, historyPredictionSequenceTable[index - 1]...)
	}

	return historyPredictionSequenceTable[0][0]
}

func getHistoryPredictionSequenceTable(history []int) [][]int {
	historyPredictionSequenceTable := [][]int{history}
	
	for historyIndex := 0; ; historyIndex++ {
		history := historyPredictionSequenceTable[historyIndex]
		historyPredictionSequenceTable = append(historyPredictionSequenceTable, []int{})
		newDiffSequenceIndex := historyIndex + 1
		for valueIndex := 0; valueIndex < len(history) - 1; valueIndex++ {
			difference := history[valueIndex + 1] - history[valueIndex]
			historyPredictionSequenceTable[newDiffSequenceIndex] = append(historyPredictionSequenceTable[newDiffSequenceIndex], difference)
		}

		isLastSequence := true
		for _, value := range historyPredictionSequenceTable[newDiffSequenceIndex] {
			if (value != 0) {
				isLastSequence = false
			}
		}

		if isLastSequence {
			break
		}
	}
	
	return historyPredictionSequenceTable
}