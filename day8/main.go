package main

import (
    "log"
    "bufio"
    "os"
    "strings"
	"regexp"
)

func main() {
    filename := "input.txt"

    partOneResult := partOne(filename)
    log.Printf("Part One: %v", partOneResult)

    // partTwoResult := partTwo(filename)
    // log.Printf("Part Two: %v", partTwoResult)

	partTwoResult := partTwoLCM(filename)
    log.Printf("Part Two: %v", partTwoResult)
}

func partOne(filename string) int {
	steps, nodes := parseMap(filename)

	stepMap := map[string]int {
		"L": 0,
		"R": 1,
	}

	nextNode := "AAA"
	stepCounter := 0
	for step := 0; nextNode != "ZZZ"; step++ {
		if step == len(steps) {step = 0}
		nextNode = nodes[nextNode][stepMap[steps[step]]]
		stepCounter++
	}

	return stepCounter
}

func partTwo(filename string) int {
	steps, startingNodes, nodes := parseMap2(filename)

	stepMap := map[string]int {
		"L": 0,
		"R": 1,
	}

	currentNodes := startingNodes
	stepCounter := 0
	for step := 0; ; step++ {
		if step == len(steps) {step = 0}
		for nodeIndex, node := range currentNodes {
			currentNodes[nodeIndex] = nodes[node][stepMap[steps[step]]]
		}
		stepCounter++
		finished := true
		for _, node := range currentNodes {
			if string(node[len(node) - 1]) != "Z" {
				finished = false
				break
			}
		}
		if finished {break}
	}

	return stepCounter
}

// Implemented LCM method from advent of code subreddit because
// the iterative brute force method takes too long.
func partTwoLCM(filename string) int {
	steps, startingNodes, nodes := parseMap2(filename)

	stepMap := map[string]int {
		"L": 0,
		"R": 1,
	}

	stepCounts := []int{}

	for _, startingNode := range startingNodes {
		nextNode := startingNode
		stepCounter := 0
		for step := 0; string(nextNode[len(nextNode) - 1]) != "Z"; step++ {
			if step == len(steps) {step = 0}
			nextNode = nodes[nextNode][stepMap[steps[step]]]
			stepCounter++
		}
		stepCounts = append(stepCounts, stepCounter)
	}

	lcm := LCM(stepCounts[0], stepCounts[1], stepCounts[2:]...)

	return lcm
}

func parseMap(mapFile string) (steps []string, nodes map[string][]string) {
	file, err := os.Open(mapFile)
    if err != nil {
        log.Fatal("Couldn't read file")
    }
    defer file.Close()

	steps = []string {}
	nodes = make(map[string][]string)
    scanner := bufio.NewScanner(file)
	lineNo := 0
    for scanner.Scan() {
		lineNo++
        line := scanner.Text()
		if line == "" {
			continue
		}

		if lineNo == 1 {
			steps = strings.Split(line, "")
			continue
		}

		re := regexp.MustCompile("[A-Z]+")
		nodeData := re.FindAllString(line, -1)
		nodes[nodeData[0]] = []string {nodeData[1], nodeData[2]}
	}

	return steps, nodes
}

func parseMap2(mapFile string) (steps []string, startingNodes []string, nodes map[string][]string) {
	file, err := os.Open(mapFile)
    if err != nil {
        log.Fatal("Couldn't read file")
    }
    defer file.Close()

	steps = []string {}
	nodes = make(map[string][]string)
	startingNodes = []string{}
	// finishingNodes = []string{}
    scanner := bufio.NewScanner(file)
	lineNo := 0
    for scanner.Scan() {
		lineNo++
        line := scanner.Text()
		if line == "" {
			continue
		}

		if lineNo == 1 {
			steps = strings.Split(line, "")
			continue
		}

		re := regexp.MustCompile("[A-Z1-9]+")
		nodeData := re.FindAllString(line, -1)
		nodes[nodeData[0]] = []string {nodeData[1], nodeData[2]}
		if (string(nodeData[0][len(nodeData[0]) - 1]) == "A") {
			startingNodes = append(startingNodes, nodeData[0])
		}
	}

	return steps, startingNodes, nodes
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}

func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}