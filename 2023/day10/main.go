package main

import (
	"errors"
	"log"
	"math"
	"os"
	"reflect"
	"strings"
)

var pipeDirections map[string][]string
var directions map[string][]int

func main() {

	directions = map[string][]int{
		"east":  {0, 1},
		"west":  {0, -1},
		"north": {-1, 0},
		"south": {1, 0},
	}

	pipeDirections = map[string][]string{
		"|": {"north", "south"},
		"-": {"east", "west"},
		"L": {"north", "east"},
		"J": {"north", "west"},
		"7": {"south", "west"},
		"F": {"south", "east"},
	}

	filename := "input.txt"

	partOneResult := partOne(filename)
	log.Printf("Part One: %v", partOneResult)

	partTwoResult := partTwo(filename)
	log.Printf("Part Two: %v", partTwoResult)
}

func partOne(filename string) int {

	pipeMap := parsePipeMap(filename)
	sIndex := findSIndex(pipeMap)
	sPipeType, stepCount := findSPipeType(pipeMap, sIndex)

	if sPipeType == "" {
		return -1
	}

	return stepCount / 2
}

func partTwo(filename string) float64 {

	pipeMap := parsePipeMap(filename)
	sIndex := findSIndex(pipeMap)
	sPipeType, _ := findSPipeType(pipeMap, sIndex)

	if sPipeType == "" {
		return -1
	}

	pipeMap[sIndex[0]][sIndex[1]] = sPipeType

	perimeterCoords := getPerimeterCoords(pipeMap, sIndex)

	area := areaByShoelaceFormula(perimeterCoords)

	perimeterCoordsCount := len(perimeterCoords)
	internalPoints := getInternalPointsCountByPicksTheorem(area, float64(perimeterCoordsCount))

	return internalPoints
}

func areaByShoelaceFormula(listOfPoints [][]int) float64 {
	listOfPoints = append(listOfPoints, listOfPoints[0])
	downMultipleList := []int{}
	for i := 0; i < len(listOfPoints)-1; i++ {
		multiple := listOfPoints[i][0] * listOfPoints[i+1][1]
		downMultipleList = append(downMultipleList, multiple)
	}

	upMultipleList := []int{}
	for i := 0; i < len(listOfPoints)-1; i++ {
		multiple := listOfPoints[i+1][0] * listOfPoints[i][1]
		upMultipleList = append(upMultipleList, multiple)
	}

	downSum := 0
	upSum := 0
	for i := 0; i < len(downMultipleList); i++ {
		downSum += downMultipleList[i]
		upSum += upMultipleList[i]
	}

	return math.Abs(float64(downSum)-float64(upSum)) / 2
}

func getInternalPointsCountByPicksTheorem(a float64, b float64) float64 {
	return (a + 1) - (b / 2)
}

func getPerimeterCoords(pipeMap [][]string, startCoords []int) [][]int {
	prevCoords := startCoords
	currentCoords := startCoords
	perimeterCoords := [][]int{}
	counter := 0
	for {
		pipeType := pipeMap[currentCoords[0]][currentCoords[1]]
		if isCoordsSimilar(currentCoords, startCoords) && counter > 0 {
			break
		}
		for _, d := range pipeDirections[pipeType] {
			nextCoords, _ := getCoordsForDirection(currentCoords, d, pipeMap)
			if !isCoordsSimilar(prevCoords, nextCoords) {
				perimeterCoords = append(perimeterCoords, currentCoords)
				prevCoords = currentCoords
				currentCoords = nextCoords
				break
			}
		}
		counter++
	}
	return perimeterCoords
}

func parsePipeMap(pipeMapFilename string) [][]string {
	file, err := os.ReadFile(pipeMapFilename)
	if err != nil {
		log.Fatal("Couldn't read file")
	}
	pipeMapStr := string(file)

	pipeMapLines := strings.Split(pipeMapStr, "\n")
	pipeMap := [][]string{}
	for _, l := range pipeMapLines {
		pipeMapLineArr := strings.Split(l, "")
		pipeMap = append(pipeMap, pipeMapLineArr)
	}
	return pipeMap
}

func findSIndex(pipeMap [][]string) []int {
	x := 0
	y := 0
	sFound := false
	for i := 0; i <= len(pipeMap)-1; i++ {
		for j := 0; j <= len(pipeMap[x])-1; j++ {
			if pipeMap[i][j] == "S" {
				y = j
				sFound = true
				break
			}
		}
		if sFound {
			x = i
			break
		}
	}
	return []int{x, y}
}

func findSPipeType(pipeMap [][]string, pipeCoords []int) (string, int) {
	sPipeType := ""
	loopStepCount := 0
	for pipeType, _ := range pipeDirections {
		newPipeMap := pipeMap
		newPipeMap[pipeCoords[0]][pipeCoords[1]] = pipeType

		isLoopbackPossible, stepCount := checkLoopback(newPipeMap, pipeCoords, pipeCoords, pipeCoords, [][]int{}, 0)
		loopStepCount = stepCount

		if isLoopbackPossible {
			sPipeType = pipeType
			break
		}
	}
	return sPipeType, loopStepCount
}

func checkLoopback(pipeMap [][]string, startingCoords []int, currentCoords []int, prevCoords []int, traversed [][]int, stepCount int) (bool, int) {
	if isCoordsSimilar(startingCoords, currentCoords) &&
		stepCount > 1 &&
		isPipeJoinedBackwards(pipeMap, startingCoords, prevCoords) {
		return true, stepCount
	} else if isCoordsSimilar(startingCoords, currentCoords) &&
		stepCount > 1 &&
		!isPipeJoinedBackwards(pipeMap, startingCoords, prevCoords) {
		return false, stepCount
	}

	totalLen := 0
	for _, m := range pipeMap {
		totalLen += len(m)
	}

	if stepCount > totalLen {
		return false, stepCount
	}

	isLoopbackPossible := false
	currentPipeType, _ := getPipe(pipeMap, currentCoords)
	for _, d := range pipeDirections[currentPipeType] {
		nextCoords, nextCoordsErr := getCoordsForDirection(currentCoords, d, pipeMap)
		if nextCoordsErr != nil {
			continue
		}

		if isCoordsSimilar(prevCoords, nextCoords) {
			continue
		}
		_, nextPipeErr := getPipe(pipeMap, nextCoords)
		if nextPipeErr != nil {
			continue
		}

		stepCount++
		isLoopbackPossible, stepCount = checkLoopback(pipeMap, startingCoords, nextCoords, currentCoords, traversed, stepCount)

		if isLoopbackPossible == true {
			return isLoopbackPossible, stepCount
		}
	}

	return isLoopbackPossible, stepCount
}

func isPipeJoinedBackwards(pipeMap [][]string, pipeCoords, backPipeCoords []int) bool {
	pipeType, err := getPipe(pipeMap, pipeCoords)
	if err != nil {
		return false
	}
	pipeDirOneCoords, _ := getCoordsForDirection(pipeCoords, pipeDirections[pipeType][0], pipeMap)
	pipeDirTwoCoords, _ := getCoordsForDirection(pipeCoords, pipeDirections[pipeType][1], pipeMap)
	if isCoordsSimilar(backPipeCoords, pipeDirOneCoords) || isCoordsSimilar(backPipeCoords, pipeDirTwoCoords) {
		return true
	}

	return false
}

func getCoordsForDirection(coords []int, direction string, pipeMap [][]string) ([]int, error) {
	nextX := coords[0] + directions[direction][0]
	nextY := coords[1] + directions[direction][1]
	if nextX < 0 || nextY < 0 {
		return nil, errors.New("Coords not possible")
	} else if nextX > len(pipeMap)-1 || nextY > len(pipeMap[nextX])-1 {
		return nil, errors.New("Coords not possible")
	}
	return []int{nextX, nextY}, nil
}

func isCoordsSimilar(coords1 []int, coords2 []int) bool {
	return reflect.DeepEqual(coords1, coords2)
}

func getPipe(pipeMap [][]string, coords []int) (string, error) {
	pipe := pipeMap[coords[0]][coords[1]]
	if pipe == "." {
		return pipe, errors.New("Not a pipe")
	}
	return pipe, nil
}
