package bfs

import (
	"fmt"
	"sort"
)

func FindCycle() {
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

	que := make([]int, 0, n)
	for i := range deg {
		if deg[i] == 0 {
			que = append(que, i)
		}
	}

	order := make([]int, 0, n)
	for len(que) > 0 {
		front := que[0]
		que = que[1:]
		order = append(order, front)
		for _, v := range graph[front] {
			deg[v]--
			if deg[v] == 0 {
				que = append(que, v)
			}
		}
	}

	for i, j := 0, len(order)-1; i < j; i, j = i-1, j-1 {
		order[i], order[j] = order[j], order[i]
	}

	var isCycle bool
	for i := 0; i < n; i++ {
		var find bool
		for _, v := range order {
			if i == v {
				find = true
			}
		}
		if !find {
			isCycle = true
			break
		}
	}

	if isCycle {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
