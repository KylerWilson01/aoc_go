package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

//go:embed input.txt
var input string

type Map struct {
	typeOfMap string
	ranges    []Range
}

type Range struct {
	destinationStart int
	sourceStart      int
	length           int
}

type Seeds []int

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
	seeds, maps := parse_input(input)

	if part == 1 {
		ans = solve_part_1(seeds, maps)
	} else {
		a := solve_part_2(seeds, maps)
		ans = a.ans
	}

	fmt.Printf("Part %d: %d\n", part, ans)
}

func solve_part_1(s Seeds, m []Map) (ans int) {
	for _, seed := range s {
		destination := seed

		for _, map_ := range m {
			for _, range_ := range map_.ranges {
				if range_.sourceStart <= destination && destination < range_.sourceStart+(range_.length) {
					destination = (destination - range_.sourceStart) + range_.destinationStart
					break
				}
			}
		}

		if destination < ans || ans == 0 {
			ans = destination
		}
	}

	return
}

type Answer struct {
	ans int
	mu  sync.Mutex
	wg  sync.WaitGroup
}

func solve_part_2(s Seeds, m []Map) (ans Answer) {
	var lowests []int
	for i := 0; i < len(s); i += 2 {
		ans.wg.Add(1)
		go func(i int) {
			var lowest int
			for j := s[i]; j < s[i]+s[i+1]-1; j++ {
				destination := j

				for _, map_ := range m {
					for _, range_ := range map_.ranges {
						if range_.sourceStart <= destination && destination < range_.sourceStart+(range_.length) {
							destination = (destination - range_.sourceStart) + range_.destinationStart
							break
						}
					}
				}

				if destination < lowest || lowest == 0 {
					lowest = destination
				}
			}
			ans.mu.Lock()
			lowests = append(lowests, lowest)
			ans.mu.Unlock()
			ans.wg.Done()
		}(i)
	}
	ans.wg.Wait()

	for _, lowest := range lowests {
		if lowest < ans.ans || ans.ans == 0 {
			ans.ans = lowest
		}
	}

	return
}

func parse_input(input string) (seeds Seeds, maps []Map) {
	input = strings.TrimRight(input, "\n")

	for i, group := range strings.Split(input, "\n\n") {
		if i == 0 {
			s := strings.Split(strings.Split(group, ": ")[1], " ")

			for _, v := range s {
				n, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, n)
			}

			continue
		}

		var m Map

		for j, line := range strings.Split(group, "\n") {
			if j == 0 {
				m.typeOfMap = strings.Split(line, " map:")[0]
				continue
			}

			var r Range

			n := strings.Split(line, " ")

			var err error

			r.destinationStart, err = strconv.Atoi(n[0])
			if err != nil {
				panic(err)
			}

			r.sourceStart, err = strconv.Atoi(n[1])
			if err != nil {
				panic(err)
			}

			r.length, err = strconv.Atoi(n[2])
			if err != nil {
				panic(err)
			}
			m.ranges = append(m.ranges, r)
		}

		maps = append(maps, m)
	}

	return
}
