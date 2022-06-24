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


type Queue struct {
	elems []interface{}
}

func (q *Queue) Size() int {
	return len(q.elems)
}

func (q *Queue) Pop() interface{} {
	elem := q.elems[0]
	q.elems = q.elems[1:]
	return elem
}

func (q *Queue) Push(elem interface{}) {
	q.elems = append(q.elems, elem)
}

func (q Queue) IsEmpty() bool {
	return len(q.elems) == 0
}

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	var q Queue = Queue{elems: make([]interface{}, 0)}
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
