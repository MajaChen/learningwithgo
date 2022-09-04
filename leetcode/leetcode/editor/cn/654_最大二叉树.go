package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findMaxIndex(nums []int) int {
	maxVal, maxIndex := -1, math.MinInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] >= maxVal {
			maxVal = nums[i]
			maxIndex = i
		}
	}
	return maxIndex
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := findMaxIndex(nums)
	node := &TreeNode{Val: nums[i]}
	node.Left = constructMaximumBinaryTree(nums[:i])
	node.Right = constructMaximumBinaryTree(nums[i+1:])
	return node
}
//leetcode submit region end(Prohibit modification and deletion)
