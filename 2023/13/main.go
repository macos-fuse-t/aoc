package main

import (
	"bufio"
	"fmt"
	"os"
)

func transpose(lines []string) []string {
	transposed := make([]string, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		for _, str := range lines {
			if i < len(str) {
				transposed[i] += string(str[i])
			}
		}
	}
	return transposed
}

func calc(lines []string, old int) int {
	for i := 0; i < len(lines)-1; i++ {
		if i == old {
			continue
		}
		if lines[i] == lines[i+1] {
			for j := 0; i-j >= 0 && i+1+j < len(lines); j++ {
				if lines[i-j] != lines[i+1+j] {
					goto next
				}
			}
			return i + 1
		}
	next:
	}
	return 0
}

func part1(lines []string) {
	s := 0

	first := 0
	for i, l := range lines {
		if l == "" || i == len(lines)-1 {
			a := lines[first:i]
			if i == len(lines)-1 {
				a = lines[first : i+1]
			}

			n1 := calc(a, -1)
			n2 := calc(transpose(a), -1)
			s += n1*100 + n2
			first = i + 1
		}
	}

	fmt.Printf("%d\n", s)
}

func newSlice(a []string, i, j int) []string {
	c := '.'
	if a[i][j] == '.' {
		c = '#'
	}
	b := make([]string, len(a))
	copy(b, a)
	r := []rune(b[i])
	r[j] = c
	b[i] = string(r)
	return b
}

func part2(lines []string) {
	s := 0

	first := 0
	for i, l := range lines {
		if l == "" || i == len(lines)-1 {
			a := lines[first:i]
			if i == len(lines)-1 {
				a = lines[first : i+1]
			}

			n1 := calc(a, -1)
			n2 := calc(transpose(a), -1)

			for k := 0; k < len(a); k++ {
				for m := 0; m < len(a[k]); m++ {
					b := newSlice(a, k, m)

					n11 := calc(b, n1-1)
					if n11 > 0 {
						n1 = n11
						n2 = 0
						goto next
					}
				}
				for m := 0; m < len(a[k]); m++ {
					b := newSlice(a, k, m)

					n22 := calc(transpose(b), n2-1)
					if n22 > 0 {
						n1 = 0
						n2 = n22
						goto next
					}
				}
			}
		next:
			s += n1*100 + n2
			first = i + 1
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
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	part1(lines)
	part2(lines)
}
