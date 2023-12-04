package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1(games []Game, limit map[string]int) {
	s := 0
	for i, game := range games {
		for _, o := range game {
			for k, v := range o {
				if v > limit[k] {
					goto next
				}
			}
		}
		s += i + 1
	next:
	}
	fmt.Printf("%d\n", s)
}

func part2(games []Game) {
	s := 0

	for _, game := range games {
		max := map[string]int{}
		for _, o := range game {
			for k, v := range o {
				if v > max[k] {
					max[k] = v
				}
			}
		}

		power := 1
		for _, v := range max {
			power *= v
		}
		s += power
	}
	fmt.Printf("%d\n", s)
}

type Outcome map[string]int
type Game []Outcome

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage: <filename>")
		os.Exit(-1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	defer file.Close()

	games := []Game{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		a := strings.Split(l, ": ")
		b := strings.Split(a[1], "; ")

		game := Game{}
		for _, g := range b {
			res := strings.Split(g, ", ")

			o := Outcome{}
			for _, r := range res {
				var c string
				var n int
				fmt.Sscanf(r, "%d %s", &n, &c)
				o[c] = n
			}
			game = append(game, o)
		}
		games = append(games, game)
	}

	part1(games, map[string]int{"red": 12, "green": 13, "blue": 14})
	part2(games)
}
