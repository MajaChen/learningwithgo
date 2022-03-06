package leetcode

import "math"

func maxProduct(nums []int) int {

	maximalProduct := math.MinInt32
	max, min := 1, 1
	for i := 0; i < len(nums); i++ {

		amax, amin := max*nums[i], min*nums[i]

		var bmax, bmin int
		if bmax, bmin = amax, amin; amax < amin {
			bmax = amin
			bmin = amax
		}
		if bmax > nums[i] {
			max = bmax
		} else {
			max = nums[i]
		}
		if bmin < nums[i] {
			min = bmin
		} else {
			min = nums[i]
		}

		if max > maximalProduct {
			maximalProduct = max
		}
	}
	return maximalProduct
}

// 得到每一层的结点，二维数组的第一维标识节点层次
func getNodesPerLayer(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var q Queue = Queue{elems: make([]interface{}, 0)}
	q.Push(root)
	for !q.IsEmpty() {
		count, tmpRes, size := 0, make([]int, 0), q.Size()
		for count < size {
			node := q.Pop().(*TreeNode)
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
			count++
		}
		res = append(res, tmpRes)
	}
	return res
}

func rightSideView(root *TreeNode) []int {

	rightViesNodes := make([]int, 0)
	for i, nodes := 0, getNodesPerLayer(root); i < len(nodes); i++ {
		rightViesNodes = append(rightViesNodes, nodes[i][len(nodes[i])-1])
	}
	return rightViesNodes
}

func traverseIslands(grid [][]byte, i, j int) {

	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != 49 {
		return
	}

	grid[i][j] = 50
	traverseIslands(grid, i+1, j)
	traverseIslands(grid, i-1, j)
	traverseIslands(grid, i, j+1)
	traverseIslands(grid, i, j-1)
}

func numIslands(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 49 {
				count++
				traverseIslands(grid, i, j)
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 50 {
				grid[i][j] = 49
			}
		}
	}

	return count
}
