package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Hand struct {
	cards    string
	bid      int
	handType int
}

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR        = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

var LETTER_MAP = map[rune]rune{
	'T': 'A',
	'J': '.',
	'Q': 'C',
	'K': 'D',
	'A': 'E',
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
	h := parse_input(input)

	if part == 1 {
		ans = solve_part_1(h)
	} else {
		ans = solve_part_2(h)
	}

	fmt.Printf("Part %d: %d\n", part, ans)
}

func solve_part_1(h []Hand) (a int) {
	for _, hand := range h {
		hand.classify()
	}
	sort.Slice(h, func(i, j int) bool {
		if h[i].handType == h[j].handType {
			return h[i].cards < h[j].cards
		}

		return h[i].handType < h[j].handType
	})

	for i, hand := range h {
		a += ((i + 1) * hand.bid)
	}

	return
}

func (h *Hand) classify() {
	counts := make(map[rune]int)
	for _, c := range h.cards {
		counts[c]++
	}

	for _, c := range counts {
		if c == 5 {
			h.handType = FIVE_OF_A_KIND
			return
		}

		if c == 4 {
			h.handType = FOUR_OF_A_KIND
			return
		}

		if c == 3 {
			if h.handType == ONE_PAIR {
				h.handType = FULL_HOUSE
				return
			}

			h.handType = THREE_OF_A_KIND
		}
		if c == 2 {
			if h.handType == THREE_OF_A_KIND {
				h.handType = FULL_HOUSE
				return
			}
			if h.handType == ONE_PAIR {
				h.handType = TWO_PAIR
				return
			}
			h.handType = ONE_PAIR
		}

		if c == 0 && h.handType < 1 {
			h.handType = HIGH_CARD
		}
	}
}

func (h *Hand) strength() {
	var mappedHand string
	for _, c := range h.cards {
		if LETTER_MAP[c] != 0 {
			mappedHand += string(LETTER_MAP[c])
		} else {
			mappedHand += string(c)
		}
	}

	h.cards = mappedHand
}

func solve_part_2(h []Hand) (a int) {
	sort.Slice(h, func(i, j int) bool {
		if h[i].handType == h[j].handType {
			return h[i].cards < h[j].cards
		}

		return h[i].handType < h[j].handType
	})
	fmt.Println(h)

	for i, hand := range h {
		a += ((i + 1) * hand.bid)
	}

	return
}

func (hand *Hand) handTypePart2() {
	maxType := HIGH_CARD
	for _, char := range []string{"A", "C", "D", "E", "9", "8", "7", "6", "5", "4", "3", "2"} {
		maxType = max(maxType, handType(strings.Replace(hand.cards, ".", char, -1)))
	}
	hand.handType = maxType
}

func handType(h string) int {
	counts := make(map[rune]int)
	for _, c := range h {
		counts[c]++
	}

	firstMax := 0
	var charWithFirstMax rune
	for _, char := range []rune{'A', 'C', 'D', '.', 'E', '9', '8', '7', '6', '5', '4', '3', '2'} {
		if counts[char] > firstMax {
			firstMax = counts[char]
			charWithFirstMax = char
		}
	}
	counts[charWithFirstMax] = 0
	secondMax := 0
	for _, char := range []rune{'A', 'C', 'D', '.', 'E', '9', '8', '7', '6', '5', '4', '3', '2'} {
		if counts[char] > secondMax {
			secondMax = counts[char]
		}
	}
	if firstMax == 5 {
		return FIVE_OF_A_KIND
	}
	if firstMax == 4 {
		return FOUR_OF_A_KIND
	}
	if firstMax == 3 {
		if secondMax == 2 {
			return FULL_HOUSE
		}
		return THREE_OF_A_KIND
	}
	if firstMax == 2 {
		if secondMax == 2 {
			return TWO_PAIR
		}
		return ONE_PAIR
	}
	return HIGH_CARD
}

func parse_input(i string) (h []Hand) {
	i = strings.TrimRight(i, "\n")

	for _, line := range strings.Split(i, "\n") {
		info := strings.Split(line, " ")

		b, err := strconv.Atoi(info[1])

		if err != nil {
			panic(err)
		}

		hand := Hand{
			cards: info[0],
			bid:   b,
		}

		hand.strength()
		hand.handTypePart2()

		h = append(h, hand)
	}

	return
}
