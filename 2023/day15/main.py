#!/usr/bin/python3
import os


def main():
    data = None
    with open("example", "r") as file:
        data = file.read()

    input = data.split(",")

    print("Part One: ", partOne(input))
    return 0


def partOne(input):
    h = 0
    for step in input:
        h += hash(step)

    return h


def hash(str):
    strList = list(str)
    h = 0
    for c in strList:
        h += ord(c)
        h *= 17
        h = h % 256

    return h


if __name__ == "__main__":
    main()
