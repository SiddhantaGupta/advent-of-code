#!/usr/bin/python3

# DP mostly copied from Jonathan Paulson
# solution by Anshuman Dash

import re
	
def solve(line, nums, part1=True):
	if part1 == False:
		line = "?".join([line]*5)
		nums *= 5

	DP = {}

	def f(i, n, b): # return how many solutions there are from this position
		# i - index in line
		# n - index in nums
		# b - size of current block
		if (i,n,b) in DP:return DP[(i,n,b)] # DP already solved it
		
		if i == len(line):	# at the end of the line, return 1 if this is a posible configuration or 0 otherwise
			return int(
				n == len(nums) and b == 0 or 			# no current block, and finished all numbers
				n == len(nums)-1 and b == nums[-1]		# one last block, and currently in a block of that size
			)

		ans = 0
		if line[i] in ".?":	# treat it like a .
			if(b == 0):
				ans += f(i+1, n, 0) # just keep moving forward
			else:			# we have a current block
				if n == len(nums): return 0	# more springs than input asks for, so not a solution
				if b == nums[n]: 			# If we currently have a continguous spring of the required size
					ans += f(i+1, n+1, 0)	# Move forward and count this block
		if line[i] in "#?": # treat it like a #
			ans += f(i+1, n, b+1) 	# no choice but to continue current block
		DP[(i,n,b)] = ans # save to DP
		return ans
	return f(0,0,0)

def parseLine(l):
	lhs,rhs = re.sub(r"\.+",".",l).split()
	nums = eval(rhs)
	return (lhs, nums)

def main(input):
	inps = [parseLine(l) for l in input.splitlines()]
	return sum(solve(*l) for l in inps), sum(solve(*l, part1 = False) for l in inps)

with open("input", "r") as file:
     ans = main(file.read())
     print(ans)

# The below code did not give a proper answer after a lot of tweaking and fixing

# . operational, # broken, ? don't know, number: broken
# with open("example", "r") as file:
#     totalArrangementCount = 0
#     count = 0

#     for line in file:
#         count += 1
#         if count != 2:
#             continue
#         print("count: ", count)
#         data = line.rstrip()
#         springRow = list(data.split(" ")[0])
#         brokenSpringsGroups = data.split(" ")[1].split(",")

#         rowStartIndex = 0
#         rowEndIndex = len(springRow)
#         for i in range(0, len(brokenSpringsGroups)):
#             x = int(brokenSpringsGroups[i])
#             rowEndIndex = rowEndIndex - (x + 1)

#         arrangements = []
#         ac = 0
#         for gs in brokenSpringsGroups:
#             gs = int(gs)
#             ac += 1
#             rowEndIndex = rowEndIndex + gs + 1
#             print("gs: ", gs)
#             print("rowEndIndex: ", rowEndIndex)
#             print("rowStartIndex: ", rowStartIndex)

#             w = springRow[rowStartIndex:rowEndIndex]
#             print("w: ", w)

#             wsi = None
#             wei = len(w)
#             for i in range(0, len(w)):
#                 s = w[i]
#                 print("s: ", s)
#                 if (wsi == None) and (s == "?" or s == "#"):
#                     print("wsi: ", i)
#                     wsi = i
#                 if wsi is not None and s == ".":
#                     wei = i
#                     if wei - wsi >= gs:
#                         break

#             print("wsi: ", wsi)
#             print("wei: ", wei)

#             ww = w[wsi:wei]

#             print("ww" , ww)

#             a = 0
#             colsConsumed = 0
#             if "#" not in ww:
#                 a = len(ww) - gs + 1
#                 colsConsumed = wsi + gs
#             else:
#                 brokenIndex = ww.index("#")
#                 print("brokenIndex: ", brokenIndex)
#                 brokenSpringCount = ww.count("#")
#                 print("brokenSpringCount: ", brokenSpringCount)
#                 print("ww[b:gs+1]: ",ww[brokenIndex:gs+1])
#                 brokenIndexEnd = None
#                 if "#" in ww[brokenIndex:gs+1]:
#                     brokenIndexEnd = len(ww[brokenIndex:gs+1]) - 1 - ww[brokenIndex:gs+1][::-1].index("#")
#                 else:
#                     brokenIndexEnd = 0
#                 brokenPipesCount = brokenIndexEnd + 1
#                 print("brokenPipesCount: ", brokenPipesCount)
#                 print("brokenIndexEnd: ", brokenIndexEnd)
#                 additionalBrokenCount = gs - brokenPipesCount
#                 print("additionalBrokenCount: ", additionalBrokenCount)
#                 idx = brokenIndex - 1 - gs
#                 print("idx: ", idx)

#                 xsi = brokenIndex - additionalBrokenCount
#                 if xsi < 0 :
#                     xsi = 0

#                 xei = brokenIndex + brokenIndexEnd + additionalBrokenCount + 1
#                 if xei > len(ww):
#                     xei = len(ww)

#                 xww = ww[xsi:xei]
#                 print("xsi: ", xsi)
#                 print("xei: ", xei)
#                 print("xww: ", xww)
#                 a = len(xww) - gs + 1
#                 colsConsumed = len(xww) + len(ww[0:xsi])
#                 # colsConsumed = wsi + gs

#             print("a: ", a)
                    
#             rowStartIndex = rowStartIndex + colsConsumed + 1
            
#             arrangements.append(a)

#             print("\n")

#         print("arrangements: ", arrangements)

#         arSum = 0
#         for i in range(0, len(arrangements)):
#             ax = 0
#             ai = arrangements[i]
#             if ai == 1:
#                 continue
#             for j in range(0, len(arrangements)):
#                 aj = arrangements[j]
#                 if aj <= 1:
#                     continue
#                 if j == i:
#                     continue
#                 # print("aj: ", aj)
#                 for k in range(1, aj):
#                     print(f"{ax} += {ai - k}")
#                     ax += ai - k
                
#             # print("ax: ", ax)

#             arSum += ax

#         # print("arSum: ", arSum)
#         if arSum == 0:
#             arSum = 1
#         totalArrangementCount += arSum

#     print(f"totalArrangementCount: ", totalArrangementCount)
