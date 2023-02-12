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
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	s, ans := make([]*TreeNode, 0), make([][]int, 0)
	s = append(s, root)
	for len(s) != 0 {
		l, t := len(s), make([]int, 0)
		for i := 0; i < l; i++ {
			node := s[i]
			t = append(t, node.Val)
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
		}
		s = s[l:]
		ans = append(ans, t)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
