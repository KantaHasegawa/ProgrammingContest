package dp

import (
	"fmt"
	"strings"
)

func LevenStein() {
	var a, b string
	fmt.Scan(&a, &b)
	s := strings.Split(a, "")
	t := strings.Split(b, "")

	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 0; i < len(s)+1; i++ {
		for j := 0; j < len(t)+1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			switch {
			case i == 0:
				// sが空文字でtはindexの文字数である
				// そのため、sとtを同じにするにはsをj回挿入 or tをj回削除となる
				// j-1で既に編集距離を計算できているため、そこに+1をすればよい
				dp[i][j] = dp[0][j-1] + 1
			case j == 0:
				dp[i][j] = dp[i-1][0] + 1
			default:
				if s[i-1] == t[j-1] {
					// 文字が同じとき,追加した文字同士の編集距離が0なので足さなくても良い
					// ただし、dp[i-1][j-1]からは同じ文字をsとtに追加しているのに対し、
					// dp[i-1][j]やdp[i][j-1]はsかtの片方に1文字しか追加していないため、同じかどうか関係なく文字数の数差分を削除or変更する分+1される
					dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]+1, dp[i][j-1]+1)
				} else {
					dp[i][j] = min(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
				}
			}
		}
	}

	fmt.Println(dp[len(s)][len(t)])
}
