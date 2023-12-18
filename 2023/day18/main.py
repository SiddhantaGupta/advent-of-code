#!/usr/bin/python3
import os
from copy import deepcopy
import time
import math

directions = {
    "U": [-1, 0],
    "D": [1, 0],
    "L": [0, -1],
    "R": [0, 1]
}

dir = ["R", "D", "L", "U"]


def main():
    data = None
    with open("input", "r") as file:
        data = file.read()

    input = data.split(os.linesep)

    start_time = time.time()
    print("Part One: ", partOne(input))
    print("--- %s seconds ---" % (time.time() - start_time))

    start_time = time.time()
    print("Part Two: ", partTwo(input))
    print("--- %s seconds ---" % (time.time() - start_time))

    # start_time = time.time()
    # print("Part Two Old: ", partTwoOld(input))
    # print("--- %s seconds ---" % (time.time() - start_time))

    return 0


def partOne(input):
    coords = (0, 0)
    diggedCoords = []

    for r in input:
        d, s, c = r.split(" ")
        s = int(s)
        c = c.replace("(", "").replace(")", "")

        for i in range(s):
            coords = nextCoords(coords, d)
            diggedCoords.append((coords[0], coords[1], c))

    area = areaByShoelaceFormula(diggedCoords)
    internalPointsCount = getInternalPointsCountByPicksTheorem(
        area, len(diggedCoords))

    return math.floor(len(diggedCoords) + internalPointsCount)


def partTwo(input):
    coords = (0, 0)
    shoeLaceDiggedCoords = [(0, 0)]
    shoeLaceUpSum = 0
    shoeLaceDownSum = 0

    count = 0
    for r in input:
        d, s, c = r.split(" ")
        s = int(c.replace("(", "").replace(")", "").replace("#", "")[:-1], 16)
        d = dir[int(c.replace(")", "")[-1])]

        for i in range(s):
            count += 1
            coords = nextCoords(coords, d)
            shoeLaceDiggedCoords.append((coords[0], coords[1]))
            downMultiple = shoeLaceDiggedCoords[0][0] * \
                shoeLaceDiggedCoords[1][1]
            upMultiple = shoeLaceDiggedCoords[1][0] * \
                shoeLaceDiggedCoords[0][1]
            shoeLaceDownSum += downMultiple
            shoeLaceUpSum += upMultiple
            shoeLaceDiggedCoords.pop(0)

    area = abs(shoeLaceDownSum - shoeLaceUpSum) / 2
    internalPointsCount = getInternalPointsCountByPicksTheorem(
        area, count)

    return math.floor(count + internalPointsCount)


def areaByShoelaceFormula(listOfPoints):
    listOfPoints.append(listOfPoints[0])
    downMultipleList = []
    for i in range(len(listOfPoints) - 1):
        multiple = listOfPoints[i][0] * listOfPoints[i+1][1]
        downMultipleList.append(multiple)

    upMultipleList = []
    for i in range(len(listOfPoints) - 1):
        multiple = listOfPoints[i+1][0] * listOfPoints[i][1]
        upMultipleList.append(multiple)

    downSum = 0
    upSum = 0
    for i in range(len(downMultipleList)):
        downSum += downMultipleList[i]
        upSum += upMultipleList[i]

    return abs(downSum - upSum) / 2


def getInternalPointsCountByPicksTheorem(area, permiterCoordsCount):
    return (area + 1) - (permiterCoordsCount / 2)


def nextCoords(coords, d=""):
    d = d if d != "" else coords[2]
    return (coords[0] + directions[d][0], coords[1] + directions[d][1])


if __name__ == "__main__":
    main()
