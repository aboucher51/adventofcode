def openFileToLines(filename):
    lines = []
    with open("2022/03/"+filename) as file:
        for line in file:
            line = line.rstrip()
            lines.append(line.split(' '))
            
    return lines

def part1(lines):
    sum = 0
    for line in lines:
        sum+=scorePocket(line[0])
    return sum

def scorePocket(line):
    firstpart, secondpart = line[:len(line)//2], line[len(line)//2:]
    commonChar = list(set(firstpart)&set(secondpart))[0]
    if commonChar.islower():
        return ord(commonChar)-96
    else:
        return ord(commonChar)-38
    
# def part2(lines):
#     sum = 0
#     for line in lines:
#         sum+=scorePocket(line[0])
#     return sum
    

print("Part 1: "+str(part1(openFileToLines("input.txt"))))
#print("Part 2: "+str(part2(openFileToLines("input.txt"))))


