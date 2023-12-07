package main

import (
	"log"
    "os"
    "strconv"
	"regexp"
    "strings"
)

func main() {
	filename := "input.txt"

    partOneResult := partOne(filename)
    partTwoResult := partTwo(filename)

    log.Printf("Part One: %v", partOneResult)
    log.Printf("Part Two: %v", partTwoResult)
}

func indexAt(s, sep string, n int) int {
    idx := strings.Index(s[n:], sep)
    if idx > -1 {
        idx += n
    }
    return idx
}

func partOne(filename string) int {
    file, err := os.ReadFile(filename)

    if err != nil {
        log.Fatal("Couldn't read file")
    }

	engineString := string(file)

	engineMatrix := strings.Split(engineString, "\n")
	var engineParts []int

	re := regexp.MustCompile("[0-9]+")
	for engineMatrixRowIndex, engineMatrixRow := range engineMatrix {
		numbers := re.FindAllString(engineMatrixRow, -1)

		lastIndex := 0
		for _, number := range numbers {
			numberColStartIndex := indexAt(engineMatrixRow, number, lastIndex)
			numberColEndIndex := numberColStartIndex + len(number) - 1
			lastIndex = numberColEndIndex

			isEnginePart := false

			var row int
			var rowEnd int

			if engineMatrixRowIndex > 0 {
				row = engineMatrixRowIndex - 1
			} else {
				row = engineMatrixRowIndex
			}

			if engineMatrixRowIndex < len(engineMatrix) - 1 {
				rowEnd = engineMatrixRowIndex + 1
			} else {
				rowEnd = engineMatrixRowIndex
			}

			for ; row <= rowEnd; row++ {

				var col int
				var colEnd int

				if numberColStartIndex > 0 {
					col = numberColStartIndex - 1
				} else {
					col = numberColStartIndex
				}

				if numberColEndIndex < len(engineMatrix[row]) - 1 {
					colEnd = numberColEndIndex + 1
				} else {
					colEnd = numberColEndIndex
				}

				for ; col <= colEnd; col++ {
					char := string(engineMatrix[row][col])
					_, err := strconv.Atoi(char)
					if char != "." && err != nil {
						isEnginePart = true
						break
					}
				}
				if isEnginePart {
					break
				}
			}

			if isEnginePart {
				partNumber, err := strconv.Atoi(number)
				if err == nil {
					engineParts = append(engineParts, partNumber)
				}
			}
			
		}

	}

	var sum int
	for _, part := range engineParts {
		sum += part
	}
   
    return sum
}

type EnginePart struct {
	number int
	rowIndex int
	colStartIndex int
	colEndIndex int
}

type Gear struct {
	rowIndex int
	colIndex int
}

func partTwo(filename string) int {
    file, err := os.ReadFile(filename)

    if err != nil {
        log.Fatal("Couldn't read file")
    }

	engineString := string(file)

	engineMatrix := strings.Split(engineString, "\n")

	var engineParts []EnginePart

	re := regexp.MustCompile("[0-9]+")
	for engineMatrixRowIndex, engineMatrixRow := range engineMatrix {
		numbers := re.FindAllString(engineMatrixRow, -1)

		lastIndex := 0
		for _, number := range numbers {
			numberColStartIndex := indexAt(engineMatrixRow, number, lastIndex)
			numberColEndIndex := numberColStartIndex + len(number) - 1
			lastIndex = numberColEndIndex

			isEnginePart := false

			var row int
			var rowEnd int

			if engineMatrixRowIndex > 0 {
				row = engineMatrixRowIndex - 1
			} else {
				row = engineMatrixRowIndex
			}

			if engineMatrixRowIndex < len(engineMatrix) - 1 {
				rowEnd = engineMatrixRowIndex + 1
			} else {
				rowEnd = engineMatrixRowIndex
			}

			for ; row <= rowEnd; row++ {

				var col int
				var colEnd int

				if numberColStartIndex > 0 {
					col = numberColStartIndex - 1
				} else {
					col = numberColStartIndex
				}

				if numberColEndIndex < len(engineMatrix[row]) - 1 {
					colEnd = numberColEndIndex + 1
				} else {
					colEnd = numberColEndIndex
				}

				for ; col <= colEnd; col++ {
					char := string(engineMatrix[row][col])
					_, err := strconv.Atoi(char)
					if char != "." && err != nil {
						isEnginePart = true
						break
					}
				}
				if isEnginePart {
					break
				}
			}

			if isEnginePart {
				partNumber, err := strconv.Atoi(number)
				enginePart := EnginePart {
					number: partNumber,
					rowIndex: engineMatrixRowIndex,
					colStartIndex: numberColStartIndex,
					colEndIndex: numberColEndIndex,
				}
				if err == nil {
					engineParts = append(engineParts, enginePart)
				}
			}
			
		}

	}

	var gearRatios []int
	for engineMatrixRowIndex, engineMatrixRow := range engineMatrix {
		for i := 0; i < len(engineMatrixRow); i++ {
			if (string(engineMatrixRow[i]) == "*") {
				potentialGear := Gear {
					rowIndex: engineMatrixRowIndex,
					colIndex: i,
				}
				gearAdjascentParts := []int{}
				for _, enginePart := range engineParts {
					if enginePart.rowIndex >= potentialGear.rowIndex - 1 && enginePart.rowIndex <= potentialGear.rowIndex + 1 {
						if (enginePart.colEndIndex >= potentialGear.colIndex - 1 && enginePart.colEndIndex <= potentialGear.colIndex + 1) ||
						(enginePart.colStartIndex >= potentialGear.colIndex - 1 && enginePart.colStartIndex <= potentialGear.colIndex + 1) ||
						(potentialGear.colIndex >= enginePart.colStartIndex && potentialGear.colIndex <= enginePart.colEndIndex) {
							gearAdjascentParts = append(gearAdjascentParts, enginePart.number)
						}
					}
				}
				if (len(gearAdjascentParts) == 2) {
					gearRatios = append(gearRatios, gearAdjascentParts[0] * gearAdjascentParts[1])
				}
			}
		}
	}

	var sum int
	for _, gearRatio := range gearRatios {
		sum += gearRatio
	}
   
    return sum
}