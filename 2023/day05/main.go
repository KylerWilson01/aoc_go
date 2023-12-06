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
		ans = solve_part_2(seeds, maps)
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

func solve_part_2(s Seeds, m []Map) (ans int) {
	c := make(chan int)
	defer close(c)
	mu := sync.Mutex{}
	var wg sync.WaitGroup

	for i := 0; i < len(s); i += 2 {
		for j := s[i]; j < s[i]+s[i+1]-1; j++ {
			go func(j int, waitingGroup *sync.WaitGroup, mumutex *sync.Mutex) {
				waitingGroup.Add(1)
				destination := j

				for _, map_ := range m {
					for _, range_ := range map_.ranges {
						if range_.sourceStart <= destination && destination < range_.sourceStart+(range_.length) {
							destination = (destination - range_.sourceStart) + range_.destinationStart
							break
						}
					}
				}

				mumutex.Lock()
				if destination < ans || ans == 0 {
					ans = destination
				}
				mumutex.Unlock()
				waitingGroup.Done()
			}(j, &wg, &mu)
		}
	}
	wg.Wait()

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
