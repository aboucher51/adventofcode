

def openFileToLines(filename):
    lines = []
    with open("2022/01/"+filename) as file:
        for line in file:
            line = line.rstrip()
            if line.isnumeric():
                lines.append(int(line))
            else:
                lines.append(0)
            
    return lines

def sumInput(lines):
    elves = []
    runningTotal = 0
    for number in lines:
        runningTotal += number
        if number == 0:
            elves.append(runningTotal)
            runningTotal = 0
    return elves

def part1(lines):
    elves = sumInput(lines)
    elves.sort(reverse=True)
    return elves[0]

def part2(lines):
    elves = sumInput(lines)
    elves.sort(reverse=True)
    return elves[0]+elves[1]+elves[2]


print("Part 1: "+str(part1(openFileToLines("input.txt"))))
print("Part 2: "+str(part2(openFileToLines("input.txt"))))


