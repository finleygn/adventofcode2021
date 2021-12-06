package main

import (
	"aoc/util/array"
	"aoc/util/input"
	"strconv"
	"strings"
)

type Board struct {
	items []int
	size  int
}

type Game struct {
	boards  []Board
	numbers []int
}

func ParseBoard(input string) Board {
	items := []int{}
	rows := strings.Split(input, "\n")
	size := len(rows)

	for _, row := range rows {
		for _, item := range strings.Fields(row) {
			v, err := strconv.Atoi(item)
			if err != nil {
				panic("Unable to convert number")
			}
			items = append(items, v)
		}
	}

	return Board{
		items,
		size,
	}
}

func ParseGame(input string) Game {
	sections := strings.Split(input, "\n\n")
	boards := []Board{}
	numbers := []int{}

	split := strings.Split(sections[0], ",")

	for _, number := range split {
		v, err := strconv.Atoi(number)
		if err != nil {
			panic("Unable to convert number")
		}
		numbers = append(numbers, v)
	}

	for _, board := range sections[1:] {
		boards = append(boards, ParseBoard(board))
	}

	return Game{
		boards:  boards,
		numbers: numbers,
	}
}

func (b *Board) has_won_row(row int, numbers []int) bool {
	match := true
	for col := 0; col < b.size; col++ {
		if !array.Contains(numbers, b.items[b.size*row+col]) {
			return false
		}
	}
	return match
}

func (b *Board) has_won_col(col int, numbers []int) bool {
	match := true
	for row := 0; row < b.size; row++ {
		if !array.Contains(numbers, b.items[b.size*row+col]) {
			return false
		}
	}
	return match
}

func (b *Board) has_won(numbers []int) bool {
	for index := 0; index < b.size; index++ {
		if b.has_won_col(index, numbers) || b.has_won_row(index, numbers) {
			return true
		}
	}
	return false
}

func (b *Board) sum(numbers []int) int {
	total := 0
	for _, item := range numbers {
		if array.Contains(b.items, item) {
			total += item
		}
	}
	return total
}

func (g *Game) solve1() int {
	for i := 1; i < len(g.numbers); i++ {
		for _, board := range g.boards {
			if board.has_won(g.numbers[:i]) {
				return board.sum(g.numbers[i:]) * g.numbers[i-1]
			}
		}
	}
	return 0
}

func (g *Game) solve2() int {
	won := []int{}
	for i := 1; i < len(g.numbers); i++ {
		for bi, board := range g.boards {
			if array.Contains(won, bi) {
				continue
			}

			if board.has_won(g.numbers[:i]) {
				won = append(won, bi)

				if len(won) == len(g.boards) {
					return board.sum(g.numbers[i:]) * g.numbers[i-1]
				}
			}
		}
	}
	return 0
}

/**
 * Bootstrap
 */
func main() {
	plain := input.ReadInputPlain("./calendar/day-04/in.txt")

	game := ParseGame(plain)
	println(game.solve1())
	println(game.solve2())
}
