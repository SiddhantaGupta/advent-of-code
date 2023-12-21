#!/usr/bin/python3
import os
import time
from copy import deepcopy


directions = {
    "N": (-1, 0),
    "S": (1, 0),
    "E": (0, 1),
    "W": (0, -1)
}


def main():
    data = None
    with open("example", "r") as file:
        data = file.read()

    input = data.split(os.linesep)
    for i in range(len(input)):
        input[i] = list(input[i])

    start_time = time.time()
    print("Part One: ", partOne(input))
    print("--- %s seconds ---" % (time.time() - start_time))

    start_time = time.time()
    print("Part Two: ", partTwo(input))
    print("--- %s seconds ---" % (time.time() - start_time))

    return 0


def partOne(input):
    startPos = findSIndex(input)
    if startPos == (-1, -1):
        return

    stepCount = 64
    tilesToTravel = [startPos]
    tilesToBeTravelled = set()
    for i in range(stepCount):
        while True:
            if len(tilesToTravel) <= 0:
                if i == stepCount - 1:
                    return len(tilesToBeTravelled)
                tilesToTravel = deepcopy(tilesToBeTravelled)
                tilesToBeTravelled = set()
                break
            coord = tilesToTravel.pop()
            for d in directions:
                nxtTile = nextTile(coord, d)
                if nxtTile[0] < 0 or nxtTile[0] > len(input) - 1 or nxtTile[1] < 0 or nxtTile[1] > len(input[0]) - 1:
                    continue
                elif input[nxtTile[0]][nxtTile[1]] == "." or input[nxtTile[0]][nxtTile[1]] == "S":
                    tilesToBeTravelled.add(nxtTile)


def partTwo(input):
    startPos = findSIndex(input)
    if startPos == (-1, -1):
        return

    stepCount = 26501365
    tilesToTravel = [startPos]
    tilesToBeTravelled = set()
    for i in range(stepCount):
        while True:
            if len(tilesToTravel) <= 0:
                if i == stepCount - 1:
                    return len(tilesToBeTravelled)

                tilesToTravel = deepcopy(tilesToBeTravelled)
                tilesToBeTravelled = set()
                break

            coord = tilesToTravel.pop()
            for d in directions:
                nxtTile = nextTile(coord, d)
                infiniteTile = getInfiniteTile(input, nxtTile)
                if input[infiniteTile[0]][infiniteTile[1]] == "." or input[infiniteTile[0]][infiniteTile[1]] == "S":
                    tilesToBeTravelled.add(nxtTile)


def getInfiniteTile(map, coords):
    coords = deepcopy(coords)
    if coords[0] < 0:
        rotation = len(map) - (abs(coords[0]) % len(map))
        rotation = 0 if rotation == len(map) else rotation
        coords = (rotation, coords[1])

    elif coords[0] > len(map) - 1:
        rotation = (coords[0] % len(map))
        coords = (rotation, coords[1])

    if coords[1] < 0:
        rotation = len(map[0]) - (abs(coords[1]) % len(map[0]))
        rotation = 0 if rotation == len(map[0]) else rotation
        coords = (coords[0], rotation)

    elif coords[1] > len(map[0]) - 1:
        rotation = (coords[1] % len(map[0]))
        coords = (coords[0], rotation)

    return coords


def findSIndex(gardenMap):
    for i in range(len(gardenMap)):
        for j in range(len(gardenMap[i])):
            if gardenMap[i][j] == "S":
                return (i, j)

    return (-1, -1)


def nextTile(block, d=""):
    d = d if d != "" else block[2]
    return (block[0] + directions[d][0], block[1] + directions[d][1])


if __name__ == "__main__":
    main()
