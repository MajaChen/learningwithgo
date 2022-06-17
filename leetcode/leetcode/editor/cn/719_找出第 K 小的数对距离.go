package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)

func count(nums []int, m int) int {
	n, c := len(nums), 0
	for i, j := 0, 1; i < n; i++ {
		for ; j < n && nums[j]-nums[i] <= m; j++ {
		}
		c += j - i - 1
	}
	return c
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	l, r, m := 0, nums[len(nums)-1]-nums[0], 0
	for l < r {
		m = (l + r) >> 1
		if c := count(nums, m); c >= k {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

//leetcode submit region end(Prohibit modification and deletion)
