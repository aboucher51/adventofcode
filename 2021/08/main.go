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
	log.Printf("Part 2: %d", PartTwo(input))
}

func ReadFileAndParseToIntArray(fileName string) []SignalEntry {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var in []SignalEntry

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		row := strings.Split(strings.TrimSpace(line), "|")
		pattern := strings.Split(strings.TrimSpace(row[0]), " ")
		output := strings.Split(strings.TrimSpace(row[1]), " ")
		in = append(in, SignalEntry{signalPattern: pattern, signalOutput: output})
	}

	return in
}

//count instances of output lengths 2, 4, 3, and 7 (corresponding to 1, 4, 7, and 8 respectively)
func PartOne(input []SignalEntry) int {
	uniqueLengths := []int{2, 4, 3, 7}
	var identifiableUniqueInputCount int
	for _, entry := range input {
		for _, outputEntry := range entry.signalOutput {
			if SliceContains(uniqueLengths, len(outputEntry)) {
				identifiableUniqueInputCount++
			}
		}
	}
	return identifiableUniqueInputCount
}

func PartTwo(input []SignalEntry) int {
	var outputSum int

	for _, sequence := range input {
		numericallyTranslatedSlice := DeduceNumberAllocationAndReturnTranslatedPattern(sequence)
		outputSlice := numericallyTranslatedSlice[len(numericallyTranslatedSlice)-4:]

		var outputNumberString string
		for _, digit := range outputSlice {
			outputNumberString += strconv.Itoa(digit)
		}
		//fmt.Println(outputNumberString)
		outputSum += Atoi(outputNumberString)
	}

	return outputSum
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(s))
	return i
}

func SliceContains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func StringSliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func StringContainsAllCharsUnordered(source string, testRunes string) bool {
	for _, char := range testRunes {
		if !strings.Contains(source, string(char)) {
			return false
		}
	}
	return true
}

type SignalEntry struct {
	signalPattern []string
	signalOutput  []string
}

func DeduceNumberAllocationAndReturnTranslatedPattern(input SignalEntry) []int {
	inputEntry := append(input.signalPattern, input.signalOutput...)
	translatedEntry := make([]int, 14) //int representation, defaults vals to 0

	var rightWires string
	var topWire string
	var foursUniqueWires string //in a standard display, this is digits bd
	var bottomWire string
	var middleWire string
	var topRightWire string
	var bottomRightWire string

	//go through and mark all unique, known numbers. All should be present in every row
	for i, sequence := range inputEntry { // 1 case, give us rows on right
		if len(sequence) == 2 {
			translatedEntry[i] = 1
			rightWires = sequence
		}
	}

	for i, sequence := range inputEntry { // 7 case. should give us top row
		if len(sequence) == 3 {
			translatedEntry[i] = 7
			if topWire == "" {
				for _, char := range sequence {
					if !strings.Contains(rightWires, string(char)) {
						topWire = string(char)
					}
				}
			}
		}
	}

	for i, sequence := range inputEntry { // 4 case, gives us unique digits
		if len(sequence) == 4 {
			translatedEntry[i] = 4
			if foursUniqueWires == "" {
				for _, char := range sequence {
					if !strings.Contains(rightWires, string(char)) {
						foursUniqueWires += string(char)
					}
				}
			}
		}
	}

	for i, sequence := range inputEntry { // 8 case. no new info, just 8's
		if len(sequence) == 7 {
			translatedEntry[i] = 8
		}
	}

	for i, sequence := range inputEntry { // 9 case, should give us bottom row
		if len(sequence) == 6 && StringContainsAllCharsUnordered(sequence, rightWires+foursUniqueWires) { //neither 6 nor 0 contain a perfect 4 shape.
			translatedEntry[i] = 9

			if bottomWire == "" {
				for _, char := range sequence {
					if !strings.Contains(rightWires+foursUniqueWires+topWire, string(char)) {
						bottomWire = string(char)
					}
				}
			}
		}
	}

	for i, sequence := range inputEntry { // 3 case
		if len(sequence) == 5 && StringContainsAllCharsUnordered(sequence, rightWires) {
			translatedEntry[i] = 3

			if middleWire == "" {
				for _, char := range sequence {
					if !strings.Contains(rightWires+topWire+bottomWire, string(char)) {
						middleWire = string(char)
					}
				}
			}
		}
	}

	for i, sequence := range inputEntry { // 6 case
		if len(sequence) == 6 && !StringContainsAllCharsUnordered(sequence, rightWires) {
			translatedEntry[i] = 6

			if topRightWire == "" {
				for _, char := range rightWires {
					if !strings.Contains(sequence, string(char)) {
						topRightWire = string(char)
					}
				}
				for _, char := range rightWires {
					if string(char) != topRightWire {
						bottomRightWire = string(char)
					}
				}
			}
		}
	}

	for i, sequence := range inputEntry { // 5 case
		if len(sequence) == 5 && !strings.Contains(sequence, topRightWire) {
			translatedEntry[i] = 5
		}
	}

	for i, sequence := range inputEntry { // 2 case
		if len(sequence) == 5 && !strings.Contains(sequence, bottomRightWire) {
			translatedEntry[i] = 2
		}
	}

	//fmt.Println(translatedEntry)
	return translatedEntry
}
