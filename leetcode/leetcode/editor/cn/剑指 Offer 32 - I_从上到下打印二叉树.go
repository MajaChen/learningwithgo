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
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	s, ans := make([]*TreeNode, 0), make([]int, 0)
	s = append(s, root)
	for len(s) != 0 {
		l := len(s)
		for i := 0; i < l; i++ {
			node := s[i]
			ans = append(ans, node.Val)
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
		}
		s = s[l:]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
