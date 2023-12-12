package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Directions string

var DirectionMap = map[Directions]string{
	"|": "NS",
	"-": "EW",
	"L": "NE",
	"J": "NW",
	"7": "SW",
	"F": "SE",
	".": "Ground",
	"S": "Starting",
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
	l := parseInput(input)

	if part == 1 {
		ans = solvePart1(l)
	} else {
		ans = solvePart2(l)
	}

	fmt.Printf("Part %d: %d\n", part, ans)
}

func solvePart1(r []Report) (a int) {
	return
}

func solvePart2(r []Report) (a int) {
	return
}

func parseInput(i string) (l []Report) {
	i = strings.TrimRight(i, "\n")

	for _, line := range strings.Split(i, "\n") {
	}

	return
}
