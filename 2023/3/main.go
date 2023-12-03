package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/exp/maps"
)

func isAdjacent(i, j int, lines []string) bool {
	for row := i - 1; row <= i+1; row++ {
		for col := j - 1; col <= j+1; col++ {
			if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[i]) {
				continue
			}
			if lines[row][col] != '.' && (lines[row][col] < '0' || lines[row][col] > '9') {
				return true
			}
		}
	}
	return false
}

func part1(lines []string) {
	s := 0

	for i, l := range lines {
		re := regexp.MustCompile(`\d+`) // numbers
		matches := re.FindAllStringIndex(l, -1)
		for _, m := range matches {
			for j := m[0]; j < m[len(m)-1]; j++ {
				if isAdjacent(i, j, lines) {
					n := l[m[0]:m[len(m)-1]]
					v, _ := strconv.Atoi(n)
					s += v
					break
				}
			}
		}
	}

	fmt.Printf("%d\n", s)
}

func getAdjacentGears(i, j int, lines []string) map[string]struct{} {
	gears := map[string]struct{}{}
	for row := i - 1; row <= i+1; row++ {
		for col := j - 1; col <= j+1; col++ {
			if row < 0 || row >= len(lines) || col < 0 || col >= len(lines[i]) {
				continue
			}
			if lines[row][col] == '*' {
				gears[fmt.Sprintf("%d:%d", row, col)] = struct{}{}
			}
		}
	}
	return gears
}

func part2(lines []string) {
	s := 0
	gears := map[string]map[int]struct{}{} // gear -> set of numbers

	for i, l := range lines {
		re := regexp.MustCompile(`\d+`) // numbers
		matches := re.FindAllStringIndex(l, -1)
		for _, m := range matches {
			for j := m[0]; j < m[len(m)-1]; j++ {
				adjGears := getAdjacentGears(i, j, lines)

				for g, _ := range adjGears {
					n := l[m[0]:m[len(m)-1]]
					v, _ := strconv.Atoi(n)
					if gears[g] == nil {
						gears[g] = map[int]struct{}{}
					}
					gears[g][v] = struct{}{}
				}
			}
		}
	}

	for _, nums := range gears {
		vals := maps.Keys(nums)
		if len(vals) != 2 {
			continue
		}
		s += vals[0] * vals[1]
	}

	fmt.Printf("%d\n", s)
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

	part1(lines)
	part2(lines)
}
