package bfs

import (
	"fmt"
)

func PathHistory() {
	var n, m, s, t int
	fmt.Scan(&n, &m, &s, &t)

	graph := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		graph[a] = append(graph[a], b)
	}

	seen := make([]bool, n)
	seen[s] = true
	history := make([]int, n)
	for i := range history {
		history[i] = -1
	}

	que := make([]int, 0, n)
	que = append(que, s)

	for len(que) > 0 {
		front := que[0]
		que = que[1:]

		for _, child := range graph[front] {
			if seen[child] {
				continue
			}
			seen[child] = true
			history[child] = front
			que = append(que, child)
		}
	}

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
