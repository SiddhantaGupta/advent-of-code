#!/usr/bin/python3
import os


def main():
    data = None
    with open("input", "r") as file:
        data = file.read()

    input = data.split(os.linesep)
    for i in range(0, len(input)):
        input[i] = list(input[i])

    print("Part One: ", partOne(input))
    print("Part Two: ", partTwo(input, 1000000000))
    return 0


def partOne(input):
    tm = transposeMatrixRight(input)
    sm = shiftRocks(tm)
    w = calculateTotalRockWeight(sm)
    return w


def partTwo(input, n):
    cache = {}
    hist = []
    m = input
    load = 0
    for i in range(0, n):
        for j in range(0, 4):
            m = transposeMatrixRight(m)
            m = shiftRocks(m)

        key = ''.join(c for r in m for c in r)
        if key in cache:
            """ explaination for finding which calc in history is the ans: 
            # complete rotation point in cycle = ((no. of times to rotate - rotations left of cycle) % cycle length) + rotations left of cycle
            # index of ans in history = complete rotation point in cycle - 1

            1. length of cycle = cycle end index - cycle start index

            2. rotations before cycle window = cycle start index // due to 0 indexing of arrays

            # rotations before cycle window will be added later
            3. number of times to rotate after reaching cycle window = number of times to rotate - rotations before cycle window

            # point in cycle window where final rotation will end
            4. final rotation point in cycle = number of times to rotate after reaching cycle window % length of cycle

            5. complete rotation point in sequence = final rotation point in cycle + rotations before cycle window

            # 1 has been subracted in the below eq to find the index since arrays are 0 indexed
            5. complete rotation index in sequence hist = complete rotation point in sequence - 1 """

            cycleStartIndex = cache[key]
            cycleEndIndex = i
            cycleLength = cycleEndIndex - cycleStartIndex

            rotationsLeftOfLoop = cycleStartIndex
            rotationsInCycle = n - rotationsLeftOfLoop

            lastRotationInCycleIndex = rotationsInCycle % cycleLength
            fullRotationEndCycleIndex = lastRotationInCycleIndex + rotationsLeftOfLoop - 1

            load = hist[fullRotationEndCycleIndex]
            break
        hist.append(calculateTotalRockWeight(transposeMatrixRight(m)))
        cache[key] = i

    return load


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
