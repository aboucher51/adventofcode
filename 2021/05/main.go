package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	coordinatePairs := ReadFileAndParse("input.txt")
	populatedGrid := BuildGrid(coordinatePairs)
	intersectionCount := CountIntersections(populatedGrid)
	log.Print(intersectionCount)
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(s))
	return i
}

func ReadFileAndParse(fileName string) []CoordinatePair {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var aggregatedLines []CoordinatePair

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		aggregatedLines = append(aggregatedLines, parseLineToCoordinatePair(line))
	}

	return aggregatedLines
}

func parseLineToCoordinatePair(line string) CoordinatePair {
	const middle_coordinate_delimiter string = " -> "
	pairs := strings.Split(line, middle_coordinate_delimiter)
	coord1 := strings.Split(pairs[0], ",")
	coord2 := strings.Split(pairs[1], ",")
	//fmt.Printf("%d,%d -> %d,%d\n", Atoi(coord1[0]), Atoi(coord1[1]), Atoi(coord2[0]), Atoi(coord2[1]))
	return CoordinatePair{x1: Atoi(coord1[0]), y1: Atoi(coord1[1]), x2: Atoi(coord2[0]), y2: Atoi(coord2[1])}
}

func BuildGrid(coordinatePairs []CoordinatePair) [999][999]int {
	var grid [999][999]int
	for _, coord := range coordinatePairs {
		if coord.x1 == coord.x2 {
			for i := Min(coord.y1, coord.y2); i <= Max(coord.y1, coord.y2); i++ {
				grid[coord.x1][i]++
			}
		} else if coord.y1 == coord.y2 {
			for i := Min(coord.x1, coord.x2); i <= Max(coord.x1, coord.x2); i++ {
				grid[i][coord.y1]++
			}
		}
	}
	return grid
}

func CountIntersections(grid [999][999]int) int {
	var count int
	for _, row := range grid {
		for _, cell := range row {
			if cell > 1 {
				count++
			}
		}
	}
	return count
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

type CoordinatePair struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

//greater than 8314
//less than 8425
