package bfs

import (
	"fmt"
	"strings"
)

// https://atcoder.jp/contests/abc007/tasks/abc007_3
func AtcoderBFSPractice() {
	var r, c, sx, sy, gx, gy int
	fmt.Scan(&r, &c)
	fmt.Scan(&sx, &sy)
	fmt.Scan(&gx, &gy)
	sx--
	sy--
	gx--
	gy--

	graph := make([][]string, r)
	for i := range graph {
		var str string
		fmt.Scan(&str)
		graph[i] = strings.Split(str, "")
	}

	dist := make([][]int, r)
	for i := range dist {
		dist[i] = make([]int, c)
	}
	dist[sx][sy] = 0

	flg := make([][]bool, r)
	for i := range flg {
		flg[i] = make([]bool, c)
	}
	flg[sx][sy] = true

	que := make([][]int, 0, r)
	que = append(que, []int{sx, sy})

	for len(que) > 0 {
		front := que[0]
		que = que[1:]

		root := nextRow(graph, front[0], front[1])
		for _, v := range root {
			if !isAccess(graph, v[0], v[1]) {
				continue
			}
			if flg[v[0]][v[1]] {
				continue
			}
			flg[v[0]][v[1]] = true
			dist[v[0]][v[1]] = dist[front[0]][front[1]] + 1
			que = append(que, v)
		}
	}
	fmt.Println(dist[gx][gy])
}

func isAccess(ary [][]string, x, y int) bool {
	defer func() {
		recover()
	}()

	return ary[x][y] != "#"
}

func nextRow(ary [][]string, x, y int) [][]int {
	top := []int{x, y + 1}
	right := []int{x + 1, y}
	bottom := []int{x, y - 1}
	left := []int{x - 1, y}

	return [][]int{top, right, bottom, left}
}
