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
	parts, symbols := parse_input(input)

	if part == 1 {
		ans = solve_part_1(parts)
	} else {
		ans = solve_part_2(symbols)
	}

	fmt.Printf("Part %d: %d\n", part, ans)
}

func solve_part_1(parts []Part) (ans int) {
	for _, part := range parts {
		if part.legal {
			ans += part.number
		}
	}
	return
}

func solve_part_2(symbols []Symbol) (ans int) {
	for _, s := range symbols {
		if s.symbol == "*" && len(s.parts) == 2 {
			ans += (s.parts[0].number * s.parts[1].number)
			continue
		}
	}
	return
}

func parse_input(input string) (parts []Part, symbols []Symbol) {
	input = strings.TrimRight(input, "\n")

	lines := strings.Split(input, "\n")
	ll := len(lines)

	partNumberIdxs := make([][][]int, ll)
	partIndecatorIdxs := make([][][]int, ll)

	rpn, err := regexp.Compile("([0-9]+)")
	if err != nil {
		panic(err)
	}

	rpi, err := regexp.Compile("([*+#$/=@%&-])")
	if err != nil {
		panic(err)
	}

	for i, line := range lines {
		partNumberMatches := rpn.FindAllStringSubmatchIndex(line, -1)
		partIndecatorMatches := rpi.FindAllStringSubmatchIndex(line, -1)

		partNumberIdxs[i] = partNumberMatches
		partIndecatorIdxs[i] = partIndecatorMatches
	}

	for i := 0; i < ll; i++ {
		for j := 0; j < len(partIndecatorIdxs[i]); j++ {
			if len(partIndecatorIdxs[i][j]) <= 0 {
				continue
			}
			// go to that symbol index
			s := &Symbol{
				symbol: lines[i][partIndecatorIdxs[i][j][0]:partIndecatorIdxs[i][j][1]],
				line:   i,
				idx:    partIndecatorIdxs[i][j][0],
			}

			s.findNumsAroundSymbols(partNumberIdxs[i-1], lines[i-1])
			s.findNumsAroundSymbols(partNumberIdxs[i], lines[i])
			s.findNumsAroundSymbols(partNumberIdxs[i+1], lines[i+1])

			// we want to check to see if there's numbers touching it
			symbols = append(symbols, *s)
			// if there is, then we want to add it to the part slice in the symbol object
		}
		for j := 0; j < len(partNumberIdxs[i]); j++ {
			if len(partNumberIdxs[i][j]) <= 0 {
				continue
			}

			p := Part{
				number:     0,
				startIndex: partNumberIdxs[i][j][0],
				endIndex:   partNumberIdxs[i][j][1],
				legal:      false,
			}

			if i != 0 {
				findIndecatorIdx(partIndecatorIdxs[i-1], lines[i-1], i-1, &p)
			}
			findIndecatorIdx(partIndecatorIdxs[i], lines[i], i, &p)
			if i < ll-1 {
				findIndecatorIdx(partIndecatorIdxs[i+1], lines[i+1], i+1, &p)
			}

			n, err := strconv.Atoi(lines[i][p.startIndex:p.endIndex])
			if err != nil {
				panic(err)
			}

			p.number = n

			if p.legal {
				parts = append(parts, p)
			}
		}
	}

	return
}

func (s *Symbol) findNumsAroundSymbols(nums [][]int, line string) {
	//build part
	var parts []Part
	for _, v := range nums {
		if (v[0] >= s.idx-1 && v[0] <= s.idx+1) || (v[1]-1 >= s.idx-1 && v[1]-1 <= s.idx+1) {
			part := Part{
				startIndex: v[0],
				endIndex:   v[1],
				number:     0,
				legal:      true,
			}

			num, err := strconv.Atoi(line[part.startIndex:part.endIndex])

			if err != nil {
				panic(err)
			}

			part.number = num

			parts = append(parts, part)
		}
	}

	s.parts = append(s.parts, parts...)
}

func findIndecatorIdx(inds [][]int, line string, idx int, part *Part) {
	if part.legal {
		return
	}
	for _, ind := range inds {
		if ind[0] >= part.startIndex-1 && ind[1] <= part.endIndex+1 {
			part.legal = true
		}
	}
}
