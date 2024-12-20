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

func dijkstra(g graph, start, end vertex, maxdist int) (int, []vertex, string) {
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
				if maxdist > 0 && alt > maxdist {
					return -1, nil, ""
				}
				dist[edge.to] = alt
				prev[edge.to] = &current
				heap.Push(&pq, &item{v: edge.to, distance: alt})
			}
		}
	}

	strpath := ""
	path := []vertex{}
	for at := &end; at != nil; at = prev[*at] {
		path = append([]vertex{*at}, path...)
		strpath = fmt.Sprintf("(%d,%d)", at.x, at.y) + strpath
	}

	if len(path) == 0 || path[0] != start {
		return -1, nil, ""
	}

	return dist[end], path, strpath
}

func calcDist(g graph, s, e vertex) map[vertex]map[vertex]int {
	dist := make(map[vertex]map[vertex]int)

	for i, u := range g.v {
		if i%100 == 0 {
			fmt.Printf("i: %d from %d\n", i, len(g.v))
		}
		d, _, _ := dijkstra(g, s, u, 0)
		if dist[s] == nil {
			dist[s] = make(map[vertex]int)
		}
		dist[s][u] = d
		d, _, _ = dijkstra(g, u, e, 0)
		if dist[u] == nil {
			dist[u] = make(map[vertex]int)
		}
		dist[u][e] = d
	}
	return dist
}

func makeGraph(grid []string) (graph, vertex, vertex) {
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

			v := vertex{x, y}
			g.v = append(g.v, v)

			if x > 0 && grid[y][x-1] != '#' {
				v2 := vertex{x - 1, y}
				g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
			}
			if y > 0 && grid[y-1][x] != '#' {
				v2 := vertex{x, y - 1}
				g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
			}
			if x < len(grid[0])-1 && grid[y][x+1] != '#' {
				v2 := vertex{x + 1, y}
				g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
			}
			if y < len(grid)-1 && grid[y+1][x] != '#' {
				v2 := vertex{x, y + 1}
				g.e[v] = append(g.e[v], edge{to: v2, weight: 1})
			}
		}
	}
	return g, vertex{startx, starty}, vertex{endx, endy}
}

func mod(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func part2(grid []string) {
	g, s, e := makeGraph(grid)
	d, _, _ := dijkstra(g, s, e, 0)

	dist := calcDist(g, s, e)
	n := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '#' {
				continue
			}

			for y1 := y - 20; y1 <= y+20; y1++ {
				for x1 := x - 20 + mod(y-y1); x1 <= x+20-mod(y-y1); x1++ {
					if x1 < 0 || x1 >= len(grid[y]) || y1 < 0 || y1 >= len(grid) {
						continue
					}
					l := mod(x-x1) + mod(y-y1)
					if l > 20 {
						continue
					}
					if grid[y1][x1] == '#' {
						continue
					}

					v1 := vertex{x, y}
					v2 := vertex{x1, y1}
					g.e[v1] = append(g.e[v1], edge{to: v2, weight: l})
					g.e[v2] = append(g.e[v2], edge{to: v1, weight: l})

					if dist[s][v1] >= 0 && dist[v2][e] >= 0 {
						d2 := dist[s][v1] + dist[v2][e] + l
						if d-d2 >= 100 {
							n++
						}
					}

					g.e[v1] = g.e[v1][:len(g.e[v1])-1]
					g.e[v2] = g.e[v2][:len(g.e[v2])-1]
				}
			}
		}
	}

	fmt.Printf("%d\n", n)
}

func part1(grid []string) {
	g, s, e := makeGraph(grid)
	d, _, p := dijkstra(g, s, e, 0)
	paths := map[string]bool{p: true}

	for y := 0; y < len(grid); y++ {
		fmt.Printf("y: %d\n", y)
		for x := 0; x < len(grid[y]); x++ {
			grid2 := make([]string, len(grid))
			copy(grid2, grid)

			grid2[y] = grid[y][:x] + "." + grid[y][x+1:]
			g2, _, _ := makeGraph(grid2)
			d2, _, p2 := dijkstra(g2, s, e, d-100)
			if d2 > 0 && d2 < d && !paths[p2] && d-d2 >= 100 {
				paths[p2] = true
			}
		}
	}

	fmt.Printf("%d\n", len(paths)-1)
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

	part1(grid)
	part2(grid)
}
