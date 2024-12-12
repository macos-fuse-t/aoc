package main

import (
	"bufio"
	"fmt"
	"os"
)

func part2(grid []string) {
	l := map[string]bool{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '.' {
				continue
			}
			for y1 := 0; y1 < len(grid); y1++ {
				for x1 := 0; x1 < len(grid[y1]); x1++ {
					if (x1 == x && y1 == y) || grid[y1][x1] != grid[y][x] {
						continue
					}

					for x2 := 0; x2 < len(grid[0]); x2++ {
						if (x2-x)*(y-y1)%(x-x1) != 0 {
							continue
						}
						y2 := (x2-x)*(y-y1)/(x-x1) + y
						if x2 >= 0 && x2 < len(grid[0]) && y2 >= 0 && y2 < len(grid) {
							l[fmt.Sprintf("%d:%d", y2, x2)] = true
						}
					}
				}
			}
		}
	}
	fmt.Printf("%d\n", len(l))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func part1(grid []string) {
	l := map[string]bool{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '.' {
				continue
			}
			for y1 := 0; y1 < len(grid); y1++ {
				for x1 := 0; x1 < len(grid[y1]); x1++ {
					if (x1 == x && y1 == y) || grid[y1][x1] != grid[y][x] {
						continue
					}
					//m := (y - y1) / (x - x1)
					x2 := 2*min(x, x1) - max(x, x1)
					y2 := (x2-x)*(y-y1)/(x-x1) + y
					if x2 >= 0 && x2 < len(grid[0]) && y2 >= 0 && y2 < len(grid) {
						l[fmt.Sprintf("%d:%d", y2, x2)] = true
					}

					x2 = 2*max(x, x1) - min(x, x1)
					y2 = (x2-x)*(y-y1)/(x-x1) + y
					if x2 >= 0 && x2 < len(grid[0]) && y2 >= 0 && y2 < len(grid) {
						l[fmt.Sprintf("%d:%d", y2, x2)] = true
					}
				}
			}
		}
	}
	//fmt.Printf("%v\n", l)
	fmt.Printf("%d\n", len(l))
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		grid = append(grid, l)
	}

	part1(grid)
	part2(grid)
}
