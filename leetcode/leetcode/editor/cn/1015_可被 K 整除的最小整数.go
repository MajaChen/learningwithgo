package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func smallestRepunitDivByK(k int) int {

	mapping := make(map[int]bool)
	for mod, count := 0, 1; ; {
		mod = (mod*10 + 1) % k

		if mod == 0 {
			return count
		}

		if _, ok := mapping[mod]; ok {
			return -1
		}

		mapping[mod] = true
		count++
	}
}

//leetcode submit region end(Prohibit modification and deletion)
