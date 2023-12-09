package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Node string

type Pattern string

type Document struct {
	node  Node
	left  Node
	right Node
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
	p, mnd := parse_input(input)

	if part == 1 {
		ans = solvePart1(p, mnd)
	} else {
		ans = solve_part_2(p, mnd)
	}

	fmt.Printf("Part %d: %d\n", part, ans)
}

func solvePart1(p Pattern, mnd map[Node]Document) (a int) {
	cn := mnd[Node("AAA")]
	endNode := mnd[Node("ZZZ")]
	foundEnd := false
	pIdx := 0

	for foundEnd == false {
		if cn.node == endNode.node {
			foundEnd = true
			break
		}

		if p[pIdx] == 'R' {
			cn = mnd[cn.right]
		} else {
			cn = mnd[cn.left]
		}

		if pIdx == len(p)-1 {
			pIdx = 0
		} else {
			pIdx++
		}

		a++
	}

	return
}

func solve_part_2(p Pattern, mnd map[Node]Document) (a int) {
	cn := []Document{}
	steps := []int{}

	for _, doc := range mnd {
		if doc.node[2] == 'A' {
			cn = append(cn, doc)
			steps = append(steps, 0)
		}
	}

	for i, doc := range cn {
		foundEnd := false
		pIdx := 0
		current := doc
		s := steps[i]
		for foundEnd == false {
			if current.node[2] == 'Z' {
				foundEnd = true
				steps[i] = s
				break
			}

			if p[pIdx] == 'R' {
				current = mnd[current.right]
			} else {
				current = mnd[current.left]
			}

			if pIdx == len(p)-1 {
				pIdx = 0
			} else {
				pIdx++
			}
			s++
		}
	}

	return LCM(1, 1, steps...)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func parse_input(i string) (p Pattern, mnd map[Node]Document) {
	i = strings.TrimRight(i, "\n")

	pd := strings.Split(i, "\n\n")
	mnd = make(map[Node]Document)

	p = Pattern(pd[0])

	for _, line := range strings.Split(pd[1], "\n") {
		parts := strings.Split(line, " = ")

		docs := strings.Split(strings.TrimSpace(strings.Map(func(r rune) rune {
			if r == '(' || r == ')' {
				return ' '
			}
			return r
		}, parts[1])), ", ")

		mnd[Node(parts[0])] = Document{
			node:  Node(parts[0]),
			left:  Node(docs[0]),
			right: Node(docs[1]),
		}
	}

	return
}
