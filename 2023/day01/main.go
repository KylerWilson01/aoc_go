package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or part 2")
	flag.Parse()

	if part == 1 {
		ans := solvePart1(input)
		fmt.Printf("Part 1: %d\n", ans)
	} else {
		ans := solvePart2(input)
		fmt.Printf("Part 2: %d\n", ans)
	}
}

func solvePart1(input string) (ans int) {
	lines := parseInput(input)

	for i := 0; i < len(lines); i++ {
		var first, last int
		line := lines[i]
		for j := 0; j < len(line); j++ {
			front_value, err1 := strconv.Atoi(string(line[j]))
			back_value, err2 := strconv.Atoi(string(line[len(line)-j-1]))

			if first == 0 && err1 == nil {
				first = front_value
			}

			if last == 0 && err2 == nil {
				last = back_value
			}

			if first != 0 && last != 0 {
				break
			}
		}

		x, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))

		if err != nil {
			panic("could not convert to int")
		}

		ans += x
	}

	return ans
}

func solvePart2(input string) (ans int) {
	lines := parseInput(input)
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i := 0; i < len(lines); i++ {
		var first, last int
		line := lines[i]
		for i, v := range numbers {
			line = strings.ReplaceAll(line, v, v+strconv.Itoa(i+1)+v)
		}

		for j := 0; j < len(line); j++ {
			front_value, err1 := strconv.Atoi(string(line[j]))
			back_value, err2 := strconv.Atoi(string(line[len(line)-j-1]))

			if first == 0 && err1 == nil {
				first = front_value
			}

			if last == 0 && err2 == nil {
				last = back_value
			}

			if first != 0 && last != 0 {
				break
			}
		}

		x, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))

		if err != nil {
			panic("could not convert to int")
		}

		ans += x
	}

	return ans
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}
