package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)

var maximalSum int

func maxSumDivThreeRe(nums []int, sum int) {

	if sum%3 == 0 {
		if sum > maximalSum {
			maximalSum = sum
		}
		return
	}

	for i := 0; sum > maximalSum && i < len(nums); i++ {
		if nums[i] < 0 {
			continue
		}

		j := nums[i]
		nums[i] = -1
		maxSumDivThreeRe(nums, sum-j)
		nums[i] = j
	}

}

func maxSumDivThree(nums []int) int {

	maximalSum = 0
	sort.Ints(nums)
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	maxSumDivThreeRe(nums, sum)

	return maximalSum
}

//leetcode submit region end(Prohibit modification and deletion)
