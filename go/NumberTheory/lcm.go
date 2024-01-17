package numbertheory

import "fmt"

// 最小公倍数を求める
// N * M = LCM(最小公倍数) * GCM(最大公約数) の式を利用する
// 変形すると、LCM = (N * M) / GCM となる
func LCM() {
	var n, m int
	fmt.Scan(&n, &m)
	a := max(n, m)
	b := min(n, m)

	var gcm int
	for {
		mod := a % b
		if mod == 0 {
			gcm = b
			break
		}
		a, b = b, mod
	}

	lcm := (n * m) / gcm
	fmt.Println(lcm)
}
