package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part2(grid []string, x, y, dir int) {
	s := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '.' {
				continue
			}

			row := []rune(grid[i])
			row[j] = '#'
			grid[i] = string(row)

			if search(grid, x, y, dir, map[string]bool{}, true) == -1 {
				s++
			}

			row[j] = '.'
			grid[i] = string(row)
		}
	}
	fmt.Printf("%d\n", s)
}

func search(grid []string, x, y, dir int, visited map[string]bool, part2 bool) int {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return len(visited)
	}
	s := fmt.Sprintf("%d-%d", x, y)
	if part2 {
		s = fmt.Sprintf("%d-%d-%d", x, y, dir)
	}
	if part2 && visited[s] {
		return -1
	}
	if grid[y][x] == '#' {
		if dir == 0 {
			y += 1
		} else if dir == 1 {
			x -= 1
		} else if dir == 2 {
			y -= 1
		} else {
			x += 1
		}
		dir += 1
		if dir > 3 {
			dir = 0
		}
	}
	visited[s] = true

	if dir == 0 {
		return search(grid, x, y-1, dir, visited, part2)
	} else if dir == 1 {
		return search(grid, x+1, y, dir, visited, part2)
	} else if dir == 2 {
		return search(grid, x, y+1, dir, visited, part2)
	}
	return search(grid, x-1, y, dir, visited, part2)
}

func part1(grid []string, x, y, dir int) {
	s := search(grid, x, y, dir, map[string]bool{}, false)
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

	grid := []string{}
	dir, x, y := 0, 0, 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		grid = append(grid, l)
		if i := strings.Index(l, "^"); i != -1 {
			x, y = i, len(grid)-1
		}
	}

	part1(grid, x, y, dir)
	file.Seek(0, 0)
	part2(grid, x, y, dir)
}
