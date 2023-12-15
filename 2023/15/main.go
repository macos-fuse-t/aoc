package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calc(s string) int {
	h := 0
	for _, c := range s {
		h = ((h + int(c)) * 17) % 256
	}
	return h
}

func part1(lines []string) {
	s := 0
	for _, str := range lines {
		s += calc(str)
	}
	fmt.Printf("%d\n", s)
}

type Item struct {
	label string
	val   int
}

func part2(lines []string) {
	l := [256][]Item{}
	for _, str := range lines {
		a := strings.FieldsFunc(str, func(c rune) bool {
			return c == '-' || c == '='
		})
		h := calc(a[0])

		for i, item := range l[h] {
			if item.label == a[0] {
				if len(a) == 1 {
					l[h] = append(l[h][:i], l[h][i+1:]...) // remove
				} else {
					val, _ := strconv.Atoi(a[1])
					l[h][i] = Item{a[0], val}
				}
				goto skip
			}
		}
		if len(a) > 1 {
			val, _ := strconv.Atoi(a[1])
			l[h] = append(l[h], Item{a[0], val})
		}
	skip:
	}

	s := 0
	for i, box := range l {
		for j, item := range box {
			s += (i + 1) * (j + 1) * item.val
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

	lines := []string{}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, s := range strings.Split(scanner.Text(), ",") {
		lines = append(lines, s)
	}
	part1(lines)
	part2(lines)
}
