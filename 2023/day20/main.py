#!/usr/bin/python3
import os
import time
from math import lcm
from copy import deepcopy

moduleTypes = {
    "broadcaster": 1,
    "flipFlop": 2,
    "conjunction": 3,
}

flipFlopStates = {
    "off": 1,
    "on": 2
}

pulseTypes = {
    "low": 1,
    "high": 2
}

def main():
    data = None
    with open("input", "r") as file:
        data = file.read()

    input = data.split(os.linesep)

    m = {}
    conjunctionModules = []
    for r in input:
        module, outputs = r.split(" -> ")
        outputs = outputs.split(", ")
        name = ""

        if "%" in module:
            name = module.replace("%", "")
            m[name] = {}
            m[name]["type"] = moduleTypes["flipFlop"]
            m[name]["state"] = flipFlopStates["off"]
        elif "&" in module:
            name = module.replace("&", "")
            m[name] = {}
            m[name]["type"] = moduleTypes["conjunction"]
            m[name]["inputs"] = []
            conjunctionModules.append(module.replace("&", ""))
        else:
            name = module
            m[name] = {}
            m[name]["type"] = moduleTypes["broadcaster"]

        m[name]["outputs"] = outputs


    conjunctionInputs = {}
    for cm in conjunctionModules:
        conjunctionInputs[cm] = {}
    
    for mod in m:
        for cm in conjunctionModules:
            if cm in m[mod]["outputs"]:
                conjunctionInputs[cm][mod] = pulseTypes["low"]

    for ci in conjunctionInputs:
        m[ci]["inputs"] = conjunctionInputs[ci]


    start_time = time.time()
    print("Part One: ", partOne(m))
    print("--- %s seconds ---" % (time.time() - start_time))

    start_time = time.time()
    print("Part Two: ", partTwo(m))
    print("--- %s seconds ---" % (time.time() - start_time))

    return 0


def partOne(input):
    buttonPressCount = 1000
    lpcs = 0
    hpcs = 0

    for i in range(buttonPressCount):
        highPulseCount, lowPulseCount, input = broadcast(input)
        lpcs += lowPulseCount + 1
        hpcs += highPulseCount

    return lpcs * hpcs


def partTwo(input):
    """ for part 2 I took hint for LCM solution from reddit:
    The final module is a conjunction and will only send a low beam to the end module
    when all input's last pulse was a high pulse.
    every input cycles to send a high pulse at an interval and the LCM of all the cycles
    is when the condition will meet."""

    buttonPressCount = 0
    cycleCount = []

    while True:
        buttonPressCount += 1
        input, cycleCount = broadcast2(input, cycleCount, buttonPressCount)
        if len(cycleCount) == len(input["rs"]["inputs"]):
            break

    return lcm(*cycleCount)


def broadcast(modules):
    m = deepcopy(modules)
    pulses = []

    for output in m["broadcaster"]["outputs"]:
        pulses.append(("broadcast", output, pulseTypes["low"]))

    highPulseCount = 0
    lowPulseCount = 0
    while True:
        if len(pulses) <= 0:
            break

        src, dest, pType = pulses.pop(0)

        if pType == pulseTypes["low"]:
            lowPulseCount += 1
        else:
            highPulseCount += 1

        if dest not in m:
            continue
        elif m[dest]["type"] == moduleTypes["flipFlop"]:
            if pType == pulseTypes["low"]:
                m[dest]["state"] = toggleFlipFlop(m[dest])
                if m[dest]["state"] == flipFlopStates["on"]:
                    # send hight pulse
                    for output in m[dest]["outputs"]:
                        pulses.append((dest, output, pulseTypes["high"]))
                elif m[dest]["state"] == flipFlopStates["off"]:
                    # send low pulse
                    for output in m[dest]["outputs"]:
                        pulses.append((dest, output, pulseTypes["low"]))

        elif m[dest]["type"] == moduleTypes["conjunction"]:
            m[dest]["inputs"][src] = pType
            allHigh = True
            for i in m[dest]["inputs"]:
                if m[dest]["inputs"][i] == pulseTypes["low"]:
                    allHigh = False
                    break
            
            if allHigh:
                # send low pulse
                for output in m[dest]["outputs"]:
                    pulses.append((dest, output, pulseTypes["low"]))
            else:
                # send high pulse
                for output in m[dest]["outputs"]:
                    pulses.append((dest, output, pulseTypes["high"]))

    return [highPulseCount, lowPulseCount, m]


def broadcast2(modules, cycleCount, buttonPressCount):
    m = deepcopy(modules)
    pulses = []

    for output in m["broadcaster"]["outputs"]:
        pulses.append(("broadcast", output, pulseTypes["low"]))


    while True:
        if len(pulses) <= 0:
            break

        src, dest, pType = pulses.pop(0)

        if dest == "rs" and pType == pulseTypes["high"]:
            cycleCount.append(buttonPressCount)

        if dest not in m:
            continue
        elif m[dest]["type"] == moduleTypes["flipFlop"]:
            if pType == pulseTypes["low"]:
                m[dest]["state"] = toggleFlipFlop(m[dest])
                if m[dest]["state"] == flipFlopStates["on"]:
                    # send hight pulse
                    for output in m[dest]["outputs"]:
                        pulses.append((dest, output, pulseTypes["high"]))
                elif m[dest]["state"] == flipFlopStates["off"]:
                    # send low pulse
                    for output in m[dest]["outputs"]:
                        pulses.append((dest, output, pulseTypes["low"]))

        elif m[dest]["type"] == moduleTypes["conjunction"]:
            m[dest]["inputs"][src] = pType
            allHigh = True
            for i in m[dest]["inputs"]:
                if m[dest]["inputs"][i] == pulseTypes["low"]:
                    allHigh = False
                    break
            
            if allHigh:
                # send low pulse
                for output in m[dest]["outputs"]:
                    pulses.append((dest, output, pulseTypes["low"]))
            else:
                # send high pulse
                for output in m[dest]["outputs"]:
                    pulses.append((dest, output, pulseTypes["high"]))

    return [m, cycleCount]


def toggleFlipFlop(flipFlop):
    return flipFlopStates["off"] if flipFlop["state"] == flipFlopStates["on"] else flipFlopStates["on"]


if __name__ == "__main__":
    main()
