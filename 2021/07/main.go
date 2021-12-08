package main

import (
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := ReadFileAndParseToIntArray("input.txt")
	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}

func ReadFileAndParseToIntArray(fileName string) []int {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var in []int

	lines := strings.Split(string(file), ",")
	for _, line := range lines {
		in = append(in, Atoi(line))
	}

	return in
}

func PartOne(input []int) int {
	median := FindMedian(input)
	//log.Print(median)
	var fuel int
	for _, position := range input {
		fuel += int(math.Abs(float64(position - median)))
	}
	return fuel
}

func FindMedian(input []int) int {
	sort.Ints(input)
	medianIndex := len(input) / 2
	if len(input)%2 == 1 {
		return input[medianIndex]
	} else {
		return (input[medianIndex] + input[medianIndex+1]) / 2
	}
}

func PartTwo(input []int) int {
	var minimumFuel int = 9223372036854775807 //int max
	sort.Ints(input)
	for i := input[0]; i < input[len(input)-1]; i++ {
		var fuel int
		for _, position := range input {
			fuel += calculateTriangularFuelCost(position, i)
		}
		if fuel < minimumFuel {
			minimumFuel = fuel
		}
	}
	return minimumFuel
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(s))
	return i
}

func calculateTriangularFuelCost(position int, destination int) int {
	return CalculateTriangularCost(int(math.Abs(float64(position - destination))))
}

func CalculateTriangularCost(n int) int {
	return (n * (n + 1)) / 2
}
