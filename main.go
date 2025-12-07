package main

import (
	"fmt"
	"os"

	"advent-of-code/day01"
	"advent-of-code/day02"
	"advent-of-code/day03"
)

func main() {
	fmt.Println("=== Advent of Code ===")

	fmt.Println("\n--- Day 01 ---")
	if err := day01.Part1(); err != nil {
		fmt.Fprintf(os.Stderr, "Day 01 Part 1 error: %v\n", err)
	}

	if err := day01.Part2(); err != nil {
		fmt.Fprintf(os.Stderr, "Day 01 Part 2 error: %v\n", err)
	}

	fmt.Println("\n--- Day 02 ---")
	day02.Part1()
	day02.Part2()

	fmt.Println("\n--- Day 03 ---")
	day03.Part1()
	day03.Part2()
}
