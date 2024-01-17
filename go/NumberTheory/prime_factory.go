package numbertheory

import (
	"fmt"
	"math"
)

// ただの素因数分解
// √N * √N > N を使う
func PrimeFactory() {
	var n int
	fmt.Scan(&n)

	root := int(math.Sqrt(float64(n)))
	ans := make([]int, 0, n)
	var mod int
	for i := 2; i <= root; i++ {
		for n%i == 0 {
			mod = n % i
			n = n / i
			ans = append(ans, i)
		}
	}

	if len(ans) == 0 {
		fmt.Println(n)
	} else {
		for i := range ans {
			fmt.Print(ans[i])
		}
		if mod != 0 {
			fmt.Print(mod)
		}
	}
}
