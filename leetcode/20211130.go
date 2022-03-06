package leetcode

import "math"

// 从startIndex开始，向后打劫所得最大值
var robMapping []int

// 从startIndex开始（包括它本身），向后打劫所得最大值
func robRe2(nums []int, startIndex int) int {
	if robMapping[startIndex] >= 0 {
		return robMapping[startIndex]
	}

	maximalRobValue := 0
	for i := startIndex + 2; i < len(nums); i++ {
		// 打劫完了“这家”，向后继续打劫所得最大值
		if value := robRe2(nums, i); value > maximalRobValue {
			maximalRobValue = value
		}
	}

	robMapping[startIndex] = nums[startIndex] + maximalRobValue
	return robMapping[startIndex]
}

func rob2(nums []int) int {
	robMapping = make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		robMapping = append(robMapping, -1)
	}
	maxRobValue := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if value := robRe2(nums, i); value > maxRobValue {
			maxRobValue = value
		}
	}
	return maxRobValue
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{Next: head}
	for p := newHead; p != nil && p.Next != nil; {
		if p.Next.Val == val { // 执行删除操作
			p.Next = p.Next.Next
		} else { // 指针后移
			p = p.Next
		}
	}
	return newHead.Next
}

func sortColors(nums []int) {

	for leftIndex, rightIndex, i := 0, len(nums)-1, 0; i <= rightIndex; {
		if nums[i] == 2 {
			swap(i, rightIndex, nums)
			rightIndex--
		} else if nums[i] == 0 {
			swap(i, leftIndex, nums)
			leftIndex++
			i++
		} else {
			i++
		}
	}

}
