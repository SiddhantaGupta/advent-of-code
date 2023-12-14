#!/usr/bin/python3
import os
from pprint import pprint


def main():
    data = None
    with open("input", "r") as file:
        data = file.read()

    input = data.split(os.linesep)
    for i in range(0, len(input)):
        input[i] = list(input[i])

    print(partOne(input))
    # print(partTwo(input))
    return 0


def partOne(input):
    tm = transposeMatrixRight(input)
    sm = shiftRocks(tm)
    w = calculateTotalRockWeight(sm)
    return w


def calculateTotalRockWeight(m):
    rockWeightSum = 0
    for i in range(0, len(m)):
        for j in range(0, len(m[i])):
            if m[i][j] == "O":
                rockWeightSum += 1 * j + 1

    return rockWeightSum


def shiftRocks(m):
    for i in range(0, len(m)):
        for j in range(len(m[i]) - 1, -1, -1):
            if m[i][j] == "O":
                eidx = getLastEmptyIndex(m[i], j)
                if eidx > j:
                    m[i][eidx] = m[i][j]
                    m[i][j] = "."

    return m


def getLastEmptyIndex(r, j):
    idx = j
    if j == len(r) - 1:
        return idx

    for i in range(j+1, len(r)):
        if r[i] in "#O":
            break
        if r[i] == ".":
            idx = i

    return idx


def transposeMatrixLeft(m):
    # using the below line copies ref and makes every iteration the same
    # mx = [[None] * len(m)] * len(m[0])
    mx = [[None for j in m] for i in m[0]]
    for i in range(0, len(m)):
        for j in range(0, len(m[i])):
            mx[j][i] = m[i][j]
    return mx


def transposeMatrixRight(m):
    mx = [[None for j in m] for i in m[0]]
    for i in range(0, len(m)):
        for j in range(0, len(m[i])):
            mx[j][len(m[0]) - 1 - i] = m[i][j]
    return mx


if __name__ == "__main__":
    main()
