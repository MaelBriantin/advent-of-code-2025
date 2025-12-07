package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE_PATH string = "./day02/input.txt"

type Part int

const (
	PART_ONE Part = iota + 1
	PART_TWO
)


func Part1() error {
	var invalidIds = []string{}
	sum, err := readFromInputFile(&invalidIds, PART_ONE)
	if err != nil {
		return err
	}
	fmt.Printf("Part 1 - Final answer: %d ⭐\n", sum)
	return nil
}

func Part2() error {
	var invalidIds = []string{}
	sum, err := readFromInputFile(&invalidIds, PART_TWO)
	if err != nil {
		return err
	}
	fmt.Printf("Part 2 - Final answer: %d ⭐\n", sum)
	return nil
}

func readFromInputFile(invalidIds *[]string, part Part) (int, error) {
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
			checkIdValidity(invalidIds, strconv.Itoa(i), part)
		}
	}
	return sumOfInvalidIds(invalidIds), nil
}

func checkIdValidity(invalidIds *[]string, id string, part Part) {
	if id[0] == 0 {
		return
	}
	if part == PART_ONE {
		if len(id)%2 != 0 {
			return
		}
		firstHalf := id[:len(id)/2]
		secondHalf := id[len(id)/2:]
		if firstHalf == secondHalf {
			*invalidIds = append(*invalidIds, id)
		}
	}
	if part == PART_TWO {
    for i := 1; i <= len(id)/2; i++ {
        if len(id)%i != 0 {
            continue
        }
        sequence := id[:i]
        repeatedSequence := strings.Repeat(sequence, len(id)/i)
        if repeatedSequence == id {
         	*invalidIds = append(*invalidIds, id)
            return
        }
    }
    return
}
}

func sumOfInvalidIds(invalidIds *[]string) int {
	var total int
	for i, id := range *invalidIds {
		stringToInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("failed to convert id %s at index %d to int: %v\n", id, i, err)
			continue
		}
		total += stringToInt
	}
	return total
}
