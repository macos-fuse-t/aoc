package main

import (
	"bufio"
	"fmt"
	"os"
)

type region struct {
	id              byte
	area, perimeter int
	sides           int
}

var regions = []*region{}
var cells = [][]*region{}

func part2(grid []string) {
	s := 0

	for y := -1; y < len(grid); y++ {
		prev_up, prev_down := byte(0), byte(0)
		for x := 0; x < len(grid[0]); x++ {
			up, down := byte(0), byte(0)
			if y >= 0 {
				up = grid[y][x]
			}
			if y < len(grid)-1 {
				down = grid[y+1][x]
			}

			if up != down {
				if (up != prev_up || prev_down == prev_up) && y >= 0 {
					cells[y][x].sides++
				}
				if (down != prev_down || prev_down == prev_up) && y < len(grid)-1 {
					cells[y+1][x].sides++
				}
			}

			prev_up = up
			prev_down = down
		}
	}
	for x := -1; x < len(grid[0]); x++ {
		prev_left, prev_right := byte(0), byte(0)
		for y := 0; y < len(grid); y++ {
			left, right := byte(0), byte(0)
			if x >= 0 {
				left = grid[y][x]
			}
			if x < len(grid[0])-1 {
				right = grid[y][x+1]
			}

			if left != right {
				if (left != prev_left || prev_left == prev_right) && x >= 0 {
					cells[y][x].sides++
				}
				if (right != prev_right || prev_right == prev_left) && x < len(grid[0])-1 {
					cells[y][x+1].sides++
				}
			}

			prev_left = left
			prev_right = right
		}
	}

	for _, r := range regions {
		s += r.area * r.sides
	}
	fmt.Printf("%d\n", s)
}

func search(grid []string, x, y int, id byte) {
	if x < 0 || y < 0 || x >= len(cells[0]) || y >= len(cells) || cells[y][x] != nil || grid[y][x] != id {
		return
	}
	if x > 0 && grid[y][x-1] == grid[y][x] && cells[y][x-1] != nil {
		cells[y][x] = cells[y][x-1]
		cells[y][x].area++
	} else if y > 0 && grid[y-1][x] == grid[y][x] && cells[y-1][x] != nil {
		cells[y][x] = cells[y-1][x]
		cells[y][x].area++
	} else if x < len(grid[0])-1 && grid[y][x+1] == grid[y][x] && cells[y][x+1] != nil {
		cells[y][x] = cells[y][x+1]
		cells[y][x].area++
	} else if y < len(grid)-1 && grid[y+1][x] == grid[y][x] && cells[y+1][x] != nil {
		cells[y][x] = cells[y+1][x]
		cells[y][x].area++
	} else {
		r := region{id: grid[y][x], area: 1, perimeter: 0}
		cells[y][x] = &r
		regions = append(regions, &r)
	}
	search(grid, x-1, y, grid[y][x])
	search(grid, x+1, y, grid[y][x])
	search(grid, x, y-1, grid[y][x])
	search(grid, x, y+1, grid[y][x])
}

func part1(grid []string) {
	s := 0
	for y := 0; y < len(grid); y++ {
		cells = append(cells, make([]*region, len(grid[y])))
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			search(grid, x, y, grid[y][x])
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			cells[y][x].perimeter += 4
			if x > 0 && cells[y][x-1] == cells[y][x] {
				cells[y][x].perimeter -= 1
			}
			if x < len(cells[0])-1 && cells[y][x+1] == cells[y][x] {
				cells[y][x].perimeter -= 1
			}
			if y > 0 && cells[y-1][x] == cells[y][x] {
				cells[y][x].perimeter -= 1
			}
			if y < len(cells)-1 && cells[y+1][x] == cells[y][x] {
				cells[y][x].perimeter -= 1
			}
		}
	}

	for _, r := range regions {
		//fmt.Printf("%v\n", *r)
		s += r.area * r.perimeter
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

	scanner := bufio.NewScanner(file)
	grid := []string{}
	for scanner.Scan() {
		l := scanner.Text()
		grid = append(grid, l)
	}

	part1(grid)
	part2(grid)
}
