package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func singleNumbers(nums []int) []int {
	s := 0
	for _, num := range nums {
		s ^= num
	}

	div := 1
	for div&s == 0 {
		div <<= 1
	}

	a, b := 0, 0
	for _, num := range nums {
		if div&num == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}

//leetcode submit region end(Prohibit modification and deletion)
