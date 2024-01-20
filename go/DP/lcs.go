package dp

import (
	"fmt"
	"strings"
)

// 最長共通部分文字列
func LCS() {
	var a, b string
	fmt.Scan(&a, &b)

	s := strings.Split(a, "")
	t := strings.Split(b, "")

	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				dp[i][j] = max(dp[i][j], dp[i-1][j-1])
			}
		}
	}

	fmt.Println(dp[len(s)][len(t)])
}

// 最長共通部分文字列(連結)
func LCSChain() {
	var a, b string
	fmt.Scan(&a, &b)

	s := strings.Split(a, "")
	t := strings.Split(b, "")

	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	var ans int
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1] + 1
				ans = max(ans, dp[i][j])
			} else {
				dp[i][j] = 0
			}
		}
	}

	fmt.Println(ans)
}
