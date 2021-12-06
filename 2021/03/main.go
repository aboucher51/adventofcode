package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input := ReadFileAndParseToIntArray("input.txt")
	log.Printf("Part 1: %d", PartOne(input))
	//log.Printf("Part 2: %d", SimulateLanternFish(input, numberOfDaysToIteratePartTwo))
}

func ReadFileAndParseToIntArray(fileName string) [][]string {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var in [][]string

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		in = append(in, strings.Split(line, ""))
	}

	return in
}

func PartOne(input [][]string) int {
	var myMapSlice []M

	for _, binaryString := range input {
		for i := 0; i < len(binaryString); i++ {
			myMapSlice[i][binaryString[i]] = myMapSlice[i][binaryString[i]] + 1
		}
	}

	return 134324
}

type M map[int]interface{}
