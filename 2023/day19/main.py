#!/usr/bin/python3
import os
import time


def main():
    data = None
    with open("input", "r") as file:
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

    # start_time = time.time()
    # print("Part Two: ", partTwo(workflowMap))
    # print("--- %s seconds ---" % (time.time() - start_time))

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
    return


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
