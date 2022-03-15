package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func longestSubsequence(arr []int, difference int) int {

	mapping, longestLength := make(map[int]int), 1
	for i := 0; i < len(arr); i++ {
		if _, ok := mapping[arr[i]-difference]; ok {
			k := mapping[arr[i]-difference]
			mapping[arr[i]] = k + 1
			if k+1 > longestLength {
				longestLength = k + 1
			}
		} else {
			if _, ok := mapping[arr[i]]; !ok {
				mapping[arr[i]] = 1
			}
		}
	}

	return longestLength
}

//leetcode submit region end(Prohibit modification and deletion)
