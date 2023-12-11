package main

import (
	"bufio"
	"fmt"
	"os"
)

func addDirection(east bool, i int, j int, graph []string, dist map[string]int, neigbors map[string][]string) {
	dist[fmt.Sprintf("%d_%d-%d_%d", i, j, i, j)] = 0
	if east {
		if j < len(graph[i])-1 {
			if graph[i][j+1] == '-' || graph[i][j+1] == 'J' || graph[i][j+1] == '7' || graph[i][j+1] == 'S' {
				u := fmt.Sprintf("%d_%d", i, j)
				v := fmt.Sprintf("%d_%d", i, j+1)
				dist[fmt.Sprintf("%s-%s", u, v)] = 1
				dist[fmt.Sprintf("%s-%s", v, u)] = 1
				neigbors[u] = append(neigbors[u], v)
				neigbors[v] = append(neigbors[v], u)
			}
		}
	} else { // south
		if i < len(graph)-1 {
			if graph[i+1][j] == '|' || graph[i+1][j] == 'J' || graph[i+1][j] == 'L' || graph[i+1][j] == 'S' {
				u := fmt.Sprintf("%d_%d", i, j)
				v := fmt.Sprintf("%d_%d", i+1, j)
				dist[fmt.Sprintf("%s-%s", u, v)] = 1
				dist[fmt.Sprintf("%s-%s", v, u)] = 1
				neigbors[u] = append(neigbors[u], v)
				neigbors[v] = append(neigbors[v], u)
			}
		}
	}
}

func buildGraph(graph []string) (string, []string, map[string]int, map[string][]string) {
	edges := map[string]int{}
	vertices := []string{}
	neigbors := map[string][]string{}
	start := ""

	for i := range graph {
		for j := range graph[i] {
			v := fmt.Sprintf("%d_%d", i, j)
			if graph[i][j] != '.' {
				vertices = append(vertices, v)
			}
			switch graph[i][j] {
			case 'S':
				start = v
				addDirection(true, i, j, graph, edges, neigbors)
				addDirection(false, i, j, graph, edges, neigbors)
			case '|', '7':
				addDirection(false, i, j, graph, edges, neigbors)
			case '-', 'L':
				addDirection(true, i, j, graph, edges, neigbors)
			case 'F':
				addDirection(true, i, j, graph, edges, neigbors)
				addDirection(false, i, j, graph, edges, neigbors)
			}
		}
	}
	return start, vertices, edges, neigbors
}

func findLoop(start, current, prev string, neighbors map[string][]string, visited map[string]bool, path *[]string) bool {
	if start == current && len(*path) > 0 {
		return true
	}
	if visited[current] {
		return false
	}
	visited[current] = true

	for _, v := range neighbors[current] {
		if v == prev {
			continue
		}
		*path = append(*path, v)
		if findLoop(start, v, current, neighbors, visited, path) {
			return true
		}
		*path = (*path)[:len(*path)-1]
	}
	return false
}

func part1(graph []string) {
	start, _, _, neighbors := buildGraph(graph)
	visited := map[string]bool{}
	path := []string{}
	findLoop(start, start, start, neighbors, visited, &path)

	fmt.Printf("%d\n", len(path)/2)
}

func part2(graph []string) {
	s := 0
	start, vertices, edges, neighbors := buildGraph(graph)
	visited := map[string]bool{}
	path := []string{}
	findLoop(start, start, start, neighbors, visited, &path)

	pathMap := map[string]bool{}
	for _, p := range path {
		pathMap[p] = true
	}
	vertMap := map[string]bool{}
	for _, v := range vertices {
		vertMap[v] = true
	}

	for i := range graph {
		for j := range graph[i] {
			v := fmt.Sprintf("%d_%d", i, j)
			if pathMap[v] {
				continue
			}
			// line
			n := 0
			for k := j + 1; k < len(graph[i]); k++ {
				kv := fmt.Sprintf("%d_%d", i, k)
				kv_1 := fmt.Sprintf("%d_%d", i, k-1)
				kv_2 := fmt.Sprintf("%d_%d", i, k+1)
				kv_3 := fmt.Sprintf("%d_%d", i-1, k)

				if (edges[fmt.Sprintf("%s-%s", kv, kv_1)] == 1 || edges[fmt.Sprintf("%s-%s", kv, kv_2)] == 1) &&
					edges[fmt.Sprintf("%s-%s", kv, kv_3)] == 0 {
					continue
				}

				if pathMap[kv] {
					n++
				}
			}
			if n%2 == 1 {
				s++
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

	graph := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		graph = append(graph, scanner.Text())
	}

	part1(graph)
	part2(graph)
}
