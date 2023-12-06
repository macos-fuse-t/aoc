package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func part1(times []int64, dist []int64) {
	s := int64(1)

	for i := range times {
		t := times[i]
		d := dist[i]

		l1 := (float64(-t) + math.Sqrt(float64(t*t-4*d))) / -2
		l2 := (float64(-t) - math.Sqrt(float64(t*t-4*d))) / -2

		if math.Ceil(l1) > l1 {
			l1 = math.Ceil(l1)
		} else {
			l1 = math.Ceil(l1 + 1)
		}

		if math.Floor(l2) < l2 {
			l2 = math.Floor(l2)
		} else {
			l2 = math.Floor(l2 - 1)
		}

		s *= int64(l2) - int64(l1) + 1
	}

	fmt.Printf("%d\n", s)
}

func part2(times []int64, dist []int64) {
	st := ""
	sd := ""
	for i := range times {
		st += fmt.Sprint(times[i])
		sd += fmt.Sprint(dist[i])
	}
	t, _ := strconv.ParseInt(st, 10, 64)
	d, _ := strconv.ParseInt(sd, 10, 64)
	part1([]int64{t}, []int64{d})
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

	times := []int64{}
	dist := []int64{}

	re := regexp.MustCompile(`\d+`) // numbers
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	matches := re.FindAllString(scanner.Text(), -1)
	for _, m := range matches {
		v, _ := strconv.ParseInt(m, 10, 64)
		times = append(times, v)
	}
	scanner.Scan()
	matches = re.FindAllString(scanner.Text(), -1)
	for _, m := range matches {
		v, _ := strconv.ParseInt(m, 10, 64)
		dist = append(dist, v)
	}

	part1(times, dist)
	part2(times, dist)
}
