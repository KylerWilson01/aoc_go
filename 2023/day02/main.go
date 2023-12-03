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

type Cube struct {
	color  string
	amount int
}

type Game struct {
	id    string
	cubes [][]Cube
}

const (
	RED   = 12
	GREEN = 13
	BLUE  = 14
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
	lines := parse_input(input)

	if part == 1 {
		ans = solve_part_1(lines)
	} else {
		ans = solve_part_2(lines)
	}

	fmt.Printf("Part %d: %d\n", part, ans)
}

func solve_part_2(games []Game) (ans int) {
	g := make(map[int]int)
	for i, game := range games {
		c := make(map[string]int)
		g[i] = 1
		for _, cubes := range game.cubes {
			for _, cube := range cubes {
				if cube.amount > c[cube.color] {
					c[cube.color] = cube.amount
				}
			}
		}
		for _, amounts := range c {
			g[i] = g[i] * amounts
		}
	}
	for _, amounts := range g {
		ans += amounts
	}

	return
}

func solve_part_1(games []Game) (ans int) {
	for _, game := range games {
		legal_game := true
		for _, cubes := range game.cubes {
			if !legal_game {
				break
			}
			for _, cube := range cubes {
				switch cube.color {
				case "red":
					if cube.amount > RED {
						legal_game = false
						break
					}
				case "green":
					if cube.amount > GREEN {
						legal_game = false
						break
					}
				case "blue":
					if cube.amount > BLUE {
						legal_game = false
						break
					}
				default:
					panic("unknown color")
				}
			}
		}
		if !legal_game {
			continue
		}
		id, err := strconv.Atoi(game.id)

		if err != nil {
			panic(err)
		}

		ans += id
	}

	return
}

func parse_cubes(input []string) (cubes [][]Cube) {
	cubes = make([][]Cube, len(input))
	for i, line := range input {
		line = strings.TrimSpace(line)
		cube_info := strings.Split(line, ",")
		cubes[i] = make([]Cube, len(cube_info))

		for j, c := range cube_info {
			c = strings.TrimSpace(c)
			info := strings.Split(string(c), " ")
			amount, err := strconv.Atoi(strings.TrimSpace(info[0]))

			if err != nil {
				panic(err)
			}

			cubes[i][j] = Cube{
				color:  info[1],
				amount: amount,
			}
		}
	}

	return cubes
}

func parse_input(input string) (games []Game) {
	input = strings.TrimRight(input, "\n")
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line = strings.Replace(line, "Game ", "", 1)
		info := strings.Split(line, ":")
		game := Game{
			id:    strings.TrimSpace(info[0]),
			cubes: parse_cubes(strings.Split(strings.TrimSpace(info[1]), ";")),
		}
		games = append(games, game)
	}

	return games
}
