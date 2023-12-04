package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Part struct {
	number     int
	startIndex int
	endIndex   int
	truePart   bool
	nextTo     string
	symbolLine int
	smybolIdx  int
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

func solve_part_1(parts []Part) (ans int) {
	for _, part := range parts {
		if part.truePart {
			ans += part.number
		}
	}
	return
}

func solve_part_2(parts []Part) (ans int) {
	return
}

func parse_input(input string) (parts []Part) {
	input = strings.ReplaceAll(strings.TrimRight(input, "\n"), ".", " ")

	lines := strings.Split(input, "\n")
	length := len(lines)
	partNumberIdxs := make([][][]int, length)
	partIndecatorIdxs := make([][][]int, length)

	for i, line := range lines {
		regexPartNumber, err := regexp.Compile("([0-9]+)")
		if err != nil {
			panic(err)
		}
		partNumberMatches := regexPartNumber.FindAllStringSubmatchIndex(line, -1)

		regexPartIndecators, err := regexp.Compile("([*+#$/=@%&-])")
		if err != nil {
			panic(err)
		}
		partIndecatorMatches := regexPartIndecators.FindAllStringSubmatchIndex(line, -1)

		partNumberIdxs[i] = partNumberMatches
		partIndecatorIdxs[i] = partIndecatorMatches
	}

	for i := 0; i < length; i++ {

		for j := 0; j < len(partNumberIdxs[i]); j++ {
			if len(partNumberIdxs[i][j]) <= 0 {
				continue
			}
			// check if there's an indecator near it
			part := Part{
				number:     0,
				startIndex: 0,
				endIndex:   0,
				truePart:   false,
			}

			part.startIndex = partNumberIdxs[i][j][0]
			part.endIndex = partNumberIdxs[i][j][1]
			n, err := strconv.Atoi(lines[i][part.startIndex:part.endIndex])
			if err != nil {
				panic(err)
			}

			if i != 0 {
				for _, inds := range partIndecatorIdxs[i-1] {
					if inds[0] >= part.startIndex-1 && inds[1] <= part.endIndex+1 {
						part.truePart = true
						part.nextTo = lines[i-1][inds[0]:inds[1]]
						part.symbolLine = i - 1
						part.smybolIdx = inds[0]
					}
				}
			}
			for _, inds := range partIndecatorIdxs[i] {
				if inds[0] >= part.startIndex-1 && inds[1] <= part.endIndex+1 {
					part.truePart = true
					part.nextTo = lines[i][inds[0]:inds[1]]
					part.symbolLine = i
					part.smybolIdx = inds[0]
				}
			}
			if i < length-1 {
				for _, inds := range partIndecatorIdxs[i+1] {
					if inds[0] >= part.startIndex-1 && inds[1] <= part.endIndex+1 {
						part.truePart = true
						part.nextTo = lines[i+1][inds[0]:inds[1]]
						part.symbolLine = i + 1
						part.smybolIdx = inds[0]
					}
				}
			}

			part.number = n
			parts = append(parts, part)
		}
	}

	fmt.Println(parts)

	return
}
