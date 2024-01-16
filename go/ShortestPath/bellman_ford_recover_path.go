package shortest_path

import "fmt"

func BellmanFordRecoverPath() {
	var n, m int
	fmt.Scan(&n, &m)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		edges[i] = append(edges[i], u, v, w)
	}

	stepDist := make([][]int, n)
	for i := range stepDist {
		stepDist[i] = make([]int, n)
		for j := range stepDist[i] {
			stepDist[i][j] = INF
		}
		stepDist[i][0] = 0
	}

	restore := make([]int, n)
	for i := range restore {
		restore[i] = -1
	}

	for k := 0; k < n; k++ {
		if k != 0 {
			copy(stepDist[k], stepDist[k-1])
		}
		for i := 0; i < m; i++ {
			u, v, w := edges[i][0], edges[i][1], edges[i][2]
			if stepDist[k][u] != INF && stepDist[k][u]+w < stepDist[k][v] {
				stepDist[k][v] = min(stepDist[k][v], stepDist[k][u]+w)
				restore[v] = u
			}
		}
	}

	ans := make([]int, 0, n)
	history := n-1
	for history != -1 {
		ans = append(ans, history)
		history = restore[history]
	}
	for i, j := 0, len(ans)-1; i < j; i,j = i+1, j-1{
		ans[i], ans[j] = ans[j], ans[i]
	}

	fmt.Println(len(ans))
	for _, v := range ans {
		fmt.Print(v, " ")
	}
}
