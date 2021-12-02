package main

import (
	"aoc/util/input"
	"fmt"
)

func sum_window(lines []int, start int, end int) int {
	window, total := lines[start:end], 0
	for _, line := range window {
		total += line
	}
	return total
}

func scan(lines []int, window_size int) int {
	total, previous := 0, 0
	for i := 0; i < len(lines); i++ {
		window := sum_window(lines, i, i+window_size)
		if window > previous && i != 0 {
			total += 1
		}
		previous = window
	}
	return total
}

func solve1(lines []int) int {
	return scan(lines, 1)
}

func solve2(lines []int) int {
	return scan(lines, 3)
}

func main() {
	lines := input.ReadInputAsInt("./calendar/day-01/in.txt", "\n")

	result := solve1(lines)
	fmt.Println(result)

	result2 := solve2(lines)
	fmt.Println(result2)
}
