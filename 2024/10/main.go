package main

import (
	"bufio"
	"fmt"
	"os"
)

func part2(nums [][]int) {
	s := 0

	for y := 0; y < len(nums); y++ {
		for x := 0; x < len(nums[y]); x++ {
			s += search2(nums, x, y, 0)
		}
	}
	fmt.Printf("%d\n", s)
}

func search2(nums [][]int, x, y int, target int) int {
	if x < 0 || x >= len(nums[0]) || y < 0 || y >= len(nums) {
		return 0
	}
	if nums[y][x] != target {
		return 0
	}
	if nums[y][x] == 9 {
		return 1
	}
	return search2(nums, x+1, y, target+1) +
		search2(nums, x-1, y, target+1) +
		search2(nums, x, y+1, target+1) +
		search2(nums, x, y-1, target+1)
}

func search(nums [][]int, x, y int, target int, targets map[string]bool) {
	if x < 0 || x >= len(nums[0]) || y < 0 || y >= len(nums) {
		return
	}
	if nums[y][x] != target {
		return
	}
	if nums[y][x] == 9 {
		targets[fmt.Sprintf("%d-%d", y, x)] = true
		return
	}
	search(nums, x+1, y, target+1, targets)
	search(nums, x-1, y, target+1, targets)
	search(nums, x, y+1, target+1, targets)
	search(nums, x, y-1, target+1, targets)
}

func part1(nums [][]int) {
	s := 0

	for y := 0; y < len(nums); y++ {
		for x := 0; x < len(nums[y]); x++ {
			t := map[string]bool{}
			search(nums, x, y, 0, t)
			s += len(t)
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

	n := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()

		nums := []int{}
		for _, c := range l {
			nums = append(nums, int(c)-'0')
		}
		n = append(n, nums)
	}

	part1(n)
	part2(n)
}
