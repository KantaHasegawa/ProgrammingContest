package bfs

import (
	"fmt"
	"sort"
)

func TopologicalSort() {
	var n, m int
	fmt.Scan(&n, &m)
	deg := make([]int, n)
	graph := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		deg[a]++
		graph[b] = append(graph[b], a)
	}

	for i := range graph {
		sort.Slice(graph[i], func(ii, jj int) bool {
			return graph[i][ii] < graph[i][jj]
		})
	}

	order := make([]int, 0, n)
	que := make([]int, 0, n)
	for i := range deg {
		if deg[i] == 0 {
			que = append(que, i)
		}
	}

	for len(que) > 0 {
		front := que[0]
		que = que[1:]
		order = append(order, front)
		for _, child := range graph[front] {
			deg[child]--
			if deg[child] == 0 {
				que = append(que, child)
			}
		}
	}

	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1{
		order[i], order[j] = order[j], order[i]
	}

	for _, v :=  range order{
		fmt.Print(v, " ")
	}
}
