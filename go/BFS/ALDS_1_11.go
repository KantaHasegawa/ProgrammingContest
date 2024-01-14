package bfs

import "fmt"

// https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_11_C&lang=ja
func Alds_1_11() {
	var n int
	fmt.Scan(&n)
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		var u, k int
		fmt.Scan(&u, &k)
		for i := 0; i < k; i++ {
			var v int
			fmt.Scan(&v)
			graph[u-1] = append(graph[u-1], v)
		}
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0

	flg := make([]bool, n)
	flg[0] = true

	que := make([]int, 0, n)
	que = append(que, 0)

	for len(que) > 0 {
		front := que[0]
		que = que[1:]

		for _, v := range graph[front] {
			if flg[v-1] {
				continue
			}
			flg[v-1] = true
			dist[v-1] = dist[front] + 1
			que = append(que, v-1)
		}
	}

	for i, v := range dist {
		fmt.Println(i+1, v)
	}
}
