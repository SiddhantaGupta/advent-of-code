#!/usr/bin/python3
import os
from pprint import pprint
from copy import deepcopy

directions = {
    "U": [-1, 0],
    "D": [1, 0],
    "L": [0, -1],
    "R": [0, 1]
}


def main():
    data = None
    with open("input", "r") as file:
        data = file.read()

    input = data.split(os.linesep)
    for i in range(len(input)):
        input[i] = list(input[i])

    print("Part One: ", partOne(input))
    print("Part Two: ", partTwo(input))
    return 0


def partOne(input):
    return computeEnergizedCells(input, (0, 0, "R"))


def partTwo(input):
    startPosList = []
    for i in range(len(input)):
        for j in range(len(input[i])):
            if (i != 0 and i != len(input) - 1) and (j != 0 and j != len(input[j]) - 1):
                continue

            if i == 0:
                startPosList.append((i, j, "D"))
            elif i == len(input) - 1:
                startPosList.append((i, j, "U"))

            if j == 0:
                startPosList.append((i, j, "R"))
            elif j == len(input[i]) - 1:
                startPosList.append((i, j, "L"))

    energizedCells = []
    for startPos in startPosList:
        energizedCells.append(computeEnergizedCells(input, startPos))

    return max(energizedCells)


def computeEnergizedCells(m, sb):
    startBeam = (sb[0] - directions[sb[2]][0],
                 sb[1] - directions[sb[2]][1],
                 sb[2])
    beams = [startBeam]
    beamHist = [startBeam]
    energizedCells = set()
    while True:
        indexToDel = []
        if len(beams) <= 0:
            break

        for b, _ in enumerate(beams):
            beams[b] = nextBeam(beams[b])

            if beams[b][0] < 0 or beams[b][0] > len(m) - 1 or beams[b][1] < 0 or beams[b][1] > len(m[0]) - 1:
                indexToDel.append(b)
                continue

            energizedCells.add((beams[b][0], beams[b][1]))

            match getTile(m, beams[b]):
                case "\\":
                    x, y, d = beams[b]
                    if beams[b][2] == "D":
                        beams[b] = (x, y, "R")
                    elif beams[b][2] == "U":
                        beams[b] = (x, y, "L")
                    elif beams[b][2] == "R":
                        beams[b] = (x, y, "D")
                    elif beams[b][2] == "L":
                        beams[b] = (x, y, "U")

                case "/":
                    x, y, d = beams[b]
                    if beams[b][2] == "D":
                        beams[b] = (x, y, "L")
                    elif beams[b][2] == "U":
                        beams[b] = (x, y, "R")
                    elif beams[b][2] == "R":
                        beams[b] = (x, y, "U")
                    elif beams[b][2] == "L":
                        beams[b] = (x, y, "D")

                case "|":
                    if beams[b][2] in "LR":
                        indexToDel.append(b)
                        newBeam1 = (beams[b][0], beams[b][1], "U")
                        newBeam2 = (beams[b][0], beams[b][1], "D")

                        if newBeam1 in beamHist and newBeam2 in beamHist:
                            continue

                        beams.append(newBeam1)
                        beams.append(newBeam2)
                        beamHist.append(deepcopy(newBeam1))
                        beamHist.append(deepcopy(newBeam2))

                case "-":
                    if beams[b][2] in "UD":
                        indexToDel.append(b)
                        newBeam1 = (beams[b][0], beams[b][1], "L")
                        newBeam2 = (beams[b][0], beams[b][1], "R")

                        if newBeam1 in beamHist and newBeam2 in beamHist:
                            continue

                        beams.append(newBeam1)
                        beams.append(newBeam2)
                        beamHist.append(deepcopy(newBeam1))
                        beamHist.append(deepcopy(newBeam2))

        indexToDel.sort()

        for i in range(len(indexToDel) - 1, -1, -1):
            del beams[indexToDel[i]]

    return len(energizedCells)


def nextBeam(beam, d=""):
    d = d if d != "" else beam[2]
    return (beam[0] + directions[d][0], beam[1] + directions[d][1], d)


def getTile(m, b):
    return m[b[0]][b[1]]


if __name__ == "__main__":
    main()
