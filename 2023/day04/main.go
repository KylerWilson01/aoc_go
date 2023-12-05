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

type ScratchCard struct {
	cardNumber     int
	winningNumbers map[int]int
	numbers        []int
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

func solve_part_1(input []ScratchCard) (ans int) {
	for _, card := range input {
		points := 0
		for _, n := range card.numbers {
			if card.winningNumbers[n] > 0 {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		ans += points
	}

	return
}

func solve_part_2(input []ScratchCard) (ans int) {
	cardMap := make(map[int]int)
	for _, card := range input {
		cardMap[card.cardNumber]++
	}

	for _, card := range input {
		wins := 0
		for _, n := range card.numbers {
			if card.winningNumbers[n] > 0 {
				wins++
			}
		}

		for j := 0; j < cardMap[card.cardNumber]; j++ {
			for i := card.cardNumber; i < card.cardNumber+wins; i++ {
				cardMap[i+1]++
			}
		}
	}
	fmt.Println(cardMap)

	for _, amount := range cardMap {
		ans += amount
	}

	return
}

func parse_input(input string) (cards []ScratchCard) {
	input = strings.TrimRight(input, "\n")

	regex := regexp.MustCompile(`([0-9]+)`)

	for _, line := range strings.Split(input, "\n") {
		part := strings.Split(line, ":")
		nums := strings.Split(part[1], "|")

		cardNum, err := strconv.Atoi(regex.FindAllStringSubmatch(part[0], 1)[0][0])

		if err != nil {
			panic(err)
		}

		card := ScratchCard{
			cardNumber:     cardNum,
			winningNumbers: make(map[int]int),
			numbers:        []int{},
		}

		for i, num := range nums {
			n := regex.FindAllStringSubmatch(num, -1)
			for _, x := range n {
				pn, err := strconv.Atoi(x[0])
				if err != nil {
					panic(err)
				}
				if i == 0 {
					card.winningNumbers[pn]++
				} else {
					card.numbers = append(card.numbers, pn)
				}
			}
		}

		cards = append(cards, card)
	}

	return
}
