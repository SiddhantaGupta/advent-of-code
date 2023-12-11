package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input"

	partOneResult := partOne(filename)
	log.Printf("Part One: %v", partOneResult)

	partTwoResult := partTwo(filename)
	log.Printf("Part Two: %v", partTwoResult)
}

func partOne(filename string) int {
	galaxyData, err := readFullFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	galaxyMatrix := parseFileStringToMatrix(galaxyData)
	expandedGalaxyMatrix := expandGalaxyMatrix(galaxyMatrix, 2)
	_, galaxyCoordsMap := numberGalaxiesInMatrix(expandedGalaxyMatrix)
	pairs := getCoordsPairList(galaxyCoordsMap)

	shortestPathStepCounts := []int{}
	for _, pair := range pairs {
		shortestPathStepCount := getShortestPathStepCount(pair[0], pair[1])
		shortestPathStepCounts = append(shortestPathStepCounts, shortestPathStepCount)
	}

	totalStepCount := 0
	for _, c := range shortestPathStepCounts {
		totalStepCount += c
	}

	return totalStepCount
}

func partTwo(filename string) int {
	galaxyData, err := readFullFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	galaxyMatrix := parseFileStringToMatrix(galaxyData)
	emptyCoords := getEmptyCoords(galaxyMatrix)
	_, galaxyCoordsMap := numberGalaxiesInMatrix(galaxyMatrix)
	pairs := getCoordsPairList(galaxyCoordsMap)

	shortestPathStepCounts := []int{}
	for _, pair := range pairs {
		shortestPathStepCount := getShortestPathStepCountWithExpansion(pair[0], pair[1], emptyCoords, 1000000)
		shortestPathStepCounts = append(shortestPathStepCounts, shortestPathStepCount)
	}

	totalStepCount := 0
	for _, c := range shortestPathStepCounts {
		totalStepCount += c
	}

	return totalStepCount
}

func getEmptyCoords(gM [][]string) [][]int {
	eM := gM
	eCs := [][]int{[]int{}, []int{}}

	// expand cols
	// emptyColCount := 0
	for i := 0; i <= len(eM[0])-1; i++ {
		emptyCol := true
		for j := 0; j <= len(eM)-1; j++ {
			// log.Println(i, j)
			if eM[j][i] != "." {
				emptyCol = false
				break
			}
		}
		if emptyCol {
			// col found i
			eCs[1] = append(eCs[1], i)
			// eCs[1] = append(eCs[1], i+(emptyColCount*n))
		}
	}

	// expand rows
	for i := 0; i <= len(eM)-1; i++ {
		emptyRow := true
		r := eM[i]
		for _, c := range r {
			if c != "." {
				emptyRow = false
				break
			}
		}
		if emptyRow {
			eCs[0] = append(eCs[0], i)
			// eCs[0] = append(eCs[0], i+(emptyColCount*n))
		}
	}

	return eCs
}

func getShortestPathStepCountWithExpansion(s []int, e []int, empty [][]int, expRate int) int {
	xl := int(math.Max(float64(s[0]), float64(e[0])))
	xs := int(math.Min(float64(s[0]), float64(e[0])))
	yl := int(math.Max(float64(s[1]), float64(e[1])))
	ys := int(math.Min(float64(s[1]), float64(e[1])))
	x := xl - xs
	y := yl - ys

	xExpCount := 0
	for _, co := range empty[0] {
		if co > xs && co < xl {
			xExpCount++
		}
	}
	x = (x - xExpCount) + (expRate * xExpCount)

	yExpCount := 0
	for _, co := range empty[1] {
		if co > ys && co < yl {
			yExpCount++
		}
	}
	y = (y - yExpCount) + (expRate * yExpCount)

	return x + y
}

func getShortestPathStepCount(s []int, e []int) int {
	r := e[0] - s[0]
	c := e[1] - s[1]
	if r < 0 {
		r = -r
	}
	if c < 0 {
		c = -c
	}
	return r + c
}

func getCoordsPairList(m map[int][]int) [][][]int {
	n := len(m)
	pairs := [][][]int{}
	for i := 1; i <= n-1; i++ {
		for j := i + 1; j <= n; j++ {
			pair1, _ := m[i]
			pair2, _ := m[j]
			pair := [][]int{pair1, pair2}
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func numberGalaxiesInMatrix(gM [][]string) ([][]string, map[int][]int) {
	ngM := gM
	n := 0
	gCs := map[int][]int{}
	for rI, r := range ngM {
		for cI, c := range r {
			if c == "#" {
				n++
				gCs[n] = []int{rI, cI}
				ngM[rI][cI] = strconv.Itoa(n)
			}
		}
	}
	return ngM, gCs
}

func expandGalaxyMatrix(gM [][]string, n int) [][]string {
	eM := gM
	eR := []string{}
	for i := 0; i < len(eM[0]); i++ {
		eR = append(eR, ".")
	}

	// expand cols
	for i := 0; i <= len(eM[0])-1; i++ {
		emptyCol := true
		for j := 0; j <= len(eM)-1; j++ {
			// log.Println(i, j)
			if eM[j][i] != "." {
				emptyCol = false
				break
			}
		}
		if emptyCol {
			for x := 1; x < n; x++ {
				for k := 0; k <= len(eM)-1; k++ {
					eM[k] = append(eM[k][:i+1], eM[k][i:]...)
					eM[k][i] = "."
				}
				i++
			}
		}
	}

	// expand rows
	for i := 0; i <= len(eM)-1; i++ {
		emptyRow := true
		r := eM[i]
		for _, c := range r {
			if c != "." {
				emptyRow = false
				break
			}
		}
		if emptyRow {
			for x := 1; x < n; x++ {
				eM = append(eM[:i+1], eM[i:]...)
				eM[i] = eR
				i++
			}
		}
	}

	return eM
}

func readFullFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	fileData := string(file)
	return fileData, nil
}

func parseFileStringToMatrix(fs string) [][]string {
	ls := strings.Split(fs, "\n")
	m := [][]string{}
	for _, l := range ls {
		byteS := strings.Split(l, "")
		stringS := []string{}
		for _, b := range byteS {
			stringS = append(stringS, string(b))
		}
		m = append(m, stringS)
	}
	return m
}
