package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type offset struct {
	i, j int
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

func solvePart1(r [][]byte) (a int) {
	// Since the input is a rectangle, we can just subtract x1,y1 from x2,y2
	return
}

func solvePart2(r [][]byte) (a int) {
	return
}

func parseInput(i string) (l [][]byte) {
	i = strings.TrimRight(i, "\n")

	emptyRows := make(map[offset]struct{})
	emptyColumns := make(map[offset]struct{})

	lines := bytes.Fields([]byte(i))

	for i := range lines {
		emptyRow := true
		emptyColumn := true
		for j := range lines {
			if lines[i][j] == '#' {
				emptyRow = false
			}
			if lines[j][i] == '#' {
				emptyColumn = false
			}

			if !emptyColumn && !emptyRow {
				break
			}
		}

		if emptyRow {
			emptyRows[offset{i, 0}] = struct{}{}
		}

		if emptyColumn {
			emptyColumns[offset{0, i}] = struct{}{}
		}
	}

	// expand the universe

	return
}
