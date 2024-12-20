package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part2(lines []string, towels []string) {
	n := 0
	for _, l := range lines {
		n += search2(l, towels)
	}
	fmt.Printf("%d\n", n)
}

var cache = map[string]int{}

func search2(p string, towels []string) int {
	if p == "" {
		return 1
	}
	if v, ok := cache[p]; ok {
		return v
	}

	n := 0
	for _, t := range towels {
		if strings.HasPrefix(p, t) {
			n += search2(p[len(t):], towels)
		}
	}
	cache[p] = n
	return n
}

func search(p string, towels []string) bool {
	if p == "" {
		return true
	}
	for _, t := range towels {
		if strings.HasPrefix(p, t) {
			r := search(p[len(t):], towels)
			if r {
				return true
			}
		}
	}
	return false
}

func part1(lines []string, towels []string) {
	n := 0
	for _, l := range lines {
		if search(l, towels) {
			n++
		}
	}
	fmt.Printf("%d\n", n)
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
	scanner.Scan()
	str := scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, l)
	}

	part := strings.Split(str, ", ")

	part1(lines, part)
	part2(lines, part)
}
