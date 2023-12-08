package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vertex struct {
	l, r string
}

func part1(path string, graph map[string]Vertex) {
	s := 0
	v := "AAA"
	for v != "ZZZ" {
		if path[s%len(path)] == 'L' {
			v = graph[v].l
		} else {
			v = graph[v].r
		}
		s++
	}

	fmt.Printf("%d\n", s)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func part2(path string, graph map[string]Vertex, va []string) {
	periods := []int{}

	for _, v := range va {
		s := 0
		d := map[string]int{}
		for {
			if v[2] == 'Z' {
				if d[fmt.Sprintf("%s_%d", v, s%len(path))] == 1 {
					break
				}
				d[fmt.Sprintf("%s_%d", v, s%len(path))] = 1
				periods = append(periods, s)
			}

			if path[s%len(path)] == 'L' {
				v = graph[v].l
			} else {
				v = graph[v].r
			}
			s++
		}
	}

	s := 1
	for _, n := range periods {
		s = lcm(s, n)
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

	graph := map[string]Vertex{}
	va := []string{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	path := scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		var v, l, r string
		fmt.Sscanf(scanner.Text(), "%s = (%3s, %3s)", &v, &l, &r)
		graph[v] = Vertex{l, r}
		if v[2] == 'A' {
			va = append(va, v)
		}
	}

	part1(path, graph)
	part2(path, graph, va)
}
