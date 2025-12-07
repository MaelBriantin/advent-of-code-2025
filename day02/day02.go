package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE_PATH string = "./day02/input.txt"

var invalidIds = []string{}

func Part1() error {
	sum, err := readFromInputFile()
	if err != nil {
		return err
	}
	fmt.Printf("Part 1 - Final answer: %d ⭐\n", sum)
	return nil
}

func readFromInputFile() (int, error) {
	content, err := os.ReadFile(INPUT_FILE_PATH)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}

	for idRange := range strings.SplitSeq((string(content)), ",") {
		ids := strings.Split(idRange, "-")
		firstIdToInt, err := strconv.Atoi(ids[0])
		if err != nil {
			return 0, fmt.Errorf("failed to convert first id to int: %w", err)
		}
		secondIdToInt, err := strconv.Atoi(ids[1])
		if err != nil {
			return 0, fmt.Errorf("failed to convert second id to int: %w", err)
		}
		for i := firstIdToInt; i <= secondIdToInt; i++ {
			checkIdValidity(strconv.Itoa(i))
		}
	}
	return sumOfInvalidIds(), nil
}

func checkIdValidity(id string) bool {
	if id[0] == 0 {
		return false
	}
	if len(id) % 2 != 0 {
		return true
	}
	idPartOne := id[:len(id)/2]
	idPartTwo := id[len(id)/2:]
	if idPartOne == idPartTwo {
		invalidIds = append(invalidIds, id)
	}
	return true
}

func sumOfInvalidIds() int {
	var total int
	for i, id := range invalidIds {
		stringToInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("failed to convert id %s at index %d to int: %v\n", id, i, err)
			continue
		}
		total += stringToInt
	}
	return total
}
