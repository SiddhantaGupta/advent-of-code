#!/usr/bin/python3
import os
from pprint import pprint


def main():
    data = None
    with open("example", "r") as file:
        data = file.read()

    input = data.split(os.linesep)
    for i in range(0, len(input)):
        input[i] = list(input[i])

    print(partOne(input))
    # print(partTwo(input))
    return 0


def partOne(input):
    tm = transposeMatrix(input)
    sm = shiftRocks(tm)
    w = calculateTotalRockWeight(sm)
    return w


def calculateTotalRockWeight(m):
    rockWeightSum = 0
    for i in range(0, len(m)):
        for j in range(0, len(m[i])):
            if m[i][j] == "O":
                rockWeightSum += 1 * (len(m[i]) - j)

    return rockWeightSum


def shiftRocks(m):
    for i in range(0, len(m)):
        for j in range(0, len(m[i])):
            if m[i][j] == "O":
                eidx = getLastEmptyIndex(m[i], j)
                if eidx < j:
                    m[i][eidx] = m[i][j]
                    m[i][j] = "."

    return m


def getLastEmptyIndex(r, j):
    if j == 0:
        return 0
    idx = j
    for i in range(j-1, -1, -1):
        if r[i] in "#O":
            break
        if r[i] == ".":
            idx = i

    return idx


def transposeMatrix(m):
    # using the below line copies ref and makes every iteration the same
    # mx = [[None] * len(m)] * len(m[0])
    mx = [[None for j in m] for i in m[0]]
    for i in range(0, len(m)):
        for j in range(0, len(m[i])):
            mx[j][i] = m[i][j]

    return mx


if __name__ == "__main__":
    main()
