package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func frequencySort(nums []int) []int {
	mapping := make(map[int]int)
	for _, num := range nums {
		mapping[num]++
	}

	sort.Slice(nums, func(i, j int) bool {
		return mapping[nums[i]] < mapping[nums[j]] ||
			(mapping[nums[i]] == mapping[nums[j]] && nums[i] > nums[j])
	})

	return nums
}
//leetcode submit region end(Prohibit modification and deletion)
