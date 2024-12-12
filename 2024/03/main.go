package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func part2(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)
	do := true
	for scanner.Scan() {
		l := scanner.Text()
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

		// Find all matches
		matches := re.FindAllStringSubmatch(l, -1)

		// Iterate over matches and print them
		for _, match := range matches {
			//fmt.Printf("Full match: %s\n", match[0])
			if match[0] == "do()" {
				do = true
				continue
			} else if match[0] == "don't()" {
				do = false
				continue
			}
			if !do {
				continue
			}
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			s += num1 * num2
		}
	}
	fmt.Printf("%d\n", s)
}

func part1(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		l := scanner.Text()
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

		// Find all matches
		matches := re.FindAllStringSubmatch(l, -1)

		// Iterate over matches and print them
		for _, match := range matches {
			//fmt.Printf("Full match: %s, Number1: %s, Number2: %s\n", match[0], match[1], match[2])
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			s += num1 * num2
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
