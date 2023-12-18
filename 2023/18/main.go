package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Line struct {
	p1 Point
	p2 Point
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func (l Line) len() int {
	if l.p1.x == l.p2.x {
		return abs(l.p1.y - l.p2.y)
	}
	return abs(l.p1.x - l.p2.x)

}

// https://en.wikipedia.org/wiki/Shoelace_formula#Example
func part1(lines []Line) {
	s := int(0)
	for i := 0; i < len(lines); i++ {
		s += (lines[i].p1.y + lines[i].p2.y) * (lines[i].p1.x - lines[i].p2.x)
	}
	for _, l := range lines {
		s += l.len()
	}

	fmt.Printf("%d\n", s/2+1)
}

func part2(lines []Line) {
	part1(lines)
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

	lines := []Line{}
	scanner := bufio.NewScanner(file)
	x1, y1, x2, y2 := 0, 0, 0, 0
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(a[1])
		switch a[0][0] {
		case 'R':
			x2 = x1 + v
		case 'L':
			x2 = x1 - v
		case 'D':
			y2 = y1 + v
		case 'U':
			y2 = y1 - v
		}
		lines = append(lines, Line{p1: Point{x1, y1}, p2: Point{x2, y2}})
		x1, y1 = x2, y2
	}

	file.Seek(0, 0)
	lines2 := []Line{}
	scanner = bufio.NewScanner(file)
	x1, y1, x2, y2 = 0, 0, 0, 0
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), " ")
		v, _ := strconv.ParseInt(a[2][2:7], 16, 64)
		switch a[2][7] {
		case '0':
			x2 = x1 + int(v)
		case '2':
			x2 = x1 - int(v)
		case '1':
			y2 = y1 + int(v)
		case '3':
			y2 = y1 - int(v)
		}
		lines2 = append(lines2, Line{p1: Point{x1, y1}, p2: Point{x2, y2}})
		x1, y1 = x2, y2
	}

	part1(lines)
	part2(lines2)
}
