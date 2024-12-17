package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cpu struct {
	A, B, C uint64
	ip      int
}

func calc(c *cpu, p []int, n []int, A uint64, l int) uint64 {
	if len(n) == 0 {
		return A
	}

	for k := uint64(0); k < 8; k++ {
		a := (k << ((3 * l) + 7)) | A

		c.A = a >> (3 * l)
		c.B = 0
		c.C = 0
		n1 := exec(c, p)
		if n1[0] == n[0] {
			r := calc(c, p, n[1:], a, l+1)
			if r > 0 {
				return r
			}
		}
	}
	return 0
}

func part2(c *cpu, p []int) {
	// B = A & 7
	// B = (A & 7) ^ 7
	// C = A >> ( (A & 7) ^ 7)
	// B = A & 7
	// A = A >> 3  !!!!
	// B=(A0 & 7)^( A0 >> ( (A0 & 7) ^ 7))

	A := uint64(0)
	for k := uint64(0); k < 128; k++ {
		a := k
		r := calc(c, p, p, a, 0)
		if r > 0 {
			A = r
			break
		}

	}
	fmt.Printf("%d\n", A)
}

func get_combo_val(c *cpu, op int) uint64 {
	switch op {
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	default:
		return uint64(op)
	}
}

func exec(c *cpu, p []int) []int {
	c.ip = 0
	n := []int{}
	for c.ip < len(p) {
		//fmt.Printf("%v\n", c)
		op := p[c.ip]
		arg := p[c.ip+1]
		c.ip += 2
		switch op {
		case 0: //adv
			c.A >>= get_combo_val(c, arg)
		case 1: // bxl
			c.B ^= uint64(arg)
		case 2: //bst
			c.B = get_combo_val(c, arg) & 7
		case 3: //jnz
			if c.A != 0 {
				c.ip = arg
			}
		case 4: // bxc
			c.B ^= c.C
		case 5: // out
			n = append(n, int(get_combo_val(c, arg)&7))
		case 6: // bdv
			c.B = c.A >> get_combo_val(c, arg)
		case 7: // cdv
			c.C = c.A >> get_combo_val(c, arg)
		}
	}
	return n
}

func part1(c *cpu, p []int) {
	s := ""
	n := exec(c, p)
	for i, v := range n {
		if i != 0 {
			s += fmt.Sprintf(",%d", v)
		} else {
			s += fmt.Sprintf("%d", v)
		}
	}

	fmt.Printf("%s\n", s)
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

	p := []int{}
	c := cpu{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	l := scanner.Text()
	fmt.Sscanf(l, "Register A: %d", &c.A)
	scanner.Scan()
	l = scanner.Text()
	fmt.Sscanf(l, "Register B: %d", &c.B)
	scanner.Scan()
	l = scanner.Text()
	fmt.Sscanf(l, "Register C: %d", &c.C)
	scanner.Scan()
	scanner.Scan()
	l = scanner.Text()
	parts := strings.TrimPrefix(l, "Program: ")
	for _, s := range strings.Split(parts, ",") {
		if n, err := strconv.Atoi(s); err == nil {
			p = append(p, n)
		}
	}

	part1(&c, p)
	part2(&c, p)
}
