package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const INPUT_FILE_PATH string = "./day01/input.txt"

const MAX_DIAL_MARK int = 99
const MIN_DIAL_MARK int = 0
const STARTING_MARK int = 50

func Part1() error {
	hit0 := 0
	actualPosition := STARTING_MARK

	if err := readFromInputFile(&hit0, &actualPosition, 1); err != nil {
		return err
	}
	fmt.Printf("Part 1 - Final answer: %d ⭐\n", hit0)
	return nil
}

func Part2() error {
	hit0 := 0
	actualPosition := STARTING_MARK

	if err := readFromInputFile(&hit0, &actualPosition, 2); err != nil {
		return err
	}
	fmt.Printf("Part 2 - Final answer: %d ⭐\n", hit0)
	return nil
}

func click(code string, hit0 *int, actualPosition *int) {
	clickInternal(code, hit0, actualPosition, false)
}

func clickEachHit(code string, hit0 *int, actualPosition *int) {
	clickInternal(code, hit0, actualPosition, true)
}

func clickInternal(code string, hit0 *int, actualPosition *int, eachHit bool) {
	letter := code[0]
	number, err := strconv.Atoi(code[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	var direction int
	switch letter {
	case 'L':
		direction = -1
	case 'R':
		direction = +1
	default:
		fmt.Fprintf(os.Stderr, "Error: neither L or R direction passed")
		return
	}

	for range number {
		*actualPosition += direction

		if *actualPosition > MAX_DIAL_MARK {
			*actualPosition = MIN_DIAL_MARK
		} else if *actualPosition < MIN_DIAL_MARK {
			*actualPosition = MAX_DIAL_MARK
		}
		if eachHit && *actualPosition == MIN_DIAL_MARK {
			*hit0++
		}
	}

	if !eachHit && *actualPosition == MIN_DIAL_MARK {
		*hit0++
	}
}

func readFromInputFile(hit0 *int, actualPosition *int, part int) error {
	fileHandle, err := os.Open(INPUT_FILE_PATH)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		if closeErr := fileHandle.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to close file: %v\n", closeErr)
		}
	}()

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		code := scanner.Text()
		switch part {
		case 1:
			click(code, hit0, actualPosition)
		case 2:
			clickEachHit(code, hit0, actualPosition)
		default:
			fmt.Fprintf(os.Stderr, "Error: invalid part %d\n", part)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading from file: %w", err)
	}

	return nil
}
