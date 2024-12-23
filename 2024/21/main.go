package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type coord struct {
	x, y int
}

type state struct {
	x, y int
	id   int
}

var keypad1 = map[byte]coord{
	'7': {0, 0}, '8': {1, 0}, '9': {2, 0},
	'4': {0, 1}, '5': {1, 1}, '6': {2, 1},
	'1': {0, 2}, '2': {1, 2}, '3': {2, 2},
	'0': {1, 3}, 'A': {2, 3}}
var keypad2 = map[byte]coord{
	'^': {1, 0}, 'A': {2, 0}, '<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
}

var coord2byte = map[coord]byte{
	{1, 0}: '^', {2, 0}: 'A', {0, 1}: '<', {1, 1}: 'v', {2, 1}: '>',
}

func search(n string, robots []state, id int, seq string, res *[]string) string {
	if id == len(robots) {
		return n
	}
	/*if id == 2 {
		fmt.Printf("%s\n", n)
		return n
	}*/

	s := &robots[id]
	old_state := *s
	keypad := keypad1
	if id > 0 {
		keypad = keypad2
	}

	if n == "" {
		r := search(seq, robots, id+1, "", res)
		if res != nil {
			if len(*res) == 0 || len(r) == len((*res)[0]) {
				*res = append(*res, r)
			} else {
				if len(r) < len((*res)[0]) {
					*res = []string{}
					*res = append(*res, r)
				}
			}
		}
		return r
	}

	d := n[0]
	if s.x == keypad[d].x && s.y == keypad[d].y {
		seq += "A"
		return search(n[1:], robots, id, seq, res)
	}
	if s.x == keypad[d].x {
		if s.y < keypad[d].y {
			seq += "v"
			s.y++
		} else {
			seq += "^"
			s.y--
		}
		r := search(n, robots, id, seq, res)
		*s = old_state
		return r
	}
	if s.y == keypad[d].y {
		if s.x < keypad[d].x {
			seq += ">"
			s.x++
		} else {
			seq += "<"
			s.x--
		}
		r := search(n, robots, id, seq, res)
		*s = old_state
		return r
	}

	if s.x < keypad[d].x {
		seq += ">"
		s.x++
	} else {
		seq += "<"
		s.x--
	}

	r1 := ""
	r2 := ""
	if !((id == 0 && s.x == 0 && s.y == 3) || (id > 0 && s.x == 0 && s.y == 0)) {
		r1 = search(n, robots, id, seq, res)
	}

	*s = old_state
	seq = seq[:len(seq)-1]
	if s.y < keypad[d].y {
		seq += "v"
		s.y++
	} else {
		seq += "^"
		s.y--
	}

	if !((id == 0 && s.x == 0 && s.y == 3) || (id > 0 && s.x == 0 && s.y == 0)) {
		r2 = search(n, robots, id, seq, res)
	}

	*s = old_state
	if r1 == "" {
		return r2
	}
	if r2 == "" {
		return r1
	}

	if len(r1) < len(r2) {
		return r1
	}
	return r2
}

func buildTable(n string, t map[string]string) {
	s := state{2, 0, 0}
	d := coord2byte[coord{s.x, s.y}]
	j := 0
	for i := 0; i < len(n); i++ {
		if n[i] == 'A' {
			d1 := coord2byte[coord{s.x, s.y}]
			str := fmt.Sprintf("%c%c", d, d1)
			t[str] = n[j : i+1]
			d = d1
			j = i + 1
			continue
		}
		if n[i] == '^' {
			s.y--
		} else if n[i] == 'v' {
			s.y++
		} else if n[i] == '<' {
			s.x--
		} else if n[i] == '>' {
			s.x++
		}
	}

	n = "><^vA"
	for i := 0; i < len(n); i++ {
		for j := i; j < len(n); j++ {
			robots := []state{
				{2, 3, 0},
				{2, 0, 1},
				{2, 0, 2},
			}
			str := fmt.Sprintf("%c%c", n[i], n[j])
			if t[str] == "" {
				robots[2].x = keypad2[str[0]].x
				robots[2].y = keypad2[str[0]].y
				p := search(str[1:], robots, 2, "", nil)
				t[str] = p
			}

			str = fmt.Sprintf("%c%c", n[j], n[i])
			if t[str] == "" {
				robots[2].x = keypad2[str[0]].x
				robots[2].y = keypad2[str[0]].y
				p := search(str[1:], robots, 2, "", nil)
				t[str] = p
			}
		}
	}
}

func part2(nums []string) {
	s := uint64(0)
	for _, n := range nums {
		robots := []state{
			{2, 3, 0},
			{2, 0, 1},
			{2, 0, 2},
		}

		res := []string{}
		search(n, robots, 0, "", &res)

		minlen := math.MaxInt
		//minp := ""
		for _, p := range res {
			t := make(map[string]string)
			buildTable(p, t)

			m := map[string]int{}
			n := map[string]int{}
			for k := range t {
				m[k] = 0
				n[k] = 0
			}

			m[fmt.Sprintf("A%c", p[0])]++
			for i := 0; i < len(p)-1; i++ {
				str := fmt.Sprintf("%c%c", p[i], p[i+1])
				m[str]++
			}

			for j := 0; j < 23; j++ {
				for k := range m {
					if m[k] == 0 {
						continue
					}
					p1 := "A" + t[k]
					for i := 0; i < len(p1)-1; i++ {
						str := fmt.Sprintf("%c%c", p1[i], p1[i+1])
						n[str] += m[k]
					}
				}
				m = n
				n = map[string]int{}
			}

			v := 0
			for k := range m {
				v += m[k]
			}
			if v < minlen {
				minlen = v
				//minp = p
			}
		}

		i, _ := strconv.Atoi(n[:len(n)-1])
		fmt.Printf("%d %d\n", minlen, i)
		s += uint64(i) * uint64(minlen)
		//break
	}
	fmt.Printf("%d\n", s)
}

func part1(nums []string) {
	s := 0
	for _, n := range nums {
		robots := []state{
			{2, 3, 0},
			{2, 0, 1},
			{2, 0, 2},
			//{2, 0, 3},
		}
		p := search(n, robots, 0, "", nil)
		i, _ := strconv.Atoi(n[:len(n)-1])
		fmt.Printf("%s %d %d\n", p, len(p), i)
		s += i * len(p)

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

	n := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		n = append(n, l)
	}

	part1(n)
	part2(n)
}
