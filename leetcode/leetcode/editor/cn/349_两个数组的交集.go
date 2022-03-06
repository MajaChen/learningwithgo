package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {

	mapping, ans := make(map[int]bool), make([]int, 0)
	for _, num := range nums1 {
		mapping[num] = true
	}

	for _, num := range nums2 {
		if mapping[num] {
			ans = append(ans, num)
			delete(mapping, num)
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
