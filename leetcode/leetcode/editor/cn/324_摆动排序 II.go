package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func wiggleSort(nums []int) {
	sort.Ints(nums)
	var i, k int
	copiedNums := make([]int, len(nums))
	copy(copiedNums, nums)
	for i = 0; i < len(nums)/2; i++ {
		nums[k] = copiedNums[i]
		k++
		nums[k] = copiedNums[len(nums)-1-i]
		k++
	}

	if len(nums)%2 != 0 {
		nums[k] = copiedNums[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
