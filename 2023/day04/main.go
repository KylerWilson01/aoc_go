package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Symbol struct {
	symbol string
	line   int
	idx    int
	parts  []Part
}

type Part struct {
	number     int
	startIndex int
	endIndex   int
	legal      bool
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part, ans int
	flag.IntVar(&part, "part", 1, "part 1 or part 2")
	flag.Parse()
	lines := parse_input(input)

	if part == 1 {
		ans = solve_part_1(lines)
	} else {
		ans = solve_part_2(lines)
	}

	fmt.Printf("Part %d: %d\n", part, ans)
}

func solve_part_1(input []string) (ans int) {
	return
}

func solve_part_2(input []string) (ans int) {
	return
}

func parse_input(input string) (lines []string) {
	return
}
