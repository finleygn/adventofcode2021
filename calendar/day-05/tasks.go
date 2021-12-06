package main

import (
	"aoc/util/input"
	"aoc/util/number"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Line struct {
	from Position
	to   Position
}

type Grid struct {
	items []Position
}

func (l *Line) is_straight() bool {
	return l.from.x == l.to.x || l.from.y == l.to.y
}

// probably really bad way of doing it perhaps come back
func (l *Line) get_positions() []Position {
	positions := []Position{}
	x_movement, y_movement := l.from.x-l.to.x, l.from.y-l.to.y
	x_difference, y_difference := int(math.Abs(float64(x_movement))), int(math.Abs(float64(y_movement)))

	for i := 0; i <= number.Max(x_difference, y_difference); i++ {
		new_x, new_y := l.from.x, l.from.y

		if x_difference != 0 {
			if l.from.x < l.to.x {
				new_x = l.from.x + i
			} else {
				new_x = l.from.x - i
			}
		}

		if y_difference != 0 {
			if l.from.y < l.to.y {
				new_y = l.from.y + i
			} else {
				new_y = l.from.y - i
			}
		}

		positions = append(positions, Position{
			x: new_x,
			y: new_y,
		})
	}

	return positions
}

func (g *Grid) apply(line Line) {
	positions := line.get_positions()
	g.items = append(g.items, positions...)
}

func (g *Grid) count(min int) int {
	count := 0
	instances := make(map[Position]int)

	for _, position := range g.items {
		instances[position] = instances[position] + 1
	}
	for _, element := range instances {
		if element >= min {
			count += 1
		}
	}

	return count
}

func parse_position(line string) Position {
	values := strings.Split(line, ",")
	position := []int{}

	for _, value := range values {
		v, err := strconv.Atoi(value)
		if err != nil {
			panic("Unable to convert number")
		}
		position = append(position, v)
	}

	return Position{
		x: position[0],
		y: position[1],
	}
}

func parse_lines(lines []string) []Line {
	result := []Line{}

	for _, line := range lines {
		values := strings.Split(line, " -> ")
		result = append(result, Line{
			from: parse_position(values[0]),
			to:   parse_position(values[1]),
		})
	}

	return result
}

func solve1(lines []Line) int {
	grid := Grid{items: []Position{}}

	for _, line := range lines {
		if line.is_straight() {
			grid.apply(line)
		}
	}

	return grid.count(2)
}

func solve2(lines []Line) int {
	grid := Grid{items: []Position{}}

	for _, line := range lines {
		grid.apply(line)
	}

	return grid.count(2)
}

/**
 * Bootstrap
 */
func main() {
	input := input.ReadInput("./calendar/day-05/in.txt", "\n")

	lines := parse_lines(input)

	fmt.Println(solve1(lines))
	fmt.Println(solve2(lines))
}
