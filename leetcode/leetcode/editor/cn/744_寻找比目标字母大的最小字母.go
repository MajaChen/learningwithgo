package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func nextGreatestLetter(letters []byte, target byte) byte {

	var l, r int
	for l, r = 0, len(letters)-1; l < r; {
		if mid := (l + r) >> 1; letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if letters[l] <= target {
		return letters[0]
	} else {
		return letters[l]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
