/*
*   Author: @v1gn35h
 */

package ds

import (
	"fmt"
	"math"
)

type Graph struct {
	v, e  int
	edges []edge
}

type edge struct {
	u, v, w int
}

// BellMan Ford Algorithm
func (grph *Graph) BellmanFord() {

	//Initialize distance for each vertice
	dis := make([]int, grph.v)
	for i := 0; i < grph.v; i++ {
		dis[i] = math.MaxInt
	}
	dis[0] = 0

	// Find distance for each vertices
	for i := 1; i < grph.v; i++ {
		for j := 0; j < grph.e; j++ {
			u := grph.edges[j].u
			v := grph.edges[j].v
			w := grph.edges[j].w

			if (dis[u] != math.MaxInt) && ((dis[u] + w) < dis[v]) {
				dis[v] = dis[u] + w
			}
		}
	}

	for i, v := range dis {
		fmt.Println(i, " => ", v)
	}

}

func (grph *Graph) dijkstra() {

	sp_set := make(map[int]bool)
	dist := make([]int, grph.v)

	for i := 0; i < grph.v; i++ {
		dist[i] = math.MaxInt
		sp_set[i] = false
	}

	dist[0] = 0

	for i := 0; i < grph.v; i++ {

		u := grph.min_val(dist, sp_set)

		if sp_set[u] == false {
			sp_set[u] = true
		}

		for j := 1; j < grph.v; j++ {
			if dist[u] != math.MaxInt && grph.find_weight(u, j) != -1 && ((dist[u] + grph.find_weight(u, j)) < dist[j]) {
				dist[j] = dist[u] + grph.find_weight(u, j)
			}
		}

	}

	for i, v := range dist {
		fmt.Println(i, " => ", v)
	}
}

func (grph *Graph) find_weight(u, v int) int {
	weight := -1

	for _, val := range grph.edges {
		if val.u == u && val.v == v {
			weight = val.w
		}
	}

	return weight
}

func (grph *Graph) min_val(dist []int, sp_set map[int]bool) int {
	min := math.MaxInt
	var min_index int

	for i := 0; i < len(dist); i++ {
		if sp_set[i] == false && dist[i] < min {
			min_index = i
		}
	}

	return min_index
}

func (grph *Graph) bfs(graph map[int][]int) {

	queue := make([]int, 0)
	visited := make(map[int]bool)

	queue = append(queue, 0)

	for len(queue) > 0 {
		x := queue[0]

		if visited[x] == false {
			visited[x] = true
			fmt.Printf("%d ", x)
		}

		for _, i := range graph[x] {
			if visited[i] == false {
				queue = append(queue, i)
			}
		}

		queue = queue[1:]
	}

}

func (grph *Graph) dfs(graph map[int][]int, start int, visited map[int]bool) {

	if visited[start] == false {
		fmt.Println(start)
		visited[start] = true
	} else {
		return
	}

	for _, v := range graph[start] {
		grph.dfs(graph, v, visited)
	}

}

func (grph *Graph) is_cyclic(graph map[int][]int, start, parent int, visited map[int]bool) bool {

	if visited[start] == false {
		visited[start] = true
	} else if start != parent {
		return true
	}

	for _, v := range graph[start] {
		grph.is_cyclic(graph, v, start, visited)
	}

	return false
}

func (grph *Graph) topological_sort(graph map[int][]int, start int, visited map[int]bool, stack []int) {

	if visited[start] == false {
		visited[start] = true
	} else {
		return
	}

	for _, v := range graph[start] {
		grph.topological_sort(graph, v, visited, stack)
	}

	stack = append(stack, start)
}

func (grph *Graph) prims_mst() {

	return

	// parent := make([]int, grph.v)
	// pqueue := make(priorityQueue.IntHeap, 0)
	// heap.Init(&pqueue)
	// IntHeap.len()

	// pqueue.Push(0)

	// for len(pqueue) > 0 {
	// 	x := pqueue.Pop()
	// }

}

func (graph *Graph) Legacy() {
	return

	//bfs(graph)
	//dfs(graph)
	// if is_cyclic(graph, 0, -1, make(map[int]bool)) == true {
	// 	fmt.Println("Yes")
	// } else {
	// 	fmt.Println("NO")
	// }
	//graph := gen_weighted_graph()
	//graph.bellman_ford()
	//graph.dijkstra()
}

func GetGraphArrayRepUMap() map[int][]int {
	vgraph := make(map[int][]int)

	for i := 0; i < 5; i++ {
		edges := make([]int, 0)
		for j := 0; j < 5; j++ {
			if j != i {
				edges = append(edges, j)
			}
		}
		vgraph[i] = edges
	}
	return vgraph
}

func GetGraphArrayRep(rmap map[int]bool) [5][5]int {
	var vgraph [5][5]int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if rmap[(i*10)+j] == true {
				vgraph[i][j] = 1
			}
		}
	}
	return vgraph
}

// Generates a random Graph data structure
func GetGraphListRep() *Graph {
	graph := &Graph{5, 8, make([]edge, 0)}

	// add edge 0-1 (or A-B in above figure)
	graph.edges = append(graph.edges, edge{0, 1, 5})

	// add edge 0-2 (or A-C in above figure)
	//graph.edges = append(graph.edges, edge{0, 2, 4})

	// add edge 1-2 (or B-C in above figure)
	graph.edges = append(graph.edges, edge{1, 2, 3})

	// add edge 1-3 (or B-D in above figure)
	graph.edges = append(graph.edges, edge{1, 3, 2})

	// add edge 1-4 (or B-E in above figure)
	graph.edges = append(graph.edges, edge{1, 4, 2})

	// add edge 3-2 (or D-C in above figure)
	//graph.edges = append(graph.edges, edge{3, 2, 5})

	// add edge 3-1 (or D-B in above figure)
	//graph.edges = append(graph.edges, edge{3, 1, 1})

	// add edge 4-3 (or E-D in above figure)
	//graph.edges = append(graph.edges, edge{4, 3, 3})

	return graph
}

// Connected components
func PrintAllConnectedComponents(graph [5][5]int, v int) {
	visited := make([]bool, v)
	for i := 0; i < v; i++ {
		fmt.Println("Visited:", visited, i)
		if visited[i] != true {
			dfs(graph, visited, i, v)
		}
		fmt.Println(" ")
	}

}

func dfs(graph [5][5]int, visited []bool, x, v int) {
	visited[x] = true
	fmt.Printf("%d ", x)
	for i := 0; i < v; i++ {
		if graph[x][i] == 1 && visited[i] != true {
			dfs(graph, visited, i, v)
		}
	}
}

func (grph *Graph) UnionsetCyclic() bool {
	uset := make([]int, grph.e)

	for i := 0; i < grph.v; i++ {
		uset[i] = i
	}

	for _, v := range grph.edges {
		x := findSet(uset, v.u)
		y := findSet(uset, v.v)

		if x == y {
			return true
		}

		uset[x] = y
	}

	return false
}

func findSet(set []int, v int) int {
	if set[v] == v {
		return v
	}

	return findSet(set, set[v])
}
