
openingChars = "[{<("
scoring = {')': 1, ']': 2, '}': 3, '>': 4}
pairs = {'}': '{', ']': '[', '>': '<', ')': '('}
inversePairs = {'{': '}', '[': ']', '<': '>', '(': ')'}

def openFileToLines(filename):
    lines = []
    with open("2021/10/"+filename) as file:
        for line in file:
            lines.append(list(line.rstrip()))
    return lines

def isValidPair(leftRune, rightrune):
    return leftRune == pairs[rightrune]

def peek(stack):
    return stack[len(stack)-1]

def isLineValid(line, workingStack):
    for rune in line:
        if rune in openingChars:
            workingStack.append(rune)
        else:
            if isValidPair(peek(workingStack), rune):
                workingStack.pop()
            else:
                return False
    return True

def part2(lines):
    scores = []
    for line in lines:
        workingStack = []
        if(isLineValid(line, workingStack)):
            lineScore = 0
            while len(workingStack) > 0:
                lineScore *= 5
                lineScore += scoring[inversePairs[workingStack.pop()]]
            scores.append(lineScore)
    scores.sort()
    return scores[round(len(scores)/2)]

print("Part 2: "+str(part2(openFileToLines("input.txt"))))

