package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func part2(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)
	grid := [][]rune{}
	moves := ""
	startx, starty := 0, 0

	state := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" || l == "\n" {
			state = 1
			continue
		}
		if state == 0 {
			l = strings.ReplaceAll(l, "#", "##")
			l = strings.ReplaceAll(l, "O", "[]")
			l = strings.ReplaceAll(l, ".", "..")
			l = strings.ReplaceAll(l, "@", "@.")
			if x := strings.Index(l, "@"); x > 0 {
				grid = append(grid, []rune(l[:x]+"."+l[x+1:]))
				starty = len(grid) - 1
				startx = x
			} else {
				grid = append(grid, []rune(l))
			}
		} else {
			moves += l
		}
	}

	s = move(grid, startx, starty, moves, true)
	fmt.Printf("%d\n", s)
}

func print(grid [][]rune, sx, sy int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if x == sx && y == sy {
				fmt.Printf("@")
			} else {
				fmt.Printf("%c", grid[y][x])
			}
		}
		fmt.Println()
	}
}

func copyGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for i := range grid {
		newGrid[i] = make([]rune, len(grid[i]))
		copy(newGrid[i], grid[i])
	}

	return newGrid
}

func move_box(grid [][]rune, x, y int, m rune, part2 bool) bool {
	if x >= len(grid[y]) || x < 0 || y >= len(grid[y]) || y < 0 ||
		grid[y][x] == '.' || grid[y][x] == '#' {
		return false
	}
	if m == '>' {
		move_box(grid, x+1, y, m, part2)
		if x >= len(grid[y])-1 || grid[y][x+1] != '.' {
			return false
		}
		grid[y][x+1] = grid[y][x]
		grid[y][x] = '.'
	} else if m == 'v' {
		if !part2 {
			move_box(grid, x, y+1, m, part2)
			if y >= len(grid)-1 || grid[y+1][x] != '.' {
				return false
			}
			grid[y+1][x] = 'O'
			grid[y][x] = '.'
		} else {
			xmin := x
			if grid[y][x] == ']' {
				xmin--
			}

			grid2 := copyGrid(grid)
			move_box(grid2, xmin, y+1, m, part2)
			move_box(grid2, xmin+1, y+1, m, part2)
			if y >= len(grid2)-1 || grid2[y+1][xmin] != '.' || grid2[y+1][xmin+1] != '.' {
				return false
			}
			move_box(grid, xmin, y+1, m, part2)
			move_box(grid, xmin+1, y+1, m, part2)
			if y >= len(grid)-1 || grid[y+1][xmin] != '.' || grid[y+1][xmin+1] != '.' {
				return false
			}
			grid[y+1][xmin] = '['
			grid[y+1][xmin+1] = ']'
			grid[y][xmin] = '.'
			grid[y][xmin+1] = '.'
		}
	} else if m == '<' {
		move_box(grid, x-1, y, m, part2)
		if x < 1 || grid[y][x-1] != '.' {
			return false
		}
		grid[y][x-1] = grid[y][x]
		grid[y][x] = '.'
	} else if m == '^' {
		if !part2 {
			move_box(grid, x, y-1, m, part2)
			if y < 1 || grid[y-1][x] != '.' {
				return false
			}
			grid[y-1][x] = 'O'
			grid[y][x] = '.'
		} else {
			xmin := x
			if grid[y][x] == ']' {
				xmin--
			}
			grid2 := copyGrid(grid)
			move_box(grid2, xmin, y-1, m, part2)
			move_box(grid2, xmin+1, y-1, m, part2)
			if grid2[y-1][xmin] != '.' || grid2[y-1][xmin+1] != '.' {
				return false
			}

			move_box(grid, xmin, y-1, m, part2)
			move_box(grid, xmin+1, y-1, m, part2)
			if y < 1 || grid[y-1][xmin] != '.' || grid[y-1][xmin+1] != '.' {
				return false
			}
			grid[y-1][xmin] = '['
			grid[y-1][xmin+1] = ']'
			grid[y][xmin] = '.'
			grid[y][xmin+1] = '.'
		}
	}
	return true
}

func move(grid [][]rune, x, y int, moves string, part2 bool) int {
	s := 0
	for _, m := range moves {
		/*if part2 {
			fmt.Printf("%d, %d, %c\n", x, y, m)
			print(grid, x, y)
		}*/
		if m == '>' {
			move_box(grid, x+1, y, m, part2)
			if x < len(grid[y])-1 && grid[y][x+1] == '.' {
				x++
			}
		} else if m == 'v' {
			move_box(grid, x, y+1, m, part2)
			if y < len(grid)-1 && grid[y+1][x] == '.' {
				y++
			}
		} else if m == '<' {
			move_box(grid, x-1, y, m, part2)
			if x > 0 && grid[y][x-1] == '.' {
				x--
			}
		} else if m == '^' {
			move_box(grid, x, y-1, m, part2)
			if y > 0 && grid[y-1][x] == '.' {
				y--
			}
		}
	}
	/*if part2 {
		print(grid, x, y)
	}*/
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'O' || grid[y][x] == '[' {
				s += x + y*100
			}
		}
	}
	return s
}

func part1(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)
	grid := [][]rune{}
	moves := ""
	startx, starty := 0, 0

	state := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" || l == "\n" {
			state = 1
			continue
		}
		if state == 0 {
			if x := strings.Index(l, "@"); x > 0 {
				grid = append(grid, []rune(l[:x]+"."+l[x+1:]))
				starty = len(grid) - 1
				startx = x
			} else {
				grid = append(grid, []rune(l))
			}
		} else {
			moves += l
		}
	}

	s = move(grid, startx, starty, moves, false)
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

	part1(file)
	file.Seek(0, 0)
	part2(file)
}
