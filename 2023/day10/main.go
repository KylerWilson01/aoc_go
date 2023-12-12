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

type direction uint

type offset struct {
	i, j int
}

const (
	stop direction = iota
	north
	east
	south
	west
)

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
	o := findStart(l)

	maxPath, maxSet := walkMaze(o, l)

	if part == 1 {
		ans = (len(maxPath) + 1) / 2
	} else {
		l[o.i][o.j] = findStartType(maxPath)
		ans = count(l, maxSet)
	}

	fmt.Printf("Part %d: %d\n", part, ans)
}

func findStart(l [][]byte) offset {
	for i, line := range l {
		for j, b := range line {
			if b == 'S' {
				return offset{i, j}
			}
		}
	}
	return offset{}
}

func next(c byte, d direction) direction {
	switch c {
	case '|':
		return [5]direction{0, north, 0, south, 0}[d]
	case '-':
		return [5]direction{0, 0, east, 0, west}[d]
	case 'L':
		return [5]direction{0, 0, 0, east, north}[d]
	case 'J':
		return [5]direction{0, 0, north, west, 0}[d]
	case '7':
		return [5]direction{0, west, south, 0, 0}[d]
	case 'F':
		return [5]direction{0, east, 0, 0, south}[d]
	default:
		return stop
	}
}

func step(d direction, o offset) offset {
	switch d {
	case north:
		return offset{o.i - 1, o.j}
	case east:
		return offset{o.i, o.j + 1}
	case south:
		return offset{o.i + 1, o.j}
	case west:
		return offset{o.i, o.j - 1}
	default:
		return o
	}
}

func walkMaze(o offset, l [][]byte) ([]offset, map[offset]struct{}) {
	var maxPath []offset
	var maxSet map[offset]struct{}
	var maxLength int

	for i := north; i < west; i++ {
		idx := step(i, o)
		d := i

		var path []offset
		seen := make(map[offset]struct{})
		for {
			if idx.i < 0 || idx.j < 0 || idx.i >= len(l) || idx.j >= len(l[idx.i]) {
				return path, nil
			}
			if _, ok := seen[idx]; ok {
				break
			}

			seen[idx] = struct{}{}
			path = append(path, idx)

			c := l[idx.i][idx.j]
			nextDir := next(c, d)
			if nextDir == stop {
				break
			}

			d = nextDir
			idx = step(nextDir, idx)
		}

		if len(path) > maxLength {
			maxLength = len(path)
			maxPath = path
			maxSet = seen
		}
	}

	return maxPath, maxSet
}

func findStartType(path []offset) byte {
	first := path[0]
	last := path[len(path)-2] // the final node is the start node, so the second to last node is actually last

	k := offset{first.i - last.i, first.j - last.j}
	switch k {
	case offset{2, 0}, offset{-2, 0}:
		return '|'
	case offset{0, 2}, offset{0, -2}:
		return '-'
	case offset{-1, -1}:
		return 'L'
	case offset{-1, 1}:
		return 'J'
	case offset{1, 1}:
		return '7'
	case offset{-1, 1}:
		return 'F'
	}
	return 0
}

func count(l [][]byte, loop map[offset]struct{}) int {
	var count int

	for i := 1; i < len(l)-1; i++ {
		var in bool
		var lastVertex byte

		for j := 0; j < len(l[i])-1; j++ {
			if _, ok := loop[offset{i, j}]; ok {
				switch l[i][j] {
				case '|':
					in = !in
					lastVertex = 0
				case 'F', 'L':
					lastVertex = l[i][j]
				case 'J':
					if lastVertex == 'F' {
						in = !in
					}
					lastVertex = 0
				case '7':
					if lastVertex == 'L' {
						in = !in
					}
					lastVertex = 0

				}
			} else {
				lastVertex = 0
				if in {
					count++
				}
			}
		}

	}

	return count
}

func parseInput(i string) (l [][]byte) {
	return bytes.Fields([]byte(i))
}
