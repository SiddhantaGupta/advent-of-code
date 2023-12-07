package main

import (
    "log"
    "bufio"
    "os"
    "strconv"
    "strings"
	"math"
	"time"
)

type SourceDestinationData struct {
	sourceName string
	destinationName string
	sourceDestinationRangeDataList []SourceDestinationRangeData
}

type SourceDestinationRangeData struct {
	sourceRangeStart int
	destinationRangeStart int
	rangeLength int
}

/*
seeds: [seeds to be planted]

<source>-to-<destination> map:
<destination-range-start> <source-range-start> <length-including-range-start>
*/
// anything not in map maps 1:1. i.e., if source 24 is not in range the destination will be 24 as well.
func main() {
    filename := "input.txt"

    partOneResult := partOne(filename)
    log.Printf("Part One: %v", partOneResult)

    partTwoResult := partTwo(filename)
    log.Printf("Part Two: %v", partTwoResult)
}

func partOne(filename string) int {
	startTime := time.Now()

	seedsToBePlanted := getSeedsToBePlanted(filename)
	sourceDestinationDataList := getSourceDestinationDataList(filename)

	lowestLocation := math.MaxInt
	for _, seed := range seedsToBePlanted {
		location := getSeedLocation(sourceDestinationDataList, seed)
		if (location < lowestLocation) {
			lowestLocation = location
		}
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	log.Println("partOne: ", executionTime)

	return lowestLocation
}

func partTwo(filename string) int {
	startTime := time.Now()

	seedRangesToBePlanted := getSeedsToBePlanted(filename)
	sourceDestinationDataList := getSourceDestinationDataList(filename)

	calculatedLocation := make(chan int)
	for i := 0; i < len(seedRangesToBePlanted); i += 2 {
		go func (sourceDestinationDataList []SourceDestinationData, seedRangeStart int, seedRangeLength int, calculatedLocation chan int) {
			lowestLocation := math.MaxInt
		
			for i := 0; i <= seedRangeLength; i++ {
				location := getSeedLocation(sourceDestinationDataList, seedRangeStart+i)
				if (location < lowestLocation) {
					lowestLocation = location
				}
			}
			
			calculatedLocation <- lowestLocation
		}(sourceDestinationDataList, seedRangesToBePlanted[i], seedRangesToBePlanted[i+1], calculatedLocation)
	}

	lowestLocation := math.MaxInt
	for i := 0; i < len(seedRangesToBePlanted); i += 2 {
		location := <- calculatedLocation
		if (location < lowestLocation) {
			lowestLocation = location
		}
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	log.Println("partTwo: ", executionTime)

	return lowestLocation
}

func getSeedLocation(sourceDestinationDataList []SourceDestinationData, seed int) int {
	sourceValue := seed	
    for _, sourceDestinationData := range sourceDestinationDataList {
		for _, sourceDestinationRangeData := range sourceDestinationData.sourceDestinationRangeDataList {
			if (sourceValue >= sourceDestinationRangeData.sourceRangeStart && sourceValue <= sourceDestinationRangeData.sourceRangeStart + sourceDestinationRangeData.rangeLength) {
				rangeDiff := sourceValue - sourceDestinationRangeData.sourceRangeStart
				sourceValue = sourceDestinationRangeData.destinationRangeStart + rangeDiff
				break
			}
		}
	}
	return sourceValue
}


func getSeedsToBePlanted(sourceDestinationDataFilename string) []int {
	file, err := os.Open(sourceDestinationDataFilename)

    if err != nil {
        log.Fatal("Couldn't read file")
    }

    defer file.Close()

    scanner := bufio.NewReader(file)

	line, err := scanner.ReadString('\n')

	if err != nil {
		log.Fatal("Could not real line")
	}

	seedsStr := strings.TrimSpace(strings.Split(line, ":")[1])
	seedStrList := strings.Split(seedsStr, " ")

	var seeds []int
	for _, seedStr := range seedStrList {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {continue}
		seeds = append(seeds, seed)
	}

	return seeds
}

func getSourceDestinationDataList(sourceDestinationDataFilename string) []SourceDestinationData {
	file, err := os.Open(sourceDestinationDataFilename)

    if err != nil {
        log.Fatal("Couldn't read file")
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

	var sourceName string
	var destinationName string
	var sourceDestinationDataList []SourceDestinationData
	mapCounter := -1
	
    for scanner.Scan() {
        line := scanner.Text()

		if line == "" || strings.Contains(line, "seeds: ") {
			continue
		}

		if strings.Contains(line, "map:") {
			mapCounter++
			mapSourceAndDestinationNames := strings.Split(strings.Split(line, " ")[0], "-to-")
			sourceName = mapSourceAndDestinationNames[0]
			destinationName = mapSourceAndDestinationNames[1]

			sourceDestinationData := SourceDestinationData {
				sourceName: sourceName,
				destinationName: destinationName,
				sourceDestinationRangeDataList: []SourceDestinationRangeData{},
			}

			sourceDestinationDataList = append(sourceDestinationDataList, sourceDestinationData)
			continue
		}

		sourceToDestinationMapValues := strings.Split(line, " ")
		destinationRangeStart, _ := strconv.Atoi(sourceToDestinationMapValues[0])
		sourceRangeStart, _ := strconv.Atoi(sourceToDestinationMapValues[1])
		rangeLength, _ := strconv.Atoi(sourceToDestinationMapValues[2])

		sourceDestinationRangeData := SourceDestinationRangeData {
			sourceRangeStart: sourceRangeStart,
			destinationRangeStart: destinationRangeStart,
			rangeLength: rangeLength,
		}

		sourceDestinationDataList[mapCounter].sourceDestinationRangeDataList = append(sourceDestinationDataList[mapCounter].sourceDestinationRangeDataList, sourceDestinationRangeData)		
	}

	return sourceDestinationDataList
}