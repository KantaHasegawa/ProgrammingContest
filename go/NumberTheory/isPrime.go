package numbertheory

import (
	"fmt"
	"math"
)

// 素数判定
// √N * √N > N であることから、Nが合成数の場合どちらか一方は√N以下であると言える
// a * bとb * aは同じなので、片方を√N以下の片方の約数を計算すればもう片方は計算しなくて良い

func IsPrime() {
	var n int
	fmt.Scan(&n)

	var ans bool
	root := int(math.Sqrt(float64(n)))
	if n%2 == 0 {
		fmt.Println(n, "is not prime")
		return
	}
	for i := 3; i < root; i = i + 2 {
		if n%i == 0 {
			ans = true
			break
		}
	}

	if ans {
		fmt.Println(n, "is not prime")
	} else {
		fmt.Println(n, "is prime")
	}
}
