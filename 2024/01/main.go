package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func part2(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)
	list1 := []int{}
	list2 := map[int]int{}

	for scanner.Scan() {
		l := scanner.Text()
		var num1, num2 int
		fmt.Sscanf(l, "%d%d", &num1, &num2)
		list1 = append(list1, num1)
		if v, ok := list2[num2]; ok {
			list2[num2] = v + 1
		} else {
			list2[num2] = 1
		}
	}
	for i := range list1 {
		s += list1[i] * list2[list1[i]]
	}
	fmt.Printf("%d\n", s)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func part1(r io.Reader) {
	s := 0
	scanner := bufio.NewScanner(r)
	list1 := []int{}
	list2 := []int{}

	for scanner.Scan() {
		l := scanner.Text()
		var num1, num2 int
		fmt.Sscanf(l, "%d%d", &num1, &num2)
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	sort.Ints(list1)
	sort.Ints(list2)

	for i := range list1 {
		s += abs(list1[i] - list2[i])
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
