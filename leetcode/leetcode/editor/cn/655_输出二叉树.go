package leetcode

import (
	"math"
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maximalHeight int

func getMaximalHeight(root *TreeNode, baseHeight int) {
	if root == nil {
		return
	}
	if baseHeight += 1; baseHeight > maximalHeight {
		maximalHeight = baseHeight
	}
	getMaximalHeight(root.Left, baseHeight)
	getMaximalHeight(root.Right, baseHeight)
}

func put(root *TreeNode, r, c, h int, matrix [][]string) {
	matrix[r][c] = strconv.Itoa(root.Val)
	if root.Left != nil {
		put(root.Left, r+1, c-int(math.Pow(2, float64(h-r-1))), h, matrix)
	}
	if root.Right != nil {
		put(root.Right, r+1, c+int(math.Pow(2, float64(h-r-1))), h, matrix)
	}
}

func printTree(root *TreeNode) [][]string {
	maximalHeight = -1
	getMaximalHeight(root, -1)
	h := maximalHeight
	m, n := h+1, int(math.Pow(2, float64(h+1))-1)
	matrix := make([][]string, 0, m)
	for i := 0; i < m; i++ {
		matrix = append(matrix, make([]string, n))
	}
	put(root, 0, (n-1)/2, h, matrix)
	return matrix
}

//leetcode submit region end(Prohibit modification and deletion)
