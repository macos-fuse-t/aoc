package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type gate struct {
	input1, input2 string
	op             string
	isLeaf         bool
	val            int
}

func validate(m map[string]gate, r string, x, y uint64, i int) bool {
	n, _ := strconv.Atoi(r[1:])
	if i < 0 {
		mask := uint64(1)<<(n+1) - 1
		res := uint64(0)
		for k := 0; k < n+1; k++ {
			r = fmt.Sprintf("z%02d", k)
			if check_loop(m, r, map[string]bool{}) {
				return false
			}
			res |= calc(m, r) << k
		}
		return res == (x+y)&mask
	}
	rx := fmt.Sprintf("x%02d", i)
	ry := fmt.Sprintf("y%02d", i)
	// first 4 bits only
	if n-i > 3 {
		m[rx] = gate{isLeaf: true, val: 0}
		m[ry] = gate{isLeaf: true, val: 0}
		return validate(m, r, x<<1, y<<1, i-1)
	}

	res := true
	for _, v := range [][]uint64{{0, 0}, {0, 1}, {1, 0}, {1, 1}} {
		m[rx] = gate{isLeaf: true, val: int(v[0])}
		m[ry] = gate{isLeaf: true, val: int(v[1])}
		if !validate(m, r, (x<<1)|v[0], (y<<1)|v[1], i-1) {
			res = false
			break
		}
	}
	return res
}

func part2(m map[string]gate) {
	vars := []string{}
	for z := 0; z < 45; z++ {
		nz := fmt.Sprintf("z%02d", z)
		r := validate(m, nz, 0, 0, z)
		if !r {
			s := 0
			for n1 := range m {
				if s == 0 {
					n1 = nz
					s++
				}
				ng1 := m[n1]
				for n2 := range m {
					if n1[0] == 'z' && n2[0] == 'z' {
						continue
					}
					ng2 := m[n2]

					m[n1] = ng2
					m[n2] = ng1
					r = validate(m, nz, 0, 0, z)
					if r {
						fmt.Printf("%s %s\n", n2, n1)
						vars = append(vars, n1, n2)
						break
					} else {
						m[n1] = ng1
						m[n2] = ng2
					}
				}
				if r {
					break
				}
			}
		}
	}

	sort.Strings(vars)
	fmt.Printf("%v\n", vars)
}

func check_loop(m map[string]gate, n string, c map[string]bool) bool {
	g := m[n]
	if c[n] {
		return true
	}
	if g.isLeaf {
		return false
	}
	c[n] = true
	return check_loop(m, g.input1, c) || check_loop(m, g.input2, c)
}

func calc(m map[string]gate, n string) uint64 {
	g := m[n]
	if g.isLeaf {
		return uint64(g.val)
	}

	if g.op == "AND" {
		return calc(m, g.input1) & calc(m, g.input2)
	}
	if g.op == "OR" {
		return calc(m, g.input1) | calc(m, g.input2)
	}
	if g.op == "XOR" {
		return calc(m, g.input1) ^ calc(m, g.input2)
	}
	return 0
}

func part1(m map[string]gate) {
	s := uint64(0)
	for i := 0; i < 64; i++ {
		r := fmt.Sprintf("z%02d", i)
		if _, ok := m[r]; !ok {
			break
		}
		s = s | calc(m, r)<<i

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

	m := map[string]gate{}

	state := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			state = 1
			continue
		}
		if state == 0 {
			p := strings.Split(l, ": ")
			m[p[0]] = gate{isLeaf: true, val: int(p[1][0]) - '0'}
		} else {
			p := strings.Split(l, " ")
			m[p[4]] = gate{input1: p[0], input2: p[2], op: p[1]}
		}
	}

	part1(m)
	part2(m)
}
