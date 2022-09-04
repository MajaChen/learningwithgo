package leetcode

import . "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var logestPathVal int

func logestUnivalPathRe(root *TreeNode, parentVal int) int {
	if root == nil {
		return 0
	}
	l, r := logestUnivalPathRe(root.Left, root.Val), logestUnivalPathRe(root.Right, root.Val)
	thisVal, returnVal := 0, 0
	if thisVal += l + r; thisVal > logestPathVal {
		logestPathVal = thisVal
	}
	if root.Val == parentVal {
		returnVal += 1 + Max(l, r)
	}

	return returnVal
}

func longestUnivaluePath(root *TreeNode) int {
	logestPathVal = 0
	if root == nil {
		return logestPathVal
	}
	if i := logestUnivalPathRe(root, root.Val) - 1; i > logestPathVal {
		logestPathVal = i
	}
	return logestPathVal
}

//leetcode submit region end(Prohibit modification and deletion)
