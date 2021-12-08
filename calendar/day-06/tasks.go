package main

import (
	"aoc/util/input"
	"aoc/util/number"
	"fmt"
)

type School struct {
	offset int
	ages   []int
}

func (s *School) get_age_index(index int) int {
	return number.Mod(index-s.offset, len(s.ages))
}

func (s *School) get_age(index int) int {
	return s.ages[s.get_age_index(index)]
}

func (s *School) set_age(index int, value int) {
	s.ages[s.get_age_index(index)] = value
}

func (s *School) step() {
	to_move := s.get_age(0)
	s.offset -= 1
	s.set_age(6, s.get_age(6)+to_move)
}

func (s *School) get_total() int {
	total := 0
	for i := 0; i < len(s.ages); i++ {
		total += s.ages[i]
	}
	return total
}

func get_ages_from_fish(fish []int, max int) []int {
	ages := []int{}

	for i := 0; i <= max; i++ {
		ages = append(ages, 0)
	}

	for _, age := range fish {
		ages[age] += 1
	}

	return ages
}

func main() {
	fish := input.ReadInputAsInt("./calendar/day-06/in.txt", ",")

	school := School{
		offset: 0,
		ages:   get_ages_from_fish(fish, 8),
	}

	for i := 0; i < 256; i++ {
		if i == 80 {
			fmt.Println(school.get_total())
		}
		school.step()
	}

	fmt.Println(school.get_total())
}
