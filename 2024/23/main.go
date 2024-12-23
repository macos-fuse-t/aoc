package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func part2(e map[string][]string) {
	m := []map[string]struct{}{}
	for k := range e {
		m = append(m, map[string]struct{}{k: {}})
	}

	for k, l := range e {
		for _, v := range l {
			for _, m1 := range m {
				_, ok1 := m1[k]
				_, ok2 := m1[v]
				if !ok1 && !ok2 || (ok1 && ok2) {
					continue
				}
				n := k
				if !ok2 {
					n = v
				}
				valid := true
				for p := range m1 {
					if !contains(e[n], p) {
						valid = false
						break
					}
				}
				if valid {
					m1[n] = struct{}{}
				}
			}
		}
	}

	max := 0
	str := ""
	for _, k := range m {
		if len(k) > max {
			max = len(k)
			l := []string{}
			for v := range k {
				l = append(l, v)
			}
			sort.Strings(l)
			str = ""
			for i, v := range l {
				if i > 0 {
					str += ","
				}
				str += v
			}
		}
	}
	fmt.Printf("%s\n", str)
}

func contains(l []string, v string) bool {
	for _, k := range l {
		if k == v {
			return true
		}
	}
	return false
}

func part1(e map[string][]string) {
	m := map[string]bool{}
	for k, l := range e {
		if k[0] == 't' {
			for _, v := range l {
				for _, k1 := range e[v] {
					if k1 != k && contains(l, k1) {
						l1 := []string{k, k1, v}
						sort.Strings(l1)
						m[l1[0]+l1[1]+l1[2]] = true
					}
				}
			}
		}
	}
	fmt.Printf("%d\n", len(m))
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

	e := map[string][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		p := strings.Split(l, "-")
		e[p[0]] = append(e[p[0]], p[1])
		e[p[1]] = append(e[p[1]], p[0])
	}

	part1(e)
	part2(e)
}
