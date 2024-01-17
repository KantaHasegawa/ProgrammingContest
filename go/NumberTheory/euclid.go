package numbertheory

import "fmt"

// ユークリッドの互除法
// AとBの最大公約数を求める A = B * q + mod の最大公約数は Bとmodの最大公約数と同じである
// プログラム上ではmodが求まればいいので、 A % B を何回も行い、modが0になった時のbが最大公約数である

func Euclid() {
	var a, b int
	fmt.Scan(&a, &b)
	a = max(a, b)
	b = min(a, b)

	var ans int
	for {
		mod := a % b
		if mod == 0 {
			ans = b
			break
		}
		a, b = b, mod
	}

	fmt.Println(ans)
}
