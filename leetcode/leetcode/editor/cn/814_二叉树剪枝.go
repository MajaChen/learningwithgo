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

func pruneTreeWithRecursion(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var fl, fr bool
	if fl = pruneTreeWithRecursion(root.Left); !fl {
		root.Left = nil
	}
	if fr = pruneTreeWithRecursion(root.Right); !fr {
		root.Right = nil
	}
	return fl || fr || root.Val == 1
}

func pruneTree(root *TreeNode) *TreeNode {
	if pruneTreeWithRecursion(root) {
		return root
	} else {
		return nil
	}
}
//leetcode submit region end(Prohibit modification and deletion)
