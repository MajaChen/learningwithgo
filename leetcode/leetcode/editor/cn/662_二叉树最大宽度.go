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

type pair struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	q := []pair{{node: root, index: 1}}
	ans := 1
	for q != nil {
		ans = Max(ans, q[len(q)-1].index-q[0].index+1)
		t := q
		q = nil
		for _, p := range t {
			if p.node.Left != nil {
				q = append(q, pair{p.node.Left, p.index * 2})
			}
			if p.node.Right != nil {
				q = append(q, pair{p.node.Right, p.index*2 + 1})
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
