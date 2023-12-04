package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func countMatches(l string) int {
	cards := strings.Split(l, ":")
	cards = strings.Split(cards[1], "|")
	re := regexp.MustCompile(`\d+`) // numbers
	matches := re.FindAllString(cards[1], -1)

	d := map[string]struct{}{}
	for _, m := range matches {
		d[m] = struct{}{}
	}

	matches = re.FindAllString(cards[0], -1)
	n := 0
	for _, m := range matches {
		if _, ok := d[m]; ok {
			n++
		}
	}
	return n
}

func part1(lines []string) {
	s := 0

	for _, l := range lines {
		n := countMatches(l)
		if n > 0 {
			s += 1 << (n - 1)
		}
	}
	fmt.Printf("%d\n", s)
}

func part2(lines []string) {
	s := 0
	cnt := make([]int, len(lines))

	for i, l := range lines {
		cnt[i]++
		n := countMatches(l)
		for j := i + 1; j < i+1+n && j < len(lines); j++ {
			cnt[j] += cnt[i]
		}
	}

	for _, n := range cnt {
		s += n
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
