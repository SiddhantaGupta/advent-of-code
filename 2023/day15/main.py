#!/usr/bin/python3
import os
from pprint import pprint


def main():
    data = None
    with open("input", "r") as file:
        data = file.read()

    input = data.split(",")

    print("Part One: ", partOne(input))
    print("Part Two: ", partTwo(input))
    return 0


def partOne(input):
    h = 0
    for step in input:
        h += hash(step)

    return h


def partTwo(input):
    boxes = []
    for i in range(0, 256):
        boxes.append([])

    for step in input:
        if "=" in step:
            label, lens = step.split("=")
            box = hash(label)
            labelIndex = getLabelIndexInBox(boxes[box], label)
            if labelIndex >= 0:
                boxes[box][labelIndex] = (label, lens)
            else:
                boxes[box].append((label, lens))

        elif "-" in step:
            label, _ = step.split("-")
            box = hash(label)
            labelIndex = getLabelIndexInBox(boxes[box], label)
            if labelIndex >= 0:
                del boxes[box][labelIndex]

    totalFocusingPower = 0
    for i in range(0, len(boxes)):
        for j in range(0, len(boxes[i])):
            totalFocusingPower += (1 + i) * (j + 1) * int(boxes[i][j][1])

    return totalFocusingPower


def getLabelIndexInBox(box, label):
    for p in range(0, len(box)):
        if box[p][0] == label:
            return p
    return -1


def hash(str):
    l = list(str)
    h = 0
    for c in l:
        h += ord(c)
        h *= 17
        h = h % 256

    return h


if __name__ == "__main__":
    main()
