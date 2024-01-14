package dfs

import (
	"fmt"
	"strings"
)

// https://algo-method.com/tasks/955
func ArgoshikiBlackArea() {
	var h, w int
	fmt.Scan(&h, &w)
	graph := make([][]string, h)
	for i := range graph {
		var str string
		fmt.Scan(&str)
		graph[i] = strings.Split(str, "")
	}

	seen := make([][]bool, h)
	for i := range seen {
		seen[i] = make([]bool, w)
	}

	var rec func(graph [][]string, x, y int)
	rec = func(graph [][]string, x, y int) {
		seen[x][y] = true
		roots := nextRow(graph, x, y)
		for _, root := range roots {
			if !isAccess(graph, root[0], root[1]) {
				continue
			}
			if seen[root[0]][root[1]] {
				continue
			}
			rec(graph, root[0], root[1])
		}
	}

	var ans int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if graph[i][j] == "#" {
				if seen[i][j] {
					continue
				}
				rec(graph, i, j)
				ans++
			}
		}
	}

	fmt.Println(ans)
}

func isAccess(ary [][]string, x, y int) bool {
	defer func() {
		recover()
	}()

	return ary[x][y] == "#"
}

func nextRow(ary [][]string, x, y int) [][]int {
	top := []int{x, y + 1}
	right := []int{x + 1, y}
	bottom := []int{x, y - 1}
	left := []int{x - 1, y}

	return [][]int{top, right, bottom, left}
}
