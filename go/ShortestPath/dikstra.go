package shortest_path

import "fmt"
import "container/heap"

func Dikstra() {
	var n, m int
	fmt.Scan(&n, &m)
	graph := make([][][]int, n)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		graph[u] = append(graph[u], []int{v, w})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[0] = 0

	done := make([]bool, n)

	for {
		selectNode := -1
		minDist := INF

		for i := 0; i < n; i++ {
			if !done[i] && dist[i] < minDist {
				minDist = dist[i]
				selectNode = i
			}
		}

		if selectNode == -1 {
			break
		}

		for _, child := range graph[selectNode] {
			dist[child[0]] = min(dist[child[0]], dist[selectNode]+child[1])
		}
		done[selectNode] = true
	}

	for _, v := range dist {
		fmt.Println(v)
	}

	fmt.Println(heap.Init)
}
