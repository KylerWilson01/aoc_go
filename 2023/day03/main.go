package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Part struct {
	number     int
	startIndex int
	endIndex   int
}

var symbols = []string{",", "(", ")", "[", "]", "{", "}", "<", ">", "!", "#", "$", "%", "^", "&", "*", "?", ":", ";", "=", "+", "-", "@"}

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

func solve_part_1(part []Part) (ans int) {
	return
}

func solve_part_2(part []Part) (ans int) {
	return
}

func parse_input(input string) (parts []Part) {
	input = strings.ReplaceAll(input, ".", " ")
	for _, symbol := range symbols {
		input = strings.ReplaceAll(input, symbol, ".")
	}

	println(input)

	return []Part{{
		number:     0,
		startIndex: 0,
		endIndex:   0,
	}}
}
