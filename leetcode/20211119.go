package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	var flagP, flagQ bool
	for {
		if p == q && p != nil && q != nil {
			return p
		}
		if p.Next == nil {
			if flagP {
				return nil
			}
			p = headB
			flagP = true
		} else {
			p = p.Next
		}
		if q.Next == nil {
			if flagQ {
				return nil
			}
			q = headA
			flagQ = true
		} else {
			q = q.Next
		}
	}
}

func rearrangeNums(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		tmp := nums[i]
		nums[i] = nums[n-i-1]
		nums[n-i-1] = tmp
	}
}

// 4 5 3 2 1
func nextPermutationV2(nums []int) {
	// 从右到左寻找第一个“非递增的元素”
	var i int
	for i = len(nums) - 2; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	var startIndex, endIndex int
	if i >= 0 { // 从左到右寻找第一个比nums[i]大的元素
		var j int
		for j = i + 1; j < len(nums) && nums[j] > nums[i]; j++ {
		}
		j--
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
	}
	// 交换nums[i]与nums[j]
	startIndex, endIndex = i+1, len(nums)
	rearrangeNums(nums[startIndex:endIndex])
}
