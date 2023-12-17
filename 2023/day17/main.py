#!/usr/bin/python3
import os
from copy import deepcopy
import time
import resource
import sys
from pprint import pprint

resource.setrlimit(resource.RLIMIT_STACK, [0x10000000, resource.RLIM_INFINITY])
sys.setrecursionlimit(0x100000)

directions = {
    "N": (-1, 0, "S"),
    "S": (1, 0, "N"),
    "E": (0, 1, "W"),
    "W": (0, -1, "E")
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

    # start_time = time.time()
    # print("Part Two: ", partTwo(input))
    # print("--- %s seconds ---" % (time.time() - start_time))

    return 0


def partOne(input):
    return pathOfLeastHeatDissapation(input)


def pathOfLeastHeatDissapation(m):
    heatLossList = []
    reached = []

    def walk(m, block, stepCount, heatLoss):
        nxtBlock = nextBlock(block)

        if nxtBlock[0] < 0 or nxtBlock[0] > len(m) - 1 or nxtBlock[1] < 0 or nxtBlock[1] > len(m[0]) - 1:
            return

        hl = heatLoss + int(m[nxtBlock[0]][nxtBlock[1]])

        a, b, c = nxtBlock
        nxtBlockReached = (a, b, c, hl)

        if nxtBlockReached in reached:
            return
        for x in heatLossList:
            if x <= hl:
                return
        if nxtBlock[0] == len(m) - 1 and nxtBlock[1] == len(m[0]) - 1:
            heatLossList.append(hl)
            return

        reached.append(nxtBlockReached)

        for direction in directions.keys():
            if direction == directions[nxtBlock[2]][2]:
                continue

            if direction == nxtBlock[2] and stepCount == 3:
                continue

            b = (nxtBlock[0], nxtBlock[1], direction)
            sc = stepCount + 1 if direction == nxtBlock[2] else 0

            walk(m, b, sc, hl)

    walk(m, (0, 0, "E"), 0, 0)
    walk(m, (0, 0, "S"), 0, 0)

    return min(heatLossList)


def nextBlock(block, d=""):
    d = d if d != "" else block[2]
    return (block[0] + directions[d][0], block[1] + directions[d][1], d)


if __name__ == "__main__":
    main()
