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

func ReadFileAndParseToIntArray(fileName string) []int {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var in []int

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		in = append(in, Atoi(line))
	}

	return in
}

func main() {
	input := ReadFileAndParseToIntArray("input.txt")

	log.Printf("Part 1: %v\n", PartOne(input))
	log.Printf("Part 2: %v\n", PartTwo(input))

}

func countIncreases(in []int) int {
	var result int

	for i := 1; i < len(in); i++ {
		//fmt.Print(in[i])
		if in[i] > in[i-1] {
			//	fmt.Print(" Increased\n")
			result++
			continue
		}
		//fmt.Print(" Decreased\n")
	}
	return result
}

func PartOne(in []int) int {
	return countIncreases(in)
}

// Part 2
func PartTwo(in []int) int {
	var windowSums []int
	for i := 0; i < len(in)-2; i++ {
		x, y, z := in[i], in[i+1], in[i+2]
		sum := x + y + z
		// fmt.Println(x, y, z, sum)
		windowSums = append(windowSums, sum)
	}
	return countIncreases(windowSums)
}
