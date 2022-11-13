package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func isGoodNum(n int) bool {
	var f bool
	for n > 0 {
		i := n%10
		if i == 3 || i == 4 || i == 7 {
			return false
		}
		if i == 2 || i == 5 || i == 6 || i == 9 {
			f = true
		}
		n /= 10
	}
	return f
}

func rotatedDigits(n int) int {
	var ans int
	for i := 1; i <= n; i++ {
		if isGoodNum(i) {
			ans++
		}
	}

	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
