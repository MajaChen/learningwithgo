package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func distributeCandies(candyType []int) int {

	n, mapping := len(candyType)/2, make(map[int]bool)
	for _, val := range candyType {
		if !mapping[val] {
			mapping[val] = true
		}
	}

	if n >= len(mapping) {
		return len(mapping)
	} else {
		return n
	}
}

//leetcode submit region end(Prohibit modification and deletion)
