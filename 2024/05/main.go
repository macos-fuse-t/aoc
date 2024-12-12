package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type DepTree map[int][]int

func part2(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)

	deps := make(DepTree)
	state := 0
	for scanner.Scan() {
		l := scanner.Text()
		if state == 0 {
			if l == "" || l == "\n" {
				state = 1
				continue
			}
			var num1, num2 int
			fmt.Sscanf(l, "%d|%d", &num1, &num2)
			if v, ok := deps[num1]; ok {
				v = append(v, num2)
				deps[num1] = v
			} else {
				deps[num1] = []int{num2}
			}
		} else {
			nums := []int{}
			valid := true
			for _, part := range strings.Split(l, ",") {
				num, _ := strconv.Atoi(part)
				for _, n := range nums {
					if contains(deps[num], n) {
						valid = false
						break
					}
				}
				nums = append(nums, num)
			}
			if !valid {
				changed := true
				for changed {
					changed = false
					for j := range nums {
						num := nums[j]
						for i := j + 1; i < len(nums); i++ {
							n := nums[i]
							if contains(deps[num], n) {
								nums[i], nums[j] = nums[j], nums[i]
								changed = true
								break
							}
						}
					}
				}
				s += nums[len(nums)/2]
			}
		}
	}
	fmt.Printf("%d\n", s)
}

func contains(nums []int, n int) bool {
	for _, v := range nums {
		if v == n {
			return true
		}
	}
	return false
}

func part1(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)

	deps := make(DepTree)
	state := 0
	for scanner.Scan() {
		l := scanner.Text()
		if state == 0 {
			if l == "" || l == "\n" {
				state = 1
				continue
			}
			var num1, num2 int
			fmt.Sscanf(l, "%d|%d", &num1, &num2)
			if v, ok := deps[num1]; ok {
				v = append(v, num2)
				deps[num1] = v
			} else {
				deps[num1] = []int{num2}
			}
		} else {
			nums := []int{}
			valid := true
			for _, part := range strings.Split(l, ",") {
				num, _ := strconv.Atoi(part)
				for _, n := range nums {
					if contains(deps[num], n) {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
				nums = append(nums, num)
			}
			if valid {
				s += nums[len(nums)/2]
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

	part1(file)
	file.Seek(0, 0)
	part2(file)
}
