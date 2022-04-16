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

var maximalPathSum int

func maxPathSumRe(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var leftMaxPathSum, rightMaxPathSum, maximal int
	if leftMaxPathSum, rightMaxPathSum, maximal = maxPathSumRe(root.Left), maxPathSumRe(root.Right), root.Val; leftMaxPathSum < rightMaxPathSum {
		if rightMaxPathSum > 0 {
			maximal = root.Val + rightMaxPathSum
		}
	} else {
		if leftMaxPathSum > 0 {
			maximal = root.Val + leftMaxPathSum
		}
	}

	if maximal > maximalPathSum {
		maximalPathSum = maximal
	}

	if root.Val+leftMaxPathSum+rightMaxPathSum > maximalPathSum {
		maximalPathSum = root.Val + leftMaxPathSum + rightMaxPathSum
	}

	return maximal
}

func maxPathSum(root *TreeNode) int {

	maximalPathSum = math.MinInt32
	maxPathSumRe(root)
	return maximalPathSum
}

//leetcode submit region end(Prohibit modification and deletion)
