package main

import (
	"io/ioutil"
	"log"
	"strconv"
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
		in = append(in, strings.Split(strings.TrimSpace(line), ""))
	}

	return in
}

func PartOne(input [][]string) int {
	myMapSlice := make([]map[string]int, len(input[0]))
	for i := 0; i < len(input[0]); i++ {
		myMapSlice[i] = make(map[string]int)
	}

	for _, binaryString := range input {
		for i := 0; i < len(binaryString); i++ {
			myMapSlice[i][binaryString[i]] = myMapSlice[i][binaryString[i]] + 1
		}
	}

	var highFrequencies string
	var lowFrequencies string

	for i := 0; i < len(input[0]); i++ {
		//log.Printf("1: %d, 0: %d\n", myMapSlice[i]["1"], myMapSlice[i]["0"])
		if myMapSlice[i]["1"] > myMapSlice[i]["0"] {
			highFrequencies += "1"
			lowFrequencies += "0"
		} else {
			highFrequencies += "0"
			lowFrequencies += "1"
		}
	}

	log.Println(highFrequencies)
	log.Println(lowFrequencies)
	gammaRate, _ := strconv.ParseInt(highFrequencies, 2, 64)
	epsilonRate, _ := strconv.ParseInt(lowFrequencies, 2, 64)
	log.Println(gammaRate)
	log.Println(epsilonRate)
	return int(gammaRate) * int(epsilonRate)
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(s))
	return i
}

type M map[int]interface{}
