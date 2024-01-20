package main

import (
	dp "root/DP"
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

func main() {
	dp.LevenStein()
}
