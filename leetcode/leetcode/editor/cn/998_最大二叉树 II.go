package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoMaxTreeRe(current, parent *TreeNode, val int) {
	node := &TreeNode{Val: val}
	if current == nil {
		parent.Right = node
		return
	}
	if val > current.Val {
		if parent.Left == current {
			parent.Left = node
			node.Left = current
		} else {
			parent.Right = node
			node.Left = current
		}
	} else {
		insertIntoMaxTreeRe(current.Right, current, val)
	}
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	vp := &TreeNode{Left: root}
	insertIntoMaxTreeRe(root, vp, val)
	return vp.Left
}
//leetcode submit region end(Prohibit modification and deletion)
