package main

import (
    "log"
    "bufio"
    "os"
    "strconv"
    "strings"
	"slices"
	"regexp"
)

type ScratchCard struct {
	id int
	winningNumbers []int
	cardNumbers []int
}

func main() {
    filename := "input.txt"

    partOneResult := partOne(filename)
    partTwoResult := partTwo(filename)

    log.Printf("Part One: %v", partOneResult)
    log.Printf("Part Two: %v", partTwoResult)
}

func partOne(filename string) int {
    file, err := os.Open(filename)

    if err != nil {
        log.Fatal("Couldn't read file")
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
	var cardPoints []int
    for scanner.Scan() {
        cardStr := scanner.Text()

		winningNumbersStr := strings.TrimSpace(strings.Split(strings.Split(cardStr, ":")[1], "|")[0])
		winningNumbersStrList := strings.Split(winningNumbersStr, " ")
		var winningNumbers []int
		for _, winningNumberStr := range winningNumbersStrList {
			winningNumber, _ := strconv.Atoi(strings.TrimSpace(winningNumberStr))
			winningNumbers = append(winningNumbers, winningNumber)
		}

		cardNumbersStr := strings.TrimSpace(strings.Split(strings.Split(cardStr, ":")[1], "|")[1])
		cardNumbersStrList := strings.Split(cardNumbersStr, " ")
		var cardNumbers []int
		for _, cardNumberStr := range cardNumbersStrList {
			cardNumber, err := strconv.Atoi(strings.TrimSpace(cardNumberStr))
			if err != nil {continue}
			cardNumbers = append(cardNumbers, cardNumber)
		}

		points := 0
		for _, cardNumber := range cardNumbers {
			if slices.Contains(winningNumbers, cardNumber) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		cardPoints = append(cardPoints, points)

	}

	sum := 0
	for _, cardPoint := range cardPoints {
		sum += cardPoint
	}

	return sum
}

func partTwo(filename string) int {
	file, err := os.ReadFile(filename)

    if err != nil {
        log.Fatal("Couldn't read file")
    }

	fileStr := string(file)
	cardStrs := strings.Split(fileStr, "\n")

	re := regexp.MustCompile("[0-9]+")

	scratchCards := make(map[int]ScratchCard)

    for _, cardStr := range cardStrs {

		cardNumberStr := re.FindAllString(strings.Split(cardStr, ":")[0], -1)[0]
		cardNumber, _ := strconv.Atoi(cardNumberStr)

		winningNumbersStr := strings.TrimSpace(strings.Split(strings.Split(cardStr, ":")[1], "|")[0])
		winningNumbersStrList := strings.Split(winningNumbersStr, " ")
		var winningNumbers []int
		for _, winningNumberStr := range winningNumbersStrList {
			winningNumber, _ := strconv.Atoi(strings.TrimSpace(winningNumberStr))
			winningNumbers = append(winningNumbers, winningNumber)
		}

		cardNumbersStr := strings.TrimSpace(strings.Split(strings.Split(cardStr, ":")[1], "|")[1])
		cardNumbersStrList := strings.Split(cardNumbersStr, " ")
		var cardNumbers []int
		for _, cardNumberStr := range cardNumbersStrList {
			cardNumber, err := strconv.Atoi(strings.TrimSpace(cardNumberStr))
			if err != nil {continue}
			cardNumbers = append(cardNumbers, cardNumber)
		}

		scratchCard := ScratchCard {
			id: cardNumber,
			winningNumbers: winningNumbers,
			cardNumbers: cardNumbers,
		}

		scratchCards[scratchCard.id] = scratchCard
	}
	scratchCardCounter := make(map[int]int)

	for i := 1; i <= len(scratchCards); i++ {
		scratchCardCounter[i] = 1
	}

	for i := 1; i <= len(scratchCards); i++ {
		winningNumberCounter := 0
		for _, currentScratchCardCardNumber := range scratchCards[i].cardNumbers {
			if slices.Contains(scratchCards[i].winningNumbers, currentScratchCardCardNumber) {
				winningNumberCounter++
			}
		}

		for j := 1; j <= winningNumberCounter; j++ {
			scratchCardCounter[i+j] += scratchCardCounter[i]
		}
	}

	sum := 0
	for _, val := range scratchCardCounter {
		sum += val
	}
	return sum
}