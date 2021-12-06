package main

import (
	"aoc/util/input"
)

// bitwise moment

func get_frequency(data []int, bit_index int) [2]int {
	count := [2]int{0, 0}
	for i := range data {
		bit := data[i] & (1 << bit_index)
		count[bit>>bit_index] += 1
	}
	return count
}

func most_frequent(count [2]int) int {
	if count[0] >= count[1] {
		return 0
	}
	return 1
}

func has_most_frequent(count [2]int) bool {
	return count[0] != count[1]
}

func solve1(binary_length int, data []int) int {
	gamma := 0
	for bit_index := 0; bit_index < binary_length; bit_index++ {
		gamma |= (most_frequent(get_frequency(data, bit_index)) & 1) << bit_index
	}
	return gamma * (^gamma & (1<<binary_length - 1))
}

func get_rating(binary_length int, data []int, use_least_frequent int, keep int) int {
	values := data

	for bit_index := binary_length - 1; bit_index >= 0; bit_index-- {
		new_values := []int{}
		count := get_frequency(values, bit_index)
		rating_bit := most_frequent(count) ^ use_least_frequent

		if !has_most_frequent(count) {
			rating_bit = keep
		}

		for _, v := range values {
			if v>>bit_index&1 == rating_bit {
				new_values = append(new_values, v)
			}
		}

		values = new_values
		if len(values) == 1 {
			break
		}
	}

	return values[0]
}

func solve2(binary_length int, data []int) int {
	oxygen_rating := get_rating(binary_length, data, 0, 1)
	co2_rating := get_rating(binary_length, data, 1, 0)

	return oxygen_rating * co2_rating
}

func main() {
	lines := input.ReadInput("./calendar/day-03/in.txt", "\n")

	result1 := solve1(
		len(lines[0]),
		input.ListToIntList(lines, 2),
	)

	println(result1)

	result2 := solve2(
		len(lines[0]),
		input.ListToIntList(lines, 2),
	)

	println(result2)
}
