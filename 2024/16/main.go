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
	dir  byte
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

		if current.x == end.x && current.y == end.y && (end.dir == '0' || end.dir == current.dir) {
			end.dir = current.dir
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

// for part2 it's better to use floyd marshall alg, dijkstra takes minutes to complete
func part2(grid []string, g graph, s, e vertex) {
	d, _ := part1(grid, g, s, e)
	m := map[string]bool{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '#' {
				continue
			}
			for _, dir := range []byte{'E', 'W', 'S', 'N'} {
				d1, v := part1(grid, g, s, vertex{x, y, dir})
				d2, _ := part1(grid, g, v, e)
				if d1+d2 == d {
					m[fmt.Sprintf("%d,%d", x, y)] = true
					break
				}
				if d1+d2 > d+4000 {
					break
				}

			}
		}
	}
	fmt.Printf("%d\n", len(m))
}

func part1(grid []string, g graph, s, e vertex) (int, vertex) {
	d, p := dijkstra(g, s, e)
	return d, p[len(p)-1]
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

	grid := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		grid = append(grid, l)
	}

	startx, starty, endx, endy := 0, 0, 0, 0
	g := graph{v: make([]vertex, 0), e: make(map[vertex][]edge)}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '#' {
				continue
			}
			if grid[y][x] == 'S' {
				startx, starty = x, y
			}
			if grid[y][x] == 'E' {
				endx, endy = x, y
			}

			m := map[byte][]byte{'E': {'N', 'S'}, 'S': {'E', 'W'}, 'W': {'N', 'S'}, 'N': {'E', 'W'}}
			for d := range m {
				v := vertex{x, y, d}
				g.v = append(g.v, v)

				if d == 'W' && x > 0 && grid[y][x-1] != '#' {
					v2 := vertex{x - 1, y, 'W'}
					g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
				} else if d == 'N' && y > 0 && grid[y-1][x] != '#' {
					v2 := vertex{x, y - 1, 'N'}
					g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
				} else if d == 'E' && x < len(grid[0])-1 && grid[y][x+1] != '#' {
					v2 := vertex{x + 1, y, 'E'}
					g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
				} else if d == 'S' && y < len(grid)-1 && grid[y+1][x] != '#' {
					v2 := vertex{x, y + 1, 'S'}
					g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
				}

				for _, d2 := range m[d] {
					v2 := vertex{x, y, d2}
					g.e[v] = append(g.e[v], edge{to: v2, weight: 1000})
				}
			}
		}
	}

	d, _ := part1(grid, g, vertex{startx, starty, 'E'}, vertex{endx, endy, '0'})
	fmt.Printf("%d\n", d)
	part2(grid, g, vertex{startx, starty, 'E'}, vertex{endx, endy, '0'})
}
