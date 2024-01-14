package bfs

import (
	"fmt"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/joi2011yo/tasks/joi2011yo_e
func Cheese() {
	var h, w, n int
	fmt.Scan(&h, &w, &n)
	cap := h
	if h < w {
		cap = w
	}

	graph := make([][]string, h)
	var start []int
	for i := range graph {
		var str string
		fmt.Scan(&str)
		strs := strings.Split(str, "")
		for ii, v := range strs {
			if v == "S" {
				start = []int{i, ii}
				v = "."
			}
			graph[i] = append(graph[i], v)
		}
	}

	var ans int

	search := func(sx, sy, goal int) {
		dist := make([][]int, h)
		for i := range dist {
			dist[i] = make([]int, w)
		}
		dist[sx][sy] = 0

		flg := make([][]bool, h)
		for i := range flg {
			flg[i] = make([]bool, w)
		}
		flg[sx][sy] = true

		que := make([][]int, 0, cap)
		que = append(que, []int{sx, sy})

		var find bool
		for len(que) > 0 && !find {
			front := que[0]
			que = que[1:]

			roots := nextRow2(graph, front[0], front[1])
			for _, root := range roots {
				if !isAccess2(graph, root[0], root[1], strconv.Itoa(goal)) {
					continue
				}
				if flg[root[0]][root[1]] {
					continue
				}
				flg[root[0]][root[1]] = true
				dist[root[0]][root[1]] = dist[front[0]][front[1]] + 1

				if graph[root[0]][root[1]] == strconv.Itoa(goal) {
					ans += dist[root[0]][root[1]]
					start = []int{root[0], root[1]}
					graph[root[0]][root[1]] = "."
					find = true
				}
				que = append(que, root)
			}
		}
	}

	for i := 1; i <= n; i++ {
		search(start[0], start[1], i)
	}

	fmt.Println(ans)
}

func isAccess2(ary [][]string, x, y int, goal string) bool {
	defer func() {
		recover()
	}()

	return ary[x][y] != "X"
}

func nextRow2(ary [][]string, x, y int) [][]int {
	top := []int{x, y + 1}
	right := []int{x + 1, y}
	bottom := []int{x, y - 1}
	left := []int{x - 1, y}

	return [][]int{top, right, bottom, left}
}
