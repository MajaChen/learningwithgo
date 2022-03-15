package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func longestLength(nums []int, target int) int {

	var longestLength int
	for i, sum, l := 0, 0, 0; i < len(nums); {
		if sum += nums[i]; sum == target {
			if i-l+1 > longestLength {
				longestLength = i - l + 1
			}
		} else if sum > target { // 尝试右移l
			for sum > target && l <= i {
				sum -= nums[l]
				l++
			}
			if sum == target {
				sum -= nums[i]
				continue
			}
		}
		i++
	}
	return longestLength
}

func minOperations(nums []int, x int) int {

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum-x == 0 {
		return len(nums)
	} else if sum-x < 0 {
		return -1
	} else {
		if l := longestLength(nums, sum-x); l == 0 {
			return -1
		} else {
			return len(nums) - l
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
