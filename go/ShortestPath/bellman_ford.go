package shortest_path

import (
	"fmt"
)

func min(i, j int) int {
	if i <= j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i <= j {
		return j
	} else {
		return i
	}
}

const INF = 2147483647
const NEG_INF = -2147483647

func Bellman_ford() {
	var n, m int
	fmt.Scan(&n, &m)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		edges[i] = append(edges[i], u, v, w)
	}

	stepDist := make([][]int, 2*n)
	for i := range stepDist {
		stepDist[i] = make([]int, n)
		for j := range stepDist[i] {
			stepDist[i][j] = INF
		}
		stepDist[i][0] = 0
	}

	for k := 0; k < n; k++ {
		if k != 0 {
			copy(stepDist[k], stepDist[k-1])
		}
		for i := 0; i < m; i++ {
			u, v, w := edges[i][0], edges[i][1], edges[i][2]
			if stepDist[k][u] != INF {
				stepDist[k][v] = min(stepDist[k][v], stepDist[k][u]+w)
			}
		}
	}

	for i := 0; i < n; i++ {
		if stepDist[n-2][i] != stepDist[n-1][i] {
			stepDist[n-1][i] = NEG_INF
		}
	}

	for k := n; k < 2*n; k++ {
		copy(stepDist[k], stepDist[k-1])
		for i := 0; i < m; i++ {
			u, v, _ := edges[i][0], edges[i][1], edges[i][2]
			if stepDist[k][u] == NEG_INF {
				stepDist[k][v] = NEG_INF
			}
		}
	}

	switch stepDist[2*n-1][n-1] {
	case INF:
		fmt.Println("impossible")
	case NEG_INF:
		fmt.Println("-inf")
	default:
		fmt.Println(stepDist[2*n-1][n-1])
	}
}
