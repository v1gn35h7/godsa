package algo

import "fmt"

func cyclicSort(graph map[int][]int) {

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
