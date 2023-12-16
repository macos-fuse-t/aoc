package main

import (
	"bufio"
	"fmt"
	"os"
)

type Beam struct {
	x, y int
	dir  byte
}

func trace(b Beam, grid []string, visited map[Beam]bool) {
	if b.y < 0 || b.y >= len(grid) || b.x < 0 || b.x >= len(grid[0]) {
		return
	}
	if visited[b] {
		return
	}
	visited[b] = true

	switch b.dir {
	case 'N':
		if grid[b.y][b.x] == '.' || grid[b.y][b.x] == '|' {
			b.y--
		} else if grid[b.y][b.x] == '\\' {
			b.x--
			b.dir = 'W'
		} else if grid[b.y][b.x] == '/' {
			b.x++
			b.dir = 'E'
		} else {
			trace(Beam{b.x, b.y, 'W'}, grid, visited)
			b.dir = 'E'
		}
	case 'S':
		if grid[b.y][b.x] == '.' || grid[b.y][b.x] == '|' {
			b.y++
		} else if grid[b.y][b.x] == '\\' {
			b.x++
			b.dir = 'E'
		} else if grid[b.y][b.x] == '/' {
			b.x--
			b.dir = 'W'
		} else {
			trace(Beam{b.x, b.y, 'W'}, grid, visited)
			b.dir = 'E'
		}
	case 'E':
		if grid[b.y][b.x] == '.' || grid[b.y][b.x] == '-' {
			b.x++
		} else if grid[b.y][b.x] == '\\' {
			b.y++
			b.dir = 'S'
		} else if grid[b.y][b.x] == '/' {
			b.y--
			b.dir = 'N'
		} else {
			trace(Beam{b.x, b.y, 'N'}, grid, visited)
			b.dir = 'S'
		}
	case 'W':
		if grid[b.y][b.x] == '.' || grid[b.y][b.x] == '-' {
			b.x--
		} else if grid[b.y][b.x] == '\\' {
			b.y--
			b.dir = 'N'
		} else if grid[b.y][b.x] == '/' {
			b.y++
			b.dir = 'S'
		} else {
			trace(Beam{b.x, b.y, 'N'}, grid, visited)
			b.dir = 'S'
		}
	}
	trace(b, grid, visited)
}

func calc(grid []string, visited map[Beam]bool) int {
	n := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if visited[Beam{x, y, 'N'}] || visited[Beam{x, y, 'S'}] || visited[Beam{x, y, 'W'}] || visited[Beam{x, y, 'E'}] {
				n++
			}
		}
	}
	return n
}

func part1(grid []string) {
	visited := map[Beam]bool{}
	trace(Beam{0, 0, 'E'}, grid, visited)
	s := calc(grid, visited)

	fmt.Printf("%d\n", s)
}

func part2(grid []string) {
	max := 0
	for y := 0; y < len(grid); y++ {
		visited := map[Beam]bool{}
		trace(Beam{0, y, 'E'}, grid, visited)
		s := calc(grid, visited)
		if s > max {
			max = s
		}

		visited = map[Beam]bool{}
		trace(Beam{len(grid[y]) - 1, y, 'W'}, grid, visited)
		s = calc(grid, visited)
		if s > max {
			max = s
		}
	}
	for x := 0; x < len(grid[0]); x++ {
		visited := map[Beam]bool{}
		trace(Beam{x, 0, 'S'}, grid, visited)
		s := calc(grid, visited)
		if s > max {
			max = s
		}

		visited = map[Beam]bool{}
		trace(Beam{x, len(grid) - 1, 'N'}, grid, visited)
		s = calc(grid, visited)
		if s > max {
			max = s
		}
	}

	fmt.Printf("%d\n", max)
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
