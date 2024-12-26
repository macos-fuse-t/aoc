package main

import (
	"bufio"
	"fmt"
	"os"
)

func part2(nums [][]int) {
	s := 0
	fmt.Printf("%d\n", s)
}

func part1(locks [][]int, keys [][]int, h int) {
	s := 0
	for _, k := range keys {
		for _, l := range locks {
			fit := true
			for i := 0; i < len(k); i++ {
				if l[i]+k[i] > h {
					fit = false
					break
				}
			}
			if fit {
				s++
			}
		}
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

	keys := [][]int{}
	locks := [][]int{}

	txt := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			n := []int{}
			for x := 0; x < len(txt[0]); x++ {
				k := 0
				for y := 0; y < len(txt); y++ {
					if txt[y][x] == '#' {
						k++
					}
				}
				n = append(n, k-1)
			}
			if txt[0][0] == '#' {
				locks = append(locks, n)
			} else {
				keys = append(keys, n)
			}
			txt = []string{}
		} else {
			txt = append(txt, l)
		}
	}

	n := []int{}
	for x := 0; x < len(txt[0]); x++ {
		k := 0
		for y := 0; y < len(txt); y++ {
			if txt[y][x] == '#' {
				k++
			}
		}
		n = append(n, k-1)
	}
	if txt[0][0] == '#' {
		locks = append(locks, n)
	} else {
		keys = append(keys, n)
	}

	part1(locks, keys, len(txt)-2)
	//part2(n)
}
