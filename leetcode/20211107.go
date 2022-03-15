package leetcode

import "math"

func maxSubArray2(nums []int) int {
	sum, maxSum := nums[0], math.MinInt32
	for i := 1; i < len(nums); i++ {
		if nums[i] >= 0 { //这里存在一个预期：sum是正值
			if sum >= 0 {
				sum += nums[i]
			} else {
				sum = nums[i]
			}
		} else {
			// 考虑要不要把sum加上去
			if sum >= 0 {

			} else {

			}
			if sum > maxSum {
				maxSum = sum
			}

			if sum += nums[i]; sum < 0 { // 不可抵扣，重开一句
				if i < len(nums)-1 {
					sum = nums[i+1]
					i++
				}
			}
		}
	}

	if sum > maxSum {
		maxSum = sum
	}

	return maxSum
}
