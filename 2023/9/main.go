package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(nums [][]int) {
	s := 0
	for _, seq := range nums {
		stop := false

		for !stop {
			stop = true
			s += seq[len(seq)-1]
			newSeq := make([]int, len(seq)-1)
			for i := 1; i < len(seq); i++ {
				newSeq[i-1] = seq[i] - seq[i-1]
				if newSeq[i-1] != 0 {
					stop = false
				}
			}
			seq = newSeq
		}
	}

	fmt.Printf("%d\n", s)
}

func part2(nums [][]int) {
	for _, seq := range nums {
		slices.Reverse(seq)
	}
	part1(nums)
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

	nums := [][]int{}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		nums = append(nums, []int{})
		split := strings.Split(scanner.Text(), " ")
		for _, s := range split {
			n, _ := strconv.Atoi(s)
			nums[i] = append(nums[i], n)
		}
		i++
	}

	part1(nums)
	part2(nums)
}
