#!/usr/bin/python3
import os


def main():
    input = None
    with open("input", "r") as file:
        data = file.read()
        input = data.split(os.linesep + os.linesep)

    print(partOne(input))
    print(partTwo(input))
    return 0


def partOne(input):
    hs = 0
    vs = 0
    for item in input:
        hm = getHorizontalMatrix(item)
        vm = getVerticalMatrix(item)
        hr = findLineOfReflextion(hm, 0)
        vr = findLineOfReflextion(vm, 0)

        if hr != (-1, -1):
            hs += (hr[0] + 1) * 100

        if vr != (-1, -1):
            vs += vr[0] + 1

    return hs + vs


def partTwo(input):
    hs = 0
    vs = 0
    for item in input:
        hm = getHorizontalMatrix(item)
        vm = getVerticalMatrix(item)
        hr = findLineOfReflextion(hm, 1)
        vr = findLineOfReflextion(vm, 1)

        if hr != (-1, -1):
            hs += (hr[0] + 1) * 100

        if vr != (-1, -1):
            vs += vr[0] + 1

    return hs + vs


def getHorizontalMatrix(s):
    return s.split("\n")


def getVerticalMatrix(s):
    return transposeMatrix(getHorizontalMatrix(s))


def transposeMatrix(m):
    mx = [""] * len(m[0])
    for i in range(0, len(m)):
        for j in range(0, len(m[i])):
            mx[j] += m[i][j]

    return mx


def findLineOfReflextion(m, smudge):
    for i in range(0, len(m) - 1):
        reflection = checkReflection(m, i, smudge)
        if reflection != (-1, -1):
            return reflection

    return (-1, -1)


def checkReflection(m, i, smudge):
    l, r = i, i+1
    diff = 1
    while True:
        if l < 0 or r > len(m) - 1:
            if smudge > 0:
                return (-1, -1)
            else:
                return (i, i+1)
        if m[l] != m[r]:
            if smudge > 0:
                smudge -= 1
                if not smudgyReflection(m[l], m[r]):
                    return (-1, -1)
            else:
                return (-1, -1)
        l -= diff
        r += diff


def smudgyReflection(a, b):
    for i in range(0, len(a)):
        ax = list(a)
        ax[i] = "#" if ax[i] == "." else "."
        ax = "".join(ax)
        if ax == b:
            return True

    return False


if __name__ == "__main__":
    main()
