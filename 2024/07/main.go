package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2(nums [][]int) {
	s := 0

	for _, n := range nums {
		if isValid2(n[2:], n[1], n[0]) {
			s += n[0]
		}
	}
	fmt.Printf("%d\n", s)
}

func isValid2(nums []int, res, target int) bool {
	if len(nums) == 0 {
		return res == target
	}

	v, _ := strconv.Atoi(fmt.Sprintf("%d%d", res, nums[0]))
	return isValid2(nums[1:], res+nums[0], target) ||
		isValid2(nums[1:], res*nums[0], target) ||
		isValid2(nums[1:], v, target)
}

func isValid(nums []int, res, target int) bool {
	if len(nums) == 0 {
		return res == target
	}
	return isValid(nums[1:], res+nums[0], target) || isValid(nums[1:], res*nums[0], target)
}

func part1(nums [][]int) {
	s := 0

	for _, n := range nums {
		if isValid(n[2:], n[1], n[0]) {
			s += n[0]
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
		parts := strings.SplitN(l, ":", 2)
		num, _ := strconv.Atoi(parts[0])
		nums = append(nums, num)

		fields := strings.Fields(strings.TrimSpace(parts[1]))
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			nums = append(nums, num)
		}
		n = append(n, nums)
	}

	part1(n)
	part2(n)
}
