package input

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInput(path string, delimiter string) []string {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	slicedContent := strings.Split(fileContent, delimiter)

	return slicedContent
}

func ReadInputAsInt(path string, delimiter string) []int {
	lines := ReadInput("./calendar/day-01/in.txt", "\n")

	var numbers []int
	for _, line := range lines {
		number, err := strconv.Atoi(line)

		if err != nil {
			panic("Invalid input")
		}

		numbers = append(numbers, number) // note the = instead of :=
	}

	return numbers
}
