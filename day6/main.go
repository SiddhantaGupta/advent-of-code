package main

import (
    "log"
    "bufio"
    "os"
    "strconv"
    "strings"
)
/*
distance = speed * (totalTimes - speed)
*/
func main() {
    filename := "input.txt"

    partOneResult := partOne(filename)
    log.Printf("Part One: %v", partOneResult)

    partTwoResult := partTwo(filename)
    log.Printf("Part Two: %v", partTwoResult)
}

func partOne(filename string) int {
	timeDistanceMap := getTimeDistanceData(filename)

	allWinTimes := [][]int{}
	for raceIndex, time := range timeDistanceMap["Time"] {
		winTimes := getWinTimes(time, timeDistanceMap["Distance"][raceIndex])
		allWinTimes = append(allWinTimes, winTimes)
	}

	allMarginOfErrors := 1
	for _, winTimes := range allWinTimes {
		allMarginOfErrors *= len(winTimes)
	}

	return allMarginOfErrors
}

func getTimeDistanceData(filename string) map[string][]int {
	file, err := os.Open(filename)

    if err != nil {
        log.Fatal("Couldn't read file")
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

	timeDistanceMap := make(map[string][]int)

    for scanner.Scan() {
        line := scanner.Text()

		data := strings.Split(line, ":")
		dataType := data[0]
		dataStrList := strings.Split(data[1], " ")

		for _, dataStrItem := range dataStrList {
			if dataStrItem == "" || dataStrItem == " " {
				continue;
			}
			dataInt, _ := strconv.Atoi(strings.TrimSpace(dataStrItem))
			timeDistanceMap[dataType] = append(timeDistanceMap[dataType], dataInt)
		} 
	}

	return timeDistanceMap
}

func getWinTimes(time int, distance int) []int {
	winTimes := []int{}
	for speed := 1; speed < time; speed++ {
		distanceTravelled := speed * (time - speed)
		if distanceTravelled > distance {
			winTimes = append(winTimes, speed)
		}
	}
	return winTimes
}

func partTwo(filename string) int {
	timeDistanceMap := getTimeDistanceData(filename)
	time, _ := join(timeDistanceMap["Time"])
	distance, _ := join(timeDistanceMap["Distance"])

	marginOfError := 0
	for speed := 1; speed < time; speed++ {
		distanceTravelled := speed * (time - speed)
		if distanceTravelled > distance {
			marginOfError++
		}
	}

	return marginOfError
}

func join(nums []int) (int, error) {
    var str string
    for i := range nums {
        str += strconv.Itoa(nums[i])
    }
    num, err := strconv.Atoi(str)
    if err != nil {
        return 0, err
    } else {
        return num, nil
    }
}