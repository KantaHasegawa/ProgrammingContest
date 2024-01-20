package dp

import (
	"fmt"
)

// 部分和問題
// https://algo-method.com/tasks/337
// dp[i][j]はa[i]を選択したorしないときに[j]になるかどうかを示す
// 例えばa[i]が7でjも7の時、7を選択する前の数値dp[i-1][j-7]がTrueであれば、dp[i][7]もtrueになる

func Bubunwa() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			choose := a[i-1]
			if 0 <= j-choose {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-choose]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	var ans bool
	for i := range dp {
		ans = ans || dp[i][m]
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
