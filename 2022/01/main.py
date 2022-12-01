

def openFileToLines(filename):
    lines = []
    with open("2022/1/"+filename) as file:
        for line in file:
            lines.append(list(line.rstrip()))
    return lines

def part1(lines):
    return ""


print("Part 2: "+str(part1(openFileToLines("input.txt"))))


