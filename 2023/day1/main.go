package main

import (
    "log"
    "bufio"
    "os"
    "strconv"
    "strings"
)

type numberInWord struct {
    digit int;
    word string;
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

    var sum int
    var allFoundNumbers []int

    for scanner.Scan() {
        line := scanner.Text()

        var firstNumber int
        var lastNumber int
        firstNumberFound := false

        for i := 0; i < len(line); i++ {
            num, err := strconv.Atoi(string(line[i]))
            if err != nil {
                continue
            }

            if !firstNumberFound {
                firstNumber = num
                firstNumberFound = true
            }

            lastNumber = num
        }

        concactedNumber, _ := strconv.Atoi(strconv.Itoa(firstNumber) + strconv.Itoa(lastNumber))
        allFoundNumbers = append(allFoundNumbers, concactedNumber)
        sum += concactedNumber

    }

    return sum
}

func partTwo(filename string) int {
    numberInWords := []numberInWord {
        {digit: 0, word: "zero"},
        {digit: 1, word: "one"},
        {digit: 2, word: "two"},
        {digit: 3, word: "three"},
        {digit: 4, word: "four"},
        {digit: 5, word: "five"},
        {digit: 6, word: "six"},
        {digit: 7, word: "seven"},
        {digit: 8, word: "eight"},
        {digit: 9, word: "nine"},
    }

    file, err := os.Open(filename)

    if err != nil {
        log.Fatal("Couldn't read file")
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    var sum int
    var allFoundNumbers []int
    lineNo := 0

    for scanner.Scan() {
        lineNo++

        line := scanner.Text()

        firstNumberInWordIndex := -1
        lastNumberInWordIndex := -1
        var firstNumberInWord int
        var lastNumberInWord int

        for j := 0; j < len(numberInWords); j++ {
            firstOccurenceIndex := strings.Index(line, numberInWords[j].word)
            lastOccurenceIndex := strings.LastIndex(line, numberInWords[j].word)

            if firstOccurenceIndex != -1 && firstOccurenceIndex < firstNumberInWordIndex || firstNumberInWordIndex == -1 {
                firstNumberInWordIndex = firstOccurenceIndex
                firstNumberInWord = numberInWords[j].digit
            }

            if lastOccurenceIndex > lastNumberInWordIndex {
                lastNumberInWordIndex = lastOccurenceIndex
                lastNumberInWord = numberInWords[j].digit
            }
        }

        firstNumberIndex := -1
        lastNumberIndex := -1
        var firstNumber int
        var lastNumber int
        firstNumberFound := false

        for i := 0; i < len(line); i++ {
            num, err := strconv.Atoi(string(line[i]))
            if err != nil {
                continue
            }

            if !firstNumberFound {
                firstNumber = num
                firstNumberIndex = i
                firstNumberFound = true
            }

            lastNumberIndex = i
            lastNumber = num
        }

        if firstNumberInWordIndex > -1 && firstNumberIndex > firstNumberInWordIndex || firstNumberIndex == -1 {
            firstNumber = firstNumberInWord
        }

        if lastNumberInWordIndex > -1 && lastNumberIndex < lastNumberInWordIndex || lastNumberIndex == -1 {
            lastNumber = lastNumberInWord
        }

        concactedNumber, _ := strconv.Atoi(strconv.Itoa(firstNumber) + strconv.Itoa(lastNumber))

        allFoundNumbers = append(allFoundNumbers, concactedNumber)
        sum += concactedNumber

    }

    return sum
}