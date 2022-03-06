package leetcode

func buildPaths(root, target *TreeNode, paths []*TreeNode) ([]*TreeNode, bool) {

	if root == nil {
		return nil, false
	}

	if root == target {
		paths = append(paths, root)
		return paths, true
	}

	if lpaths, flag := buildPaths(root.Left, target, paths); flag {
		paths = lpaths
		paths = append(paths, root)
		return paths, true
	}

	if rpaths, flag := buildPaths(root.Right, target, paths); flag {
		paths = rpaths
		paths = append(paths, root)
		return paths, true
	}

	return nil, false
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	ppaths := make([]*TreeNode, 0)
	qpaths := make([]*TreeNode, 0)
	ppaths, _ = buildPaths(root, p, ppaths)
	qpaths, _ = buildPaths(root, q, qpaths)

	pk, qk := -1, -1
	for i, j := len(ppaths)-1, len(qpaths)-1; i >= 0 && j >= 0; {
		if ppaths[i] != qpaths[j] {
			pk, qk = i+1, j+1
			break
		}
		i--
		j--
	}

	if pk >= len(ppaths) && qk >= len(qpaths) {
		return nil
	}

	if pk < 0 && qk < 0 {
		if len(ppaths) < len(qpaths) {
			return ppaths[0]
		} else {
			return qpaths[0]
		}
	}

	return ppaths[pk]
}

func productExceptSelf(nums []int) []int {

	multiA, multiB := make([]int, len(nums)), make([]int, len(nums))
	for i, seed := 0, 1; i < len(nums); i++ {
		multiA[i] = seed * nums[i]
		seed *= nums[i]
	}

	for i, seed := len(nums)-1, 1; i >= 0; i-- {
		multiB[i] = seed * nums[i]
		seed *= nums[i]
	}

	res := make([]int, len(nums))
	for i := 1; i < len(nums)-1; i++ {
		res[i] = multiA[i-1] * multiB[i+1]
	}
	res[0] = multiB[1]
	res[len(nums)-1] = multiA[len(nums)-2]
	return res
}

func searchMatrix(matrix [][]int, target int) bool {

	for i, j := 0, len(matrix[0]); i < len(matrix) && j >= 0; {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target { //往左走
			j -= 1
		} else {
			i += 1
		}
	}
	return false

}
