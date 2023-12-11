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

type Report struct {
	numbers    []int
	nextNumber int
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
	for _, report := range r {
		tempNums := [][]int{}
		tempNums = append(tempNums, report.numbers)

		foundZeros := false
		for foundZeros == false {
			tt := []int{}
			tna := tempNums[len(tempNums)-1]
			fz := true

			for i := 1; i < len(tna); i++ {
				tpn := tna[i-1]
				tcn := tna[i]
				tdn := tcn - tpn
				tt = append(tt, tdn)

				if tdn != 0 {
					fz = false
				}
			}

			foundZeros = fz
			tempNums = append(tempNums, tt)
		}

		for i := len(tempNums) - 1; i > 0; i-- {
			ltn := tempNums[i][len(tempNums[i])-1]
			lstn := tempNums[i-1][len(tempNums[i-1])-1]

			tempNums[i-1] = append(tempNums[i-1], ltn+lstn)

			if i == 1 {
				a += (ltn + lstn)
			}
		}
	}

	return
}

func solvePart2(r []Report) (a int) {
	for _, report := range r {
		tempNums := [][]int{}
		tempNums = append(tempNums, report.numbers)

		foundZeros := false
		for foundZeros == false {
			tt := []int{}
			tna := tempNums[len(tempNums)-1]
			fz := true

			for i := 1; i < len(tna); i++ {
				tpn := tna[i-1]
				tcn := tna[i]
				tdn := tcn - tpn
				tt = append(tt, tdn)

				if tdn != 0 {
					fz = false
				}
			}

			foundZeros = fz
			tempNums = append(tempNums, tt)
		}

		for i := len(tempNums) - 1; i > 0; i-- {
			ltn := tempNums[i][0]
			lstn := tempNums[i-1][0]

			tempNums[i-1] = append([]int{lstn - ltn}, tempNums[i-1]...)

			if i == 1 {
				a += (lstn - ltn)
			}
		}
	}
	return
}

func parseInput(i string) (l []Report) {
	i = strings.TrimRight(i, "\n")

	for _, line := range strings.Split(i, "\n") {
		nums := strings.Split(line, " ")
		r := Report{
			numbers:    []int{},
			nextNumber: 0,
		}

		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			r.numbers = append(r.numbers, n)
		}

		l = append(l, r)
	}

	return
}
