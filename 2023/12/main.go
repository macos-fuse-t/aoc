package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validate(s string, cond []int) (bool, bool) {
	nw := strings.Count(s, "?")
	nn := strings.Count(s, "#")

	if nn > 0 && len(cond) == 0 {
		return false, true
	}
	s = strings.ReplaceAll(s, "?", ".")
	end := nw == 0

	a := strings.FieldsFunc(s, func(r rune) bool {
		return r == '.'
	})
	if len(a) != len(cond) {
		return false, end
	}
	for i, c := range cond {
		if c != len(a[i]) {
			return false, end
		}
	}

	return true, true
}

func eat(s string, cnt int) (int, bool) {
	if len(s) < cnt {
		return 0, false
	}
	for i := 0; i < cnt; i++ {
		if s[i] == '.' {
			return 0, false
		}
	}
	if len(s) == cnt {
		return cnt, true
	}

	if s[cnt] == '.' {
		return cnt + 1, true
	}
	if s[cnt] == '?' {
		return cnt + 1, true
	}
	return 0, false
}

func countArrangements(s string, cond []int, visited map[string]int) int {
	key := s + fmt.Sprintf("%v", cond)
	if n, ok := visited[key]; ok {
		return n
	}

	visited[key] = 0
	if valid, end := validate(s, cond); end {
		if valid {
			visited[key] = 1
			return 1
		}
		return 0
	}

	if s[0] == '.' {
		n := countArrangements(s[1:], cond, visited)
		visited[key] = n
		return n
	}

	n := 0
	if s[0] == '?' {
		n += countArrangements(s[1:], cond, visited) // .
	}

	cnt, ok := eat(s, cond[0])
	if !ok {
		visited[key] = n
		return n
	}
	n += countArrangements(s[cnt:], cond[1:], visited) // #
	visited[key] = n
	return n
}

func part1(lines []string) {
	s := 0

	for _, l := range lines {
		cond := []int{}
		a := strings.Split(l, " ")
		for _, c := range strings.Split(a[1], ",") {
			v, _ := strconv.Atoi(c)
			cond = append(cond, v)
		}
		visited := map[string]int{}
		s += countArrangements(a[0], cond, visited)
	}

	fmt.Printf("%d\n", s)
}

func part2(lines []string) {
	s := 0

	for _, l := range lines {
		origCond := []int{}
		a := strings.Split(l, " ")
		for _, c := range strings.Split(a[1], ",") {
			v, _ := strconv.Atoi(c)
			origCond = append(origCond, v)
		}

		cond := []int{}
		str := ""
		for i := 0; i < 5; i++ {
			cond = append(cond, origCond...)
			if i != 0 {
				str = str + "?"
			}
			str = str + a[0]
		}

		visited := map[string]int{}
		s += countArrangements(str, cond, visited)
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
