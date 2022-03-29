package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func minArray(numbers []int) int {

	var l, r int
	for l, r = 0, len(numbers)-1; l < r; {
		if mid := (l + r) >> 1; numbers[mid] > numbers[r] { // 放弃左侧
			l = mid + 1
		} else if numbers[mid] < numbers[r] { // 放弃右侧
			r = mid
		} else { // 无法取舍，缩小区间
			r = r - 1
		}
	}
	return numbers[l]
}

//leetcode submit region end(Prohibit modification and deletion)
