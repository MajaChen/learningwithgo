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

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		t := nums[i]
		nums[i] = nums[n-i-1]
		nums[n-i-1] = t
	}
}
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	s, ans, c := make([]*TreeNode, 0), make([][]int, 0), 1
	s = append(s, root)
	for len(s) != 0 {
		l, t := len(s), make([]int, 0)
		c++
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
		if c%2 != 0 {
			reverse(t)
		}
		ans = append(ans, t)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
