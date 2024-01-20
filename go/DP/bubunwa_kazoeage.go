package dp

import "fmt"

// 部分和数え上げ問題
// dp[0][0]を、0個選んだ時0になるパターンは1通りとみなし、部分話数え上げと同様にdp表からパターンを足していく
// dp[i][j]は。a[i]を選ばないパターンdp[i-1][j]通りと、a[i]を選ぶパターンdp[i-1][j-a[i]]通りの合計になる

func max(i, j int) int {
	if i <= j {
		return j
	} else {
		return i
	}
}

func BubunwaKazoeage() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			choose := a[i-1]
			dp[i][j] += dp[i-1][j]
			if choose <= j {
				dp[i][j] += dp[i-1][j-choose]
			}
		}
	}

	var ans int
	for i := range dp {
		ans = max(ans, dp[i][m+1])
	}

	fmt.Println(ans)
}
