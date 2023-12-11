package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func part1(lines []string, factor int) {
	s := 0

	emptyRows := map[int]bool{}
	emptyCols := map[int]bool{}
	for i, l := range lines {
		if strings.Contains(l, "#") {
			continue
		}
		emptyRows[i] = true
	}
	for i := range lines[0] {
		l := ""
		for j := range lines {
			l += string(lines[j][i])
		}
		if strings.Contains(l, "#") {
			continue
		}
		emptyCols[i] = true
	}

	stars := map[int]Coord{}
	n := 0
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '#' {
				nx := 0
				ny := 0
				for k := 0; k < i; k++ {
					if emptyRows[k] {
						ny++
					}
				}
				for k := 0; k < j; k++ {
					if emptyCols[k] {
						nx++
					}
				}
				stars[n] = Coord{i + ny*(factor-1), j + nx*(factor-1)}
				n++
			}
		}
	}

	for i := range stars {
		for j := i + 1; j < len(stars); j++ {
			s += abs(stars[i].x-stars[j].x) + abs(stars[i].y-stars[j].y)
		}
	}

	fmt.Printf("%d\n", s)
}

func part2(lines []string) {
	part1(lines, 1000000)
}

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

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	part1(lines, 2)
	part2(lines)
}
