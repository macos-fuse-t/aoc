package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Rule struct {
	part byte
	comp byte
	val  int
	next string
}

func (r Rule) execute(parts map[byte]int) (bool, string) {
	switch r.comp {
	case '<':
		return parts[r.part] < r.val, r.next
	case '>':
		return parts[r.part] > r.val, r.next
	default:
		return true, r.next
	}
}

type Workflow struct {
	name  string
	rules []Rule
}

func (w Workflow) execute(parts map[byte]int) string {
	for _, r := range w.rules {
		if ok, next := r.execute(parts); ok {
			return next
		}
	}
	return w.rules[len(w.rules)-1].next
}

func exec(workflows map[string]Workflow, parts map[byte]int) bool {
	next := "in"
	for {
		next = workflows[next].execute(parts)
		if next == "R" {
			return false
		} else if next == "A" {
			break
		}
	}
	return true
}

func part1(workflows map[string]Workflow, allParts []map[byte]int) {
	s := 0
	for _, parts := range allParts {
		if exec(workflows, parts) {
			for _, p := range parts {
				s += p
			}
		}
	}
	fmt.Printf("%d\n", s)
}

type Range struct {
	min, max int
}

func part2(workflows map[string]Workflow) {
	dmin := map[byte][]int{}
	dmax := map[byte][]int{}
	for _, w := range workflows {
		for _, r := range w.rules {
			if r.comp == '>' {
				dmin[r.part] = append(dmin[r.part], r.val+1)
				dmax[r.part] = append(dmax[r.part], r.val)
			} else if r.comp == '<' {
				dmin[r.part] = append(dmin[r.part], r.val)
				dmax[r.part] = append(dmax[r.part], r.val-1)
			}
		}
	}
	for p := range dmin {
		dmin[p] = append(dmin[p], 1)
		sort.Slice(dmin[p], func(i, j int) bool {
			return dmin[p][i] < dmin[p][j]
		})
	}
	for p := range dmax {
		dmax[p] = append(dmax[p], 4000)
		sort.Slice(dmax[p], func(i, j int) bool {
			return dmax[p][i] < dmax[p][j]
		})
	}

	n := uint64(0)
	for xi, x := range dmin['x'] {
		for mi, m := range dmin['m'] {
			for si, s := range dmin['s'] {
				for ai, a := range dmin['a'] {
					parts := map[byte]int{'x': x, 'm': m, 's': s, 'a': a}
					if exec(workflows, parts) {
						n += (uint64(dmax['a'][ai]) - uint64(a) + 1) *
							(uint64(dmax['s'][si]) - uint64(s) + 1) *
							(uint64(dmax['m'][mi]) - uint64(m) + 1) *
							(uint64(dmax['x'][xi]) - uint64(x) + 1)
					}
				}
			}
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

	workflows := map[string]Workflow{}
	allParts := []map[byte]int{}

	state := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if state == 0 {
			if scanner.Text() == "" {
				state = 1
				continue
			}

			p := regexp.MustCompile(`^(.+)\{(.+)\}$`).FindStringSubmatch(l)
			name, rules := p[1], p[2]

			rule := strings.Split(rules, ",")

			w := Workflow{name: name}
			for _, s := range rule {
				var r Rule
				if strings.ContainsAny(s, "<>") {
					fmt.Sscanf(s[2:], "%d:%s", &r.val, &r.next)
					r.part = s[0]
					r.comp = s[1]
				} else {
					r.next = s
				}
				w.rules = append(w.rules, r)
			}
			workflows[w.name] = w
		} else {
			var x, m, a, s int
			fmt.Sscanf(l, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
			allParts = append(allParts, map[byte]int{'x': x, 'm': m, 'a': a, 's': s})
		}
	}

	part1(workflows, allParts)
	part2(workflows)
}
