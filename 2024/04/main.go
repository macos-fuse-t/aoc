package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func part2(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)
	txt := []string{}

	for scanner.Scan() {
		l := scanner.Text()
		txt = append(txt, l)
	}

	for i := 0; i < len(txt); i++ {
		for j := 0; j < len(txt[i]); j++ {
			if search(txt, i, j, 1, 1, "MAS", []int{}) > 0 && search(txt, i, j+2, 1, -1, "MAS", []int{}) > 0 {
				s += 1
			}
			if search(txt, i, j, 1, 1, "MAS", []int{}) > 0 && search(txt, i, j+2, 1, -1, "SAM", []int{}) > 0 {
				s += 1
			}
			if search(txt, i, j, 1, 1, "SAM", []int{}) > 0 && search(txt, i, j+2, 1, -1, "MAS", []int{}) > 0 {
				s += 1
			}
			if search(txt, i, j, 1, 1, "SAM", []int{}) > 0 && search(txt, i, j+2, 1, -1, "SAM", []int{}) > 0 {
				s += 1
			}
		}
	}
	fmt.Printf("%d\n", s)
}

func search(txt []string, i, j, dy, dx int, s string, indices []int) int {
	if s == "" {
		return 1
	}
	if i < 0 || i >= len(txt) || j < 0 || j >= len(txt[0]) {
		return 0
	}
	n := 0
	if s[0] == txt[i][j] {
		indices = append(indices, i*len(txt)+j)
		n += search(txt, i+dy, j+dx, dy, dx, s[1:], indices)
	}
	return n
}

func part1(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)
	txt := []string{}

	for scanner.Scan() {
		l := scanner.Text()
		txt = append(txt, l)
	}

	for i := 0; i < len(txt); i++ {
		for j := 0; j < len(txt[i]); j++ {
			s += search(txt, i, j, 0, 1, "XMAS", []int{})
			s += search(txt, i, j, 0, -1, "XMAS", []int{})
			s += search(txt, i, j, 1, 0, "XMAS", []int{})
			s += search(txt, i, j, -1, 0, "XMAS", []int{})
			s += search(txt, i, j, 1, 1, "XMAS", []int{})
			s += search(txt, i, j, 1, -1, "XMAS", []int{})
			s += search(txt, i, j, -1, 1, "XMAS", []int{})
			s += search(txt, i, j, -1, -1, "XMAS", []int{})
		}
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

	part1(file)
	file.Seek(0, 0)
	part2(file)
}
