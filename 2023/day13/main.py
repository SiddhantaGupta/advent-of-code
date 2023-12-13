#!/usr/bin/python3
import os
from pprint import pprint

dataArr = None
with open("input", "r") as file:
    d = file.read()
    dataArr = d.split(os.linesep + os.linesep)

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

def smudgyReflection(a, b):
    for i in range(0, len(a)):
        ax = list(a)
        ax[i] = "#" if ax[i] == "." else "."
        ax = "".join(ax)
        if ax == b:
            return True

    return False

def checkReflection(m, i, smudge):
    l, r = i, i+1
    diff = 1
    while True:
        if l < 0 or r > len(m) - 1:
            if smudge > 0:
                return(-1, -1)
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

def main(input):
    hs = 0
    vs = 0
    for data in input:
        hm = getHorizontalMatrix(data)
        vm = getVerticalMatrix(data)
        hr = findLineOfReflextion(hm, 1)
        vr = findLineOfReflextion(vm, 1)

        if hr != (-1, -1):
            hs += (hr[0] + 1) * 100

        if vr != (-1, -1):
            vs += vr[0] + 1

    print(hs + vs)
    return 0

if __name__ == "__main__":
    main(dataArr)