winningPairs = {
    'A':'Y',
    'B':'Z',
    'C':'X'
}
losingPairs = {
    'A':'Z',
    'B':'X',
    'C':'Y'
}

def openFileToLines(filename):
    lines = []
    with open("2022/02/"+filename) as file:
        for line in file:
            line = line.rstrip()
            lines.append(line.split(' '))
            
    return lines

def part1(lines):
    runningTotal = 0
    for match in lines:
        runningTotal += scoreMatchPart1(match)
    return runningTotal

def scoreMatchPart1(input):
    sum = ord(input[1])-87
    if ord(input[0]) == ord(input[1])-23:
        return sum + 3
    elif input[1] == winningPairs[input[0]]:
        return sum + 6
    else:
        return sum


def part2(lines):
    runningTotal = 0
    for match in lines:
        runningTotal += scoreMatchPart2(match)
    return runningTotal

def scoreMatchPart2(input):
    sum = 0
    if input[1] == 'Y':
        return 3 + ord(input[0])-64
    elif input[1] == 'X':
        return 0 + ord(losingPairs[input[0]])-87
    else:
        return 6 + ord(winningPairs[input[0]])-87

    

print("Part 1: "+str(part1(openFileToLines("input.txt"))))
print("Part 2: "+str(part2(openFileToLines("input.txt"))))


