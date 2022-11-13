package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func minSwap(nums1 []int, nums2 []int) int {
	a, b, n := 0, 1, len(nums1)
	var sa, sb int
	for i := 1; i < len(nums1); i++ {
		sa, sb = a, b
		a, b = n, n
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			a = min(a, sa)
			b = min(b, sb+1)
		}
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			a = min(a, sb)
			b = min(b, sa+1)
		}
	}
	return min(a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
