package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func part2(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		l := scanner.Text()
		nums := []int{}

		fields := strings.Fields(strings.TrimSpace(l))
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			nums = append(nums, num)
		}

		for i := range nums {
			temp := make([]int, len(nums)-1)
			copy(temp[:i], nums[:i])
			copy(temp[i:], nums[i+1:])
			if isSafe(temp) {
				s += 1
				break
			}
		}
	}
	fmt.Printf("%d\n", s)
}

func isSafe(nums []int) bool {
	increasing := nums[1] > nums[0]
	safe := true
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		if increasing && d != 1 && d != 2 && d != 3 {
			safe = false
			break
		}
		if !increasing && d != -1 && d != -2 && d != -3 {
			safe = false
			break
		}
	}
	return safe
}

func part1(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		l := scanner.Text()
		nums := []int{}

		fields := strings.Fields(strings.TrimSpace(l))
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			nums = append(nums, num)
		}

		if isSafe(nums) {
			s += 1
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

	part1(file)
	file.Seek(0, 0)
	part2(file)
}
