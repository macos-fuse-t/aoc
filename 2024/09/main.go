package main

import (
	"bufio"
	"fmt"
	"os"
)

type space struct {
	id    int
	start int
	len   int
}

func part2(l []int) {
	s := uint64(0)
	free := []space{}
	files := []space{}

	b := 0
	for i := 0; i < len(l); i += 2 {
		files = append(files, space{id: i / 2, start: b, len: l[i]})
		b += l[i]
		if i == len(l)-1 {
			break
		}
		free = append(free, space{start: b, len: l[i+1]})
		b += l[i+1]
	}

	for i := len(files) - 1; i >= 0; i-- {
		file := &files[i]
		for j := range free {
			fr := &free[j]
			if file.start <= fr.start {
				break
			}
			if fr.len >= file.len {
				file.start = fr.start
				fr.len -= file.len
				fr.start += file.len
			}
		}
	}
	for _, file := range files {
		for i := 0; i < file.len; i++ {
			s += uint64(file.id) * uint64(file.start+i)
		}
	}

	fmt.Printf("%d\n", s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func part1(nums []int) {
	l := make([]int, len(nums))
	copy(l, nums)
	s := uint64(0)
	pos1 := 0
	pos2 := len(l) - 1
	blocks := []int{}

	for pos1 < pos2 {
		b := pos1 / 2
		blocks = append(blocks, b, l[pos1])

		for l[pos1+1] > 0 && pos1 < pos2 {
			free := l[pos1+1]
			move := l[pos2]
			b = pos2 / 2
			blocks = append(blocks, b, min(free, move))
			l[pos2] -= min(free, move)
			l[pos1+1] -= min(free, move)
			if l[pos2] == 0 {
				pos2 -= 2
			}
		}
		pos1 += 2
	}
	for l[pos1] > 0 {
		b := pos1 / 2
		blocks = append(blocks, b, l[pos1])
		pos1 += 2
	}
	b := 0
	for i := 0; i < len(blocks)-1; i += 2 {
		for j := 0; j < blocks[i+1]; j++ {
			s += uint64(b) * uint64(blocks[i])
			b++
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

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	l := scanner.Text()
	nums := make([]int, len(l))
	for i, char := range l {
		nums[i] = int(char - '0')
	}

	part1(nums)
	part2(nums)
}
