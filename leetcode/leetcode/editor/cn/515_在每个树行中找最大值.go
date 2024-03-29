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

import . "learning/common"

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	var q Queue = Queue{Elems: make([]interface{}, 0)}
	q.Push(root)
	for !q.IsEmpty() {
		size, count, max := q.Size(), 0, math.MinInt32
		for count < size {
			node := q.Pop().(*TreeNode)
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
			count++
		}
		count = 0
		res = append(res, max)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
