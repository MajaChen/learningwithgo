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

func addOneRowTraverse(root *TreeNode, val, curDepth, depth int) {
	if root == nil {
		return
	}
	if curDepth++; curDepth == depth-1 {
		root.Left = &TreeNode{Val: val, Left: root.Left}
		root.Right = &TreeNode{Val: val, Right: root.Right}
		return
	}
	addOneRowTraverse(root.Left, val, curDepth, depth)
	addOneRowTraverse(root.Right, val, curDepth, depth)
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	vRoot := &TreeNode{Left: root}
	addOneRowTraverse(vRoot, val, -1, depth)
	return vRoot.Left
}

//leetcode submit region end(Prohibit modification and deletion)
