package dfs

import (
	"fmt"
	"sort"
)

// https://algo-method.com/tasks/969
func ExistPath() {
	var n, m, s, t int
	fmt.Scan(&n, &m, &s, &t)

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
	history := make([]int, n)
	for i := range history {
		history[i] = -1
	}

	var rec func([][]int, int)
	rec = func(graph [][]int, current int) {
		seen[current] = true
		for _, child := range graph[current] {
			if seen[child] {
				continue
			}
			history[child] = current
			rec(graph, child)
		}
	}
	rec(graph, s)

	ans := make([]int, 0, n)
	current := t
	for current != -1 {
		ans = append(ans, current)
		current = history[current]
	}

	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	fmt.Println(len(ans))
	for i, v := range ans {
		fmt.Print(v)
		if len(ans)-1 != i {
			fmt.Print(" ")
		} else {
			fmt.Println("")
		}
	}
}
