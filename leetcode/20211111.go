package leetcode

import (
	"math"
)

func isSymmetricRe(nodeA, nodeB *TreeNode) bool {
	if nodeA == nil && nodeB == nil {
		return true
	} else if nodeA != nil && nodeB != nil {
		if nodeA.Val != nodeB.Val {
			return false
		}
		return isSymmetricRe(nodeA.Left, nodeB.Right) && isSymmetricRe(nodeA.Right, nodeB.Left)
	} else {
		return false
	}
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricRe(root.Left, root.Right)
}

func maxArea(height []int) int {
	maxArea := math.MinInt32
	for l, r := 0, len(height)-1; l < r; {
		var minHeight int
		var flag bool
		if minHeight = height[l]; minHeight > height[r] {
			minHeight = height[r]
			flag = true
		}
		if (r-l)*minHeight > maxArea {
			maxArea = (r - l) * minHeight
		}

		if flag {
			r--
		} else {
			l++
		}
	}
	return maxArea
}
