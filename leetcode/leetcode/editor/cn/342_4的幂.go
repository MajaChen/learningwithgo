package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfFour(n int) bool {

	if n == 0 {
		return false
	}

	return int(math.Pow(4, float64(int(math.Log(float64(n))/math.Log(4))))) == n
}

//leetcode submit region end(Prohibit modification and deletion)
