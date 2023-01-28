package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)

func findNthDigit(n int) int {
	i, b := 1, 0 // i 几位数
	for ; int(math.Pow(float64(10), float64(i-1)))*9*i < n; i++ {
		b += int(math.Pow(float64(10), float64(i-1))) * 9 * i
	}
	x := int(math.Pow10(i-1)) + (n-b-1)/i // x 第几个数，n-b-1是考虑0
	y := (n - b - 1) % i                  // y 第几位数
	s := fmt.Sprint(x)
	a, _ := strconv.Atoi(s[y : y+1])
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
