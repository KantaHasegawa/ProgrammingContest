package dfs

import (
	"fmt"
	"sort"
)

func FindCycle() {
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
	finish := make([]bool, n)

	isCycle := true
	var rec func(current int)
	rec = func(current int) {
		seen[current] = true

		for _, child := range graph[current] {
			if seen[child] {
				if !finish[child] {
					isCycle = false
				}
				continue
			}
			rec(child)
		}
		finish[current] = true
	}

	ans := "No"
	for i := 0; i < n; i++ {
		if seen[i] {
			continue
		}
		rec(i)
		if !isCycle {
			ans = "Yes"
		}
	}

	fmt.Println(ans)
}
