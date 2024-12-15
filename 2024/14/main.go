package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func print(grid [][]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] > 0 {
				fmt.Printf("%d", grid[y][x])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

type item struct {
	x0, y0, vx, vy int
	cycle          int
}

var w = 101
var h = 103

func calc(i item, n int) (int, int) {
	x := i.x0 + i.vx*n
	y := i.y0 + i.vy*n
	if x >= 0 {
		x = x % w
	} else {
		x = (w - (-x)%w) % w
	}
	if y >= 0 {
		y = y % h
	} else {
		y = (h - (-y)%h) % h
	}
	return x, y
}

func part2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	items := []item{}

	for scanner.Scan() {
		l := scanner.Text()
		var x0, y0, vx, vy int
		fmt.Sscanf(l, "p=%d,%d v=%d,%d", &x0, &y0, &vx, &vy)
		items = append(items, item{x0, y0, vx, vy, 0})
	}

	for n := 0; n < w*h; n++ {
		grid := make([][]int, h)
		for y := 0; y < h; y++ {
			grid[y] = make([]int, w)
		}
		for _, i := range items {
			x, y := calc(i, n)
			grid[y][x]++
		}

		f := true
		for k := 1; k < 500; k++ {
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					if grid[y][x] > 0 && grid[y][x] != k {
						f = false
						break
					}
				}
				if !f {
					break
				}
			}
			if f {
				break
			}
		}
		if f {
			fmt.Printf("n %d\n", n)
			print(grid)
		}
	}
}

func part1(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)
	n := 100

	q1, q2, q3, q4 := 0, 0, 0, 0

	for scanner.Scan() {
		l := scanner.Text()
		var x0, y0, vx, vy int
		fmt.Sscanf(l, "p=%d,%d v=%d,%d", &x0, &y0, &vx, &vy)

		x, y := calc(item{x0, y0, vx, vy, 0}, n)
		if x < w/2 && y < h/2 {
			q1++
		} else if x > w/2 && y < h/2 {
			q2++
		} else if x < w/2 && y > h/2 {
			q3++
		} else if x > w/2 && y > h/2 {
			q4++
		}
	}
	s = q1 * q2 * q3 * q4
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
