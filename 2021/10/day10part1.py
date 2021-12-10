import os
print (os.getcwd())

heightmap = []
with open("2021/09/input.txt") as file:
    for line in file:
        heightmap.append(list(line.rstrip()))

localMinima = []

for i in range(len(heightmap)):
    for j in range(len(heightmap[i])):
        currentPoint = heightmap[i][j]
        if(    (i==0 or (heightmap[i-1][j] > currentPoint))
           and (j==0 or (heightmap[i][j-1] > currentPoint))
           and (i==len(heightmap)-1 or (heightmap[i+1][j] > currentPoint))
           and (j==len(heightmap[i])-1 or (heightmap[i][j+1] > currentPoint))):
            localMinima.append(currentPoint)

#print(localMinima)

localMinimaSum = 0
for localMinimum in localMinima:
    localMinimaSum += int(localMinimum)+1

print("Part 1: "+str(localMinimaSum))
