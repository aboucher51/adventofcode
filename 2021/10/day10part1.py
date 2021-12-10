
lines = []
with open("2021/10/input.txt") as file:
    for line in file:
        lines.append(list(line.rstrip()))

openingChars = "[{<("
pairs = {'}': '{', ']': '[', '>': '<', ')': '('}
scoring = {')': 3, ']': 57, '}': 1197, '>': 25137}

workingStack = []
score = 0

for line in lines:
    for rune in line:
        if rune in openingChars:
            workingStack.append(rune)
        else:
            if workingStack[len(workingStack)-1] == pairs[rune]:
                workingStack.pop()
            else:
                score += scoring[rune]
                break

print("Part 1: "+str(score))
