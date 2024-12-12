package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = map[string]int{}

func transform1(n int) []int {
	nums := []int{}
	s := strconv.Itoa(n)
	if n == 0 {
		nums = append(nums, 1)
	} else if len(s)%2 == 0 {
		n1, _ := strconv.Atoi(s[:len(s)/2])
		n2, _ := strconv.Atoi(s[len(s)/2:])
		nums = append(nums, n1, n2)
	} else {
		nums = append(nums, n*2024)
	}

	return nums
}

func transform(n int, cnt int) int {
	s := 0
	if cnt == 0 {
		return 1
	}
	c := fmt.Sprintf("%d-%d", cnt, n)
	if v, ok := cache[c]; ok {
		return v
	}
	l := transform1(n)
	for _, n1 := range l {
		s += transform(n1, cnt-1)
	}
	cache[c] = s
	return s
}

func part1(nums []int) {
	s := 0
	for _, n := range nums {
		s += transform(n, 25)
	}

	fmt.Printf("%d\n", s)
}

func part2(nums []int) {
	s := 0
	for _, n := range nums {
		s += transform(n, 75)
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
	nums := []int{}

	for scanner.Scan() {
		l := scanner.Text()
		parts := strings.Split(l, " ")
		for _, s := range parts {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
	}

	part1(nums)
	part2(nums)
}
