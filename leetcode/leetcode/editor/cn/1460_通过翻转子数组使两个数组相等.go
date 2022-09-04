package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	if len(target) == 0 {
		return true
	}
	sa := make(map[int]int)
	for i := 0; i < len(target); i++ {
		sa[target[i]]++
		sa[arr[i]]--
	}
	for _, j := range sa {
		if j != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
