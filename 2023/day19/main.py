#!/usr/bin/python3
import os
import time
from copy import deepcopy

def main():
    data = None
    with open("example", "r") as file:
        data = file.read()

    workflowStr, partStr = data.split(os.linesep + os.linesep)

    workflowList = workflowStr.split(os.linesep)
    workflowMap = {}
    for workflow in workflowList:
        name, ruleset = workflow.split("{")
        ruleset = ruleset.replace("}", "")
        rulesetList = ruleset.split(",")
        workflowMap[name] = rulesetList

    partList = partStr.split(os.linesep)
    for i, part in enumerate(partList):
        part = part.replace("{", "").replace("}", "")
        partData = part.split(",")
        partMap = {}
        for data in partData:
            n, v = data.split("=")
            partMap[n] = int(v)

        partList[i] = partMap


    start_time = time.time()
    print("Part One: ", partOne(workflowMap, partList))
    print("--- %s seconds ---" % (time.time() - start_time))

    start_time = time.time()
    print("Part Two: ", partTwo(workflowMap))
    print("--- %s seconds ---" % (time.time() - start_time))

    return 0


def partOne(workflowMap, partList):
    acceptedPartsXmasValueSum = 0
    for part in partList:
        partCategory = getPartCategory(workflowMap, part, "in")
        if partCategory == "A":
            for v in part.values():
                acceptedPartsXmasValueSum += v
            
    return acceptedPartsXmasValueSum


def partTwo(workflowMap):
    acceptedPaths = getAcceptedPaths(workflowMap, "in", [], [])

    aps = []
    for ap in acceptedPaths:
        maxs = {
            "x": 4000,
            "m": 4000,
            "a": 4000,
            "s": 4000
        }
        mins = {
            "x": 0,
            "m": 0,
            "a": 0,
            "s": 0
        }
        for r in ap:
            if ":" in r:
                """ Explaination for off by 1:
                while subtracting min is exclusive and max is inclusive in math
                so to make min inclusive we subtract 1
                and to make max exclusive we subtract 1
                """
                if ">=" in r:
                    mins[r[0]] = max(mins[r[0]], int(r.split(":")[0].split(">=")[1]) - 1)
                elif "<=" in r:
                    maxs[r[0]] = min(maxs[r[0]], int(r.split(":")[0].split("<=")[1]))
                elif ">" in r:
                    mins[r[0]] = max(mins[r[0]], int(r.split(":")[0].split(">")[1]))
                elif "<" in r:
                    maxs[r[0]] = min(maxs[r[0]], int(r.split(":")[0].split("<")[1]) - 1)

        p = {}
        for k in maxs.keys():
            p[k] = maxs[k] - mins[k]

        aps.append(p)

    t = 0
    for ps in aps:
        m = 1
        for v in ps.values():
            m *= v
        t += m
    
    return t


def getAcceptedPaths(wfm, wf, path, paths):
    for i, rule in enumerate(wfm[wf]):
        npath = deepcopy(path)

        for j in range(i):
            if wfm[wf][j][1] == ">":
                npath.append(wfm[wf][j].replace(">", "<="))
            else:
                npath.append(wfm[wf][j].replace("<", ">="))

        npath.append((rule))

        nextWf = getRuleNextWf(rule)
        if nextWf == "A":
            paths.append(npath)
            continue
        elif nextWf == "R":
            continue

        getAcceptedPaths(wfm, getRuleNextWf(rule), npath, paths)

    return paths

def getRuleNextWf(rule):
    if ":" in rule:
        return rule.split(":")[1]
    return rule


def getPartCategory(workflowMap, part, entryPoint): # Returns A or R
    ep = entryPoint
    while True:
        if ep in "AR":
            break

        for rule in workflowMap[ep]:
            # no condition, always applies
            if ":" not in rule:
                ep = rule
                break

            isMatch = matchWorkflowCondition(rule.split(":")[0], part)
            if isMatch:
                ep = rule.split(":")[1]
                break

    return ep


def matchWorkflowCondition(condition, part):
    return eval(condition.replace(condition[0], str(part[condition[0]])))


if __name__ == "__main__":
    main()
