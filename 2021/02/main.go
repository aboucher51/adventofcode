package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(s))
	return i
}

func ReadFileAndParseToIntArray(fileName string) []Vector {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var in []Vector

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		var vectorString []string = strings.Split(strings.TrimSpace(line), " ")
		in = append(in, Vector{direction: VectorDirectionValueOf(vectorString[0]), magnitude: Atoi(vectorString[1])})
	}

	return in
}

func main() {
	input := ReadFileAndParseToIntArray("input.txt")

	log.Printf("Part 1: %v\n", PartOne(input))
	log.Printf("Part 2: %v\n", PartTwo(input))

}

func PartOne(vectors []Vector) int {
	var horizontal int
	var depth int

	for _, vector := range vectors {
		if vector.direction == Forward {
			horizontal += vector.magnitude
		} else if vector.direction == Up {
			depth -= vector.magnitude
		} else {
			depth += vector.magnitude
		}
	}

	return horizontal * depth
}

func PartTwo(vectors []Vector) int {
	var horizontal int
	var depth int
	var aim int

	for _, vector := range vectors {
		if vector.direction == Forward {
			horizontal += vector.magnitude
			depth += aim * vector.magnitude
		} else if vector.direction == Up {
			aim -= vector.magnitude
		} else {
			aim += vector.magnitude
		}
	}

	return horizontal * depth
}

func VectorDirectionValueOf(direction string) VectorDirection {
	switch direction {
	case "up":
		return Up
	case "down":
		return Down
	case "forward":
		return Forward
	}
	return Up
}

type Vector struct {
	direction VectorDirection
	magnitude int
}

type VectorDirection int64

const (
	Forward VectorDirection = iota
	Up
	Down
)
