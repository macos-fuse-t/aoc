package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Line struct{ p1, p2 int }
type Brick struct{ lx, ly, lz Line }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lineIntersecting(l1, l2 Line) bool {
	return !(l1.p1 > l2.p2 || l1.p2 < l2.p1)
}

func brickIntersecting(c1, c2 Brick) bool {
	return lineIntersecting(c1.lx, c2.lx) &&
		lineIntersecting(c1.ly, c2.ly) &&
		lineIntersecting(c1.lz, c2.lz)
}

func fell(bricks []Brick) (int, []Brick) {
	n := 0
	fallen := []Brick{}
	for _, b := range bricks {
		b1 := b
		for b1.lz.p1 > 1 {
			b1.lz.p1--
			b1.lz.p2--
			for _, b2 := range fallen {
				if brickIntersecting(b1, b2) {
					b1.lz.p1++
					b1.lz.p2++
					//fmt.Printf("%v intersecting %v\n", b1, b2)
					goto next
				}
			}
		}
	next:
		fallen = append(fallen, b1)
		if b1.lz.p1 < b.lz.p1 {
			n++
		}
	}

	//fmt.Printf("orig: %v\nnew %v\n", bricks, fallen)
	return n, fallen
}

func remove(slice []Brick, i int) []Brick {
	a := make([]Brick, 0)
	a = append(a, slice[:i]...)
	a = append(a, slice[i+1:]...)
	return a
}

func part1(bricks []Brick) {
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].lz.p1 < bricks[j].lz.p1
	})

	n := 0
	_, bricks = fell(bricks)
	for i := range bricks {
		bricks1 := remove(bricks, i)
		if n1, _ := fell(bricks1); n1 == 0 {
			n++
		}
	}
	fmt.Printf("%d\n", n)
}

func part2(bricks []Brick) {
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].lz.p1 < bricks[j].lz.p1
	})

	n := 0
	_, bricks = fell(bricks)
	for i := range bricks {
		bricks1 := remove(bricks, i)
		n1, _ := fell(bricks1)
		n += n1
	}
	fmt.Printf("%d\n", n)
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
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
	bricks := []Brick{}
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), "~")
		a1 := strings.Split(a[0], ",")
		a2 := strings.Split(a[1], ",")

		bricks = append(bricks, Brick{
			lx: Line{atoi(a1[0]), atoi(a2[0])},
			ly: Line{atoi(a1[1]), atoi(a2[1])},
			lz: Line{atoi(a1[2]), atoi(a2[2])},
		})
	}

	part1(bricks)
	part2(bricks)
}
