package main

import (
	"aoc/util/input"
	"fmt"
	"strconv"
	"strings"
)

type Move struct {
	direction string
	magnitude int
}

func parseMoves(lines []string) []Move {
	moves := make([]Move, len(lines))

	for _, line := range lines {
		v := strings.Split(line, " ")
		magnitude, err := strconv.Atoi(v[1])

		if err != nil {
			panic("Invalid magnitude")
		}

		moves = append(moves, Move{
			direction: v[0],
			magnitude: magnitude,
		})
	}

	return moves
}

type Helicopter struct {
	x   int
	y   int
	aim int
}

func (h *Helicopter) move(move Move) {
	switch move.direction {
	case "forward":
		h.x += move.magnitude
		h.y += move.magnitude * h.aim
	case "up":
		h.aim -= move.magnitude
	case "down":
		h.aim += move.magnitude
	}
}

func (h *Helicopter) navigate_linear() int {
	// Aim acts like y would in linear navigation
	return h.aim * h.x
}

func (h *Helicopter) navigate_aim() int {
	return h.y * h.x
}

func main() {
	lines := input.ReadInput("./calendar/day-02/in.txt", "\n")
	moves := parseMoves(lines)

	helicopter := Helicopter{x: 0, y: 0, aim: 0}

	for _, m := range moves {
		helicopter.move(m)
	}

	// Solution 1
	fmt.Println(helicopter.navigate_linear())

	// Solution 2
	fmt.Println(helicopter.navigate_aim())
}
