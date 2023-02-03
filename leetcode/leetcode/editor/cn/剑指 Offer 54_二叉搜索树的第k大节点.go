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

var kthLargestSlice []int

func kthLargestTraverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	// 左节点
	kthLargestTraverse(root.Left, k)
	// 根节点
	kthLargestSlice = append(kthLargestSlice, root.Val)
	// 右节点
	kthLargestTraverse(root.Right, k)
}

func kthLargest(root *TreeNode, k int) int {
	kthLargestTraverse(root, k)
	return kthLargestSlice[len(kthLargestSlice)-k]
}

//leetcode submit region end(Prohibit modification and deletion)
