package main

import (
    "log"
    "bufio"
    "os"
    "strconv"
    "strings"
)

type Bag struct {
    id int
    cubes map[string]int
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

    var games []Bag

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        res := strings.Split(line, ":")

        bag := Bag{id: 0, cubes: make(map[string]int)}
        bag.id, _ = strconv.Atoi(strings.Split(res[0], " ")[1])
        reveals := strings.Split(res[1], ";")

        for revealIndex := 0; revealIndex < len(reveals); revealIndex++ {

            reveal := reveals[revealIndex]
            cubeStrs := strings.Split(reveal, ",")

            for _, cubeStr := range cubeStrs {

                cubeStr = strings.TrimSpace(cubeStr)
                cubeCount, _ := strconv.Atoi(strings.Split(cubeStr, " ")[0])
                cubeColor := strings.Split(cubeStr, " ")[1]

                if bag.cubes[cubeColor] < cubeCount {
                    bag.cubes[cubeColor] = cubeCount
                }
                
            }
        }
        games = append(games, bag)
    }
    loadedBag := make(map[string]int)
    loadedBag["red"] = 12
    loadedBag["green"] = 13
    loadedBag["blue"] = 14

    sumOfIdsOfPossibleGames := 0
    for _, game := range games {
        isGamePossible := true
        for cubeColor, cubeCount := range game.cubes {
            if cubeCount > loadedBag[cubeColor] {
                isGamePossible = false
            }
        }
        if isGamePossible {
            sumOfIdsOfPossibleGames += game.id
        }
    }

    return sumOfIdsOfPossibleGames
}

func partTwo(filename string) int {
    file, err := os.Open(filename)

    if err != nil {
        log.Fatal("Couldn't read file")
    }

    defer file.Close()

    var games []Bag

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        res := strings.Split(line, ":")

        bag := Bag{id: 0, cubes: make(map[string]int)}
        bag.id, _ = strconv.Atoi(strings.Split(res[0], " ")[1])
        reveals := strings.Split(res[1], ";")

        for revealIndex := 0; revealIndex < len(reveals); revealIndex++ {

            reveal := reveals[revealIndex]
            cubeStrs := strings.Split(reveal, ",")

            for _, cubeStr := range cubeStrs {

                cubeStr = strings.TrimSpace(cubeStr)
                cubeCount, _ := strconv.Atoi(strings.Split(cubeStr, " ")[0])
                cubeColor := strings.Split(cubeStr, " ")[1]

                if bag.cubes[cubeColor] < cubeCount {
                    bag.cubes[cubeColor] = cubeCount
                }
                
            }
        }
        games = append(games, bag)
    }

    var cubePowerSum int

    for _, game := range games {
        cubePower := 1
        for _, cubeCount := range game.cubes {
            cubePower *= cubeCount
        }
        cubePowerSum += cubePower
    }

    return cubePowerSum
}