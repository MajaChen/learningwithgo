package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

func rearrangeNums(nums []int){
	n := len(nums)
	for i := 0;i < n/2;i++{
		tmp := nums[i]
		nums[i] = nums[n-i-1]
		nums[n-i-1] = tmp
	}
}

// 4 5 3 2 1
func nextPermutation(nums []int){
	// 从右到左寻找第一个“非递增的元素”
	var i int
	for i = len(nums)-2;i >= 0 && nums[i] >= nums[i+1];i--{}
	var startIndex,endIndex int
	if i >= 0{ // 从左到右寻找最后一个比nums[i]大的元素
		var j int
		for j = i+1;j < len(nums) && nums[j] > nums[i];j++{}
		j--
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
	}
	// 交换nums[i]与nums[j]
	startIndex,endIndex = i+1,len(nums)
	rearrangeNums(nums[startIndex:endIndex])
}

func nextGreaterElement(n int) int {
	nums, nn := make([]int, 0, 32), n
	for n != 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	rearrangeNums(nums)
	nextPermutation(nums)
	ans := 0
	for i, m := len(nums)-1, 1; i >= 0; i-- {
		ans += nums[i]*m
		m *= 10
	}
	if ans <= nn || ans > math.MaxInt32 {
		return -1
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
