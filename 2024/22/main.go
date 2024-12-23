package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	children [19]*Node
	value    *int
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{}}
}

func offset(d int) int {
	return d + 9
}

func (t *Trie) Insert(seq []int, leafValue int) {
	cur := t.root
	for _, digit := range seq {
		idx := offset(digit)
		if cur.children[idx] == nil {
			cur.children[idx] = &Node{}
		}
		cur = cur.children[idx]
	}
	cur.value = &leafValue
}

func (t *Trie) Search(seq []int) (int, bool) {
	cur := t.root
	for _, digit := range seq {
		idx := offset(digit)
		if cur.children[idx] == nil {
			return 0, false
		}
		cur = cur.children[idx]
	}
	if cur.value != nil {
		return *cur.value, true
	}
	return 0, false
}

func part2(nums []int) {
	s := 0
	t := NewTrie()
	for _, n := range nums {
		l := calc(n, 2000)
		t1 := NewTrie()
		for i := 1; i < 1997; i++ {
			k := []int{l[i]%10 - l[i-1]%10, l[i+1]%10 - l[i]%10, l[i+2]%10 - l[i+1]%10, l[i+3]%10 - l[i+2]%10}
			v := l[i+3] % 10

			if _, ok := t1.Search(k); ok {
				continue
			}
			t1.Insert(k, v)

			oldv, ok := t.Search(k)
			if !ok {
				t.Insert(k, v)
				if v > s {
					s = v
				}
			} else {
				t.Insert(k, v+oldv)
				if v+oldv > s {
					s = v + oldv
				}
			}
		}
	}

	fmt.Printf("%d\n", s)
}

func calc(n int, i int) []int {
	l := []int{}
	for i > 0 {
		n = ((n * 64) ^ n) % 16777216
		n = ((n / 32) ^ n) % 16777216
		n = ((n * 2048) ^ n) % 16777216
		l = append(l, n)
		i--
	}
	return l
}

func part1(nums []int) {
	s := 0
	for _, n := range nums {
		l := calc(n, 2000)
		s += l[len(l)-1]
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

	nums := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		n, _ := strconv.Atoi(l)
		nums = append(nums, n)
	}

	part1(nums)
	part2(nums)
}
