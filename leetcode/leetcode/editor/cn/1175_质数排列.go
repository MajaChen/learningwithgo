package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

var (
	MagicNum = int64(math.Pow(10, 9)) + int64(7)
)

func countPrime(n int) int {
	if n <= 0 {
		return 0
	}
	counter, isPrime := 0, make([]int, n+1)
	for i := 2; i <= n; i++ {
		if isPrime[i] == 0 {
			counter++
			for j := i * i; j <= n; j += i {
				isPrime[j] = 1
			}
		}
	}
	return counter
}

func mmulti(n int) int64 {
	var mm int64 = 1
	for i := 1; i <= n; i++ {
		mm = mm * int64(i) % MagicNum
	}
	return mm
}

func numPrimeArrangements(n int) int {
	x := countPrime(n)
	y := n - x
	return int((mmulti(x) * mmulti(y)) % MagicNum)
}

//leetcode submit region end(Prohibit modification and deletion)
