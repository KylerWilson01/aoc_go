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

type Time int
type Distance int

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
	t, d := parse_input(input)

	if part == 1 {
		ans = solve_part_1(t, d)
	} else {
		ans = solve_part_2(t, d)
	}

	fmt.Printf("Part %d: %d\n", part, ans)
}

func solve_part_1(t []Time, d []Distance) (a int) {
	if len(t) != len(d) {
		panic(fmt.Sprintf("len(t) != len(d): %d != %d", len(t), len(d)))
	}

	a = 1

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < len(t); i++ {
		wg.Add(1)

		go func(i int) {
			mmpms := 0
			ttr := t[i]

			tbd := []int{}

			for tt := 0; tt < int(t[i]); tt++ {
				ttr--
				mmpms++

				if int(d[i]) < (mmpms * int(ttr)) {
					tbd = append(tbd, mmpms)
				}
			}

			mu.Lock()
			a *= len(tbd)
			mu.Unlock()

			wg.Done()
		}(i)
	}

	wg.Wait()

	return
}

func solve_part_2(t []Time, d []Distance) (a int) {
	if len(t) != len(d) {
		panic(fmt.Sprintf("len(t)!= len(d): %d!= %d", len(t), len(d)))
	}

	// take all the times in t, convert them to strings, and concatenate them
	tsa := make([]string, len(t))
	for i := 0; i < len(t); i++ {
		tsa[i] = strconv.Itoa(int(t[i]))
	}
	ts := strings.Join(tsa, "")

	dsa := make([]string, len(t))
	for i := 0; i < len(t); i++ {
		dsa[i] = strconv.Itoa(int(d[i]))
	}
	ds := strings.Join(dsa, "")

	dis, err := strconv.Atoi(ds)

	if err != nil {
		panic(err)
	}

	mmpms := 0
	ttr, err := strconv.Atoi(ts)
	if err != nil {
		panic(err)
	}

	fmt.Println(ttr, dis)

	tbd := 0

	l := ttr

	for tt := 0; tt < l; tt++ {
		ttr--
		mmpms++

		if dis < (mmpms * ttr) {
			tbd++
		}
	}

	a = tbd

	return
}

func parse_input(i string) (t []Time, d []Distance) {
	i = strings.TrimRight(i, "\n")

	for _, line := range strings.Split(i, "\n") {
		parts := strings.Split(line, ": ")
		name := parts[0]
		nums := parts[1]

		for _, n := range strings.Split(nums, " ") {
			if n == "" {
				continue
			}

			x, err := strconv.Atoi(strings.TrimSpace(n))
			if err != nil {
				panic(err)
			}

			if name == "Time" {
				t = append(t, Time(x))
			} else if name == "Distance" {
				d = append(d, Distance(x))
			}
		}
	}

	return
}
