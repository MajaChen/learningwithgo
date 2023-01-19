package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

/*
func totalDigits(length int) (digits int) {
    for curLength, curCount := 1, 9; curLength <= length; curLength++ {
        digits += curLength * curCount
        curCount *= 10
    }
    return
}

func findNthDigit(n int) int {
    d := 1 + sort.Search(8, func(length int) bool {
        return totalDigits(length+1) >= n
    })
    prevDigits := totalDigits(d - 1)
    index := n - prevDigits - 1
    start := int(math.Pow10(d - 1))
    num := start + index/d
    digitIndex := index % d
    return num / int(math.Pow10(d-digitIndex-1)) % 10
}

*/

//leetcode submit region begin(Prohibit modification and deletion)

func findNthDigit(n int) int {
	i, b := 1, 0
	for ; int(math.Pow(float64(10), float64(i-1)))*9*i < n; i++ {
		b += int(math.Pow(float64(10), float64(i-1))) * 9 * i
	}
	x := int(math.Pow10(i-1)) + (n-b-1)/i
	y := (n - b - 1) % i
	s := fmt.Sprint(x)
	a, _ := strconv.Atoi(s[y : y+1])
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
