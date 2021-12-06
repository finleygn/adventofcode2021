package old

import (
	"aoc/util/input"
)

// bitwise moment

type iter func(int)

func masked_iter(length int, callback iter) {
	for col := 0; col < length; col++ {
		mask := 1 << col
		callback(mask)
	}
}

func solve1(bin_length int, data []int) int {
	gamma := 0

	masked_iter(bin_length, func(mask int) {
		total := 0

		for row := 0; row < len(data); row++ {
			total += (mask & data[row])
		}

		if total > len(data)*mask/2 {
			gamma |= mask
		}
	})

	return gamma * (^gamma & (1<<bin_length - 1))
}

func main() {
	lines := input.ReadInput("./calendar/day-03/in.txt", "\n")
	data := input.ListToIntList(lines, 2)

	result1 := solve1(
		len(lines[0]),
		data,
	)

	println(result1)
}
