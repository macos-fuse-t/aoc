package main

import (
	"bufio"
	"fmt"
	"os"
)

func part2(eqs []eq) {
	s := 0
	for _, eq := range eqs {
		eq.c0 += 10000000000000
		eq.c1 += 10000000000000
		D := eq.a0*eq.b1 - eq.b0*eq.a1
		Dx := eq.c0*eq.b1 - eq.b0*eq.c1
		Dy := eq.a0*eq.c1 - eq.c0*eq.a1
		if D == 0 {
			continue
		}
		if Dx%D != 0 || Dy%D != 0 {
			continue
		}
		s += 3*Dx/D + 1*Dy/D
	}
	fmt.Printf("%d\n", s)
}

type eq struct {
	a0, a1, b0, b1, c0, c1 int
}

func part1(eqs []eq) {
	s := 0

	for _, eq := range eqs {
		D := eq.a0*eq.b1 - eq.b0*eq.a1
		Dx := eq.c0*eq.b1 - eq.b0*eq.c1
		Dy := eq.a0*eq.c1 - eq.c0*eq.a1
		if D == 0 {
			continue
		}
		if Dx%D != 0 || Dy%D != 0 {
			continue
		}
		s += 3*Dx/D + 1*Dy/D
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

	scanner := bufio.NewScanner(file)
	eqs := []eq{}

	for scanner.Scan() {
		var a0, a1, b0, b1, c0, c1 int

		l := scanner.Text()
		fmt.Sscanf(l, "Button A: X+%d, Y+%d", &a0, &a1)
		scanner.Scan()
		l = scanner.Text()
		fmt.Sscanf(l, "Button B: X+%d, Y+%d", &b0, &b1)
		scanner.Scan()
		l = scanner.Text()
		fmt.Sscanf(l, "Prize: X=%d, Y=%d", &c0, &c1)
		scanner.Scan()
		scanner.Text()
		eqs = append(eqs, eq{a0, a1, b0, b1, c0, c1})
	}

	part1(eqs)
	part2(eqs)
}
