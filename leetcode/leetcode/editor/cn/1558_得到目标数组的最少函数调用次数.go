package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func minOperations2(nums []int) int {

	addCount, maxMultiCount := 0, 0
	for i := 0; i < len(nums); i++ {
		multiCount := 0
		for nums[i] > 0 {
			if nums[i]%2 == 0 {
				multiCount += 1
				nums[i] /= 2
			} else if nums[i]%2 == 1 {
				addCount += 1
				nums[i] -= 1
			}
		}
		if multiCount > maxMultiCount {
			maxMultiCount = multiCount
		}
	}

	return addCount + maxMultiCount
}

//leetcode submit region end(Prohibit modification and deletion)
