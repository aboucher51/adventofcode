package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const lanternfish_adolescence int = 2
const lanternfish_cycle int = 6
const lanternfish_cohort_length = lanternfish_adolescence + lanternfish_cycle + 1 + 1 //newborn fish are marked an extra day so they dont iterate twice

const numberOfDaysToIteratePartOne int = 80
const numberOfDaysToIteratePartTwo int = 256

func main() {
	input := ReadFileAndParseToIntArray("input.txt")
	log.Printf("Part 1: %d", SimulateLanternFish(input, numberOfDaysToIteratePartOne))
	log.Printf("Part 2: %d", SimulateLanternFish(input, numberOfDaysToIteratePartTwo))
}

func SimulateLanternFish(input []int, numberOfDaysToIterate int) int {
	//initialize cohorts. a cohort is a group of functionally identical lanternfish
	var cohorts []int = make([]int, lanternfish_cohort_length)
	for _, fishNumber := range input {
		cohorts[fishNumber]++
	}

	for i := 0; i < numberOfDaysToIterate; i++ {
		// if i < 10 {
		// 	log.Println(cohorts)
		// }
		cohorts = IterateCohorts(cohorts)
		// if i < 10 {
		// 	log.Println(cohorts)
		// 	log.Println()
		// }
	}

	var arraySum int
	for _, fishNumber := range cohorts {
		arraySum += fishNumber
	}

	return arraySum
}

func IterateCohorts(cohorts []int) []int {

	var newCohorts []int = make([]int, len(cohorts))

	cohorts[lanternfish_cycle+lanternfish_adolescence+1] += cohorts[0]
	cohorts[lanternfish_cycle+1] += cohorts[0]

	for i := 0; i < len(cohorts)-1; i++ {
		newCohorts[i] = cohorts[i+1]
	}

	return newCohorts
}

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

	lines := strings.Split(string(file), ",")
	for _, line := range lines {
		in = append(in, Atoi(line))
	}

	return in
}
