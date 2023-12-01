package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func part2(r io.Reader) {
	dict := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	s := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := scanner.Text()

		for i := 0; i < len(l); i++ {
			c := l[i]
			if c >= '1' && c <= '9' {
				s += int(c-'0') * 10
				goto next
			}
			for k, v := range dict {
				if i+len(k) <= len(l) && l[i:i+len(k)] == k {
					s += v * 10
					goto next
				}
			}
		}
	next:
		for i := len(l) - 1; i >= 0; i-- {
			c := l[i]
			if c >= '1' && c <= '9' {
				s += int(c - '0')
				goto next2
			}
			for k, v := range dict {
				if i+len(k) <= len(l) && l[i:i+len(k)] == k {
					s += v
					goto next2
				}
			}
		}
	next2:
	}
	fmt.Printf("%d\n", s)
}

func part1(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := scanner.Text()

		for i := 0; i < len(l); i++ {
			c := l[i]
			if c >= '1' && c <= '9' {
				s += int(c-'0') * 10
				break
			}
		}

		for i := len(l) - 1; i >= 0; i-- {
			c := l[i]
			if c >= '1' && c <= '9' {
				s += int(c - '0')
				break
			}
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
