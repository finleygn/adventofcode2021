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

func ReadInputPlain(path string) string {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	return fileContent
}

func ReadInputAsInt(path string, delimiter string) []int {
	lines := ReadInput("./calendar/day-01/in.txt", "\n")
	return ListToIntList(lines, 10)
}

func ListToIntList(lines []string, base int) []int {
	var numbers []int
	for _, line := range lines {
		number, err := strconv.ParseInt(line, base, 0)

		if err != nil {
			panic("Invalid input")
		}

		numbers = append(numbers, int(number))
	}

	return numbers
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
