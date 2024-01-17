package numbertheory

import (
	"fmt"
	"math"
)

// エラトステネスの篩
// 整数Nまでの範囲で素数を列挙する
// 1. 素数の倍数は素数ではない
// 2. √N * √N > N であることから、Nが合成数の場合どちらか一方は√N以下であると言える
// 3. a * bとb * aは同じなので、片方を√N以下の片方の約数を計算すればもう片方は計算しなくて良い
// 以上の条件から、Nまでの数を配列で表し、√Nまで素数の倍数を消し続ければ残った数が素数となる

func Eratosthenes() {
	var n int
	fmt.Scan(&n)

	hurui := make([]bool, n+1)
	for i := range hurui {
		if i == 0 || i == 1 {
			continue
		}
		hurui[i] = true
	}

	root := int(math.Sqrt(float64(n)))
	for i := 0; i < root; i++ {
		if !hurui[i] {
			continue
		}
		for j := i * 2; j <= n; j = j + i {
			hurui[j] = false
		}
	}

	for i, v := range hurui {
		if v {
			fmt.Println(i)
		}
	}
}
