package dfs

import (
	"fmt"
	"sort"
)

// https://algo-method.com/tasks/963
func TopologicalSort() {
	var n, m int
	fmt.Scan(&n, &m)

	graph := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		graph[a] = append(graph[a], b)
	}
	for i := range graph {
		sort.Slice(graph[i], func(ii, jj int) bool {
			return graph[i][ii] < graph[i][jj]
		})
	}

	seen := make([]bool, n)

	order := make([]int, 0, n)
	var rec func(graph [][]int, seen []bool, current int)
	rec = func(graph [][]int, seen []bool, current int) {
		seen[current] = true
		for _, child := range graph[current] {
			if seen[child] {
				continue
			}
			rec(graph, seen, child)
		}
		order = append(order, current)
	}

	for i := 0; i < n; i++ {
		if seen[i] {
			continue
		}
		rec(graph, seen, i)
	}

	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}

	for _, v := range order {
		fmt.Print(v, " ")
	}
}
