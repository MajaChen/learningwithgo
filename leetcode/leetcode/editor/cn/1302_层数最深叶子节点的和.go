package leetcode

import ."learning/common"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := Queue{Elems: make([]interface{}, 0)}
	sum := 0
	q.Push(root)
	for !q.IsEmpty() {
		sum = 0
		for i := q.Size(); i > 0; i-- {
			node := q.Pop().(*TreeNode)
			sum += node.Val
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}
	return sum
}
//leetcode submit region end(Prohibit modification and deletion)
