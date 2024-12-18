package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type vertex struct {
	x, y int
}

type edge struct {
	to     vertex
	weight int
}

type graph struct {
	v []vertex
	e map[vertex][]edge
}

type item struct {
	v        vertex
	distance int
	index    int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	it := x.(*item)
	it.index = n
	*pq = append(*pq, it)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	it := old[n-1]
	old[n-1] = nil
	it.index = -1
	*pq = old[0 : n-1]
	return it
}

func dijkstra(g graph, start, end vertex) (int, []vertex) {
	dist := make(map[vertex]int)
	for _, v := range g.v {
		dist[v] = math.MaxInt
	}
	dist[start] = 0

	prev := make(map[vertex]*vertex)

	pq := make(priorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &item{v: start, distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*item).v

		if current == end {
			break
		}

		for _, edge := range g.e[current] {
			alt := dist[current] + edge.weight
			if alt < dist[edge.to] {
				dist[edge.to] = alt
				prev[edge.to] = &current
				heap.Push(&pq, &item{v: edge.to, distance: alt})
			}
		}
	}

	path := []vertex{}
	for at := &end; at != nil; at = prev[*at] {
		path = append([]vertex{*at}, path...)
	}

	if len(path) == 0 || path[0] != start {
		return -1, nil
	}

	return dist[end], path
}

func part2(m map[string]int) {
	s := 0
	min_steps := 0
	max_steps := len(m) - 1
	for min_steps < max_steps-1 {
		cur_steps := (min_steps + max_steps) / 2
		g := make_graph(m, cur_steps)
		s, _ = dijkstra(g, vertex{0, 0}, vertex{w - 1, h - 1})
		if s != -1 {
			min_steps = cur_steps
		} else {
			max_steps = cur_steps
		}
	}

	for k, v := range m {
		if v == min_steps {
			fmt.Printf("%s, %d\n", k, v)
		}
	}
}

func part1(m map[string]int) {
	s := 0
	g := make_graph(m, steps)
	s, _ = dijkstra(g, vertex{0, 0}, vertex{w - 1, h - 1})
	fmt.Printf("%d\n", s)
}

func make_graph(m map[string]int, maxsteps int) graph {
	g := graph{v: make([]vertex, 0), e: make(map[vertex][]edge)}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if s, ok := m[fmt.Sprintf("%d,%d", x, y)]; ok && s < maxsteps {
				continue
			}
			v := vertex{x, y}
			g.v = append(g.v, v)
			if s, ok := m[fmt.Sprintf("%d,%d", x-1, y)]; x > 0 && (!ok || s >= maxsteps) {
				v2 := vertex{x - 1, y}
				g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
			}
			if s, ok := m[fmt.Sprintf("%d,%d", x, y-1)]; y > 0 && (!ok || s >= maxsteps) {
				v2 := vertex{x, y - 1}
				g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
			}
			if s, ok := m[fmt.Sprintf("%d,%d", x+1, y)]; x < w-1 && (!ok || s >= maxsteps) {
				v2 := vertex{x + 1, y}
				g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
			}
			if s, ok := m[fmt.Sprintf("%d,%d", x, y+1)]; y < h-1 && (!ok || s >= maxsteps) {
				v2 := vertex{x, y + 1}
				g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
			}
		}
	}
	return g
}

var w = 71
var h = 71
var steps = 1024

//var w = 7
//var h = 7
//var steps = 12

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

	i := 0
	m := map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		var x, y int
		fmt.Sscanf(l, "%d,%d", &x, &y)
		m[fmt.Sprintf("%d,%d", x, y)] = i
		i++
	}

	part1(m)
	part2(m)
}
