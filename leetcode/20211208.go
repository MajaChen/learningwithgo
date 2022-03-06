package leetcode

var paths [][]int

func pathSumRe(root *TreeNode, targetSum int, localPaths []int) {

	if root == nil {
		return
	}

	localPaths = append(localPaths, root.Val)
	if targetSum -= root.Val; targetSum == 0 && root.Left == nil && root.Right == nil {
		dst := make([]int, len(localPaths))
		copy(dst, localPaths)
		paths = append(paths, dst)
		return
	}

	l := len(localPaths)
	pathSumRe(root.Left, targetSum, localPaths)
	pathSumRe(root.Right, targetSum, localPaths)
	localPaths = localPaths[:l]
}

func pathSum(root *TreeNode, targetSum int) [][]int {

	paths = make([][]int, 0)
	pathSumRe(root, targetSum, make([]int, 0))
	return paths
}

var pathNumSum int

func sumPaths(nums []int) int {
	var sum int
	for i, multi := len(nums)-1, 1; i >= 0; i-- {
		sum += nums[i] * multi
		multi *= 10
	}
	return sum
}

func sumNumbersRe(root *TreeNode, localPaths []int) {

	if root == nil {
		return
	}

	localPaths = append(localPaths, root.Val)
	if root.Left == nil && root.Right == nil {
		pathNumSum += sumPaths(localPaths)
		return
	}

	l := len(localPaths)
	sumNumbersRe(root.Left, localPaths)
	sumNumbersRe(root.Right, localPaths)
	localPaths = localPaths[:l]
}

func sumNumbers(root *TreeNode) int {

	pathNumSum = 0
	sumNumbersRe(root, make([]int, 0))
	return pathNumSum
}

var solvePathMapping [][]bool
var solveVisitedMapping [][]int

// 从（i，j）出发，看是否可以突围
func solveRe(board [][]byte, i, j int) bool {

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return true
	}

	if solvePathMapping[i][j] {
		return false
	}

	if solveVisitedMapping[i][j] >= 0 {
		return solveVisitedMapping[i][j] == 1
	}

	solvePathMapping[i][j] = true
	if !solveRe(board, i-1, j) && !solveRe(board, i+1, j) && !solveRe(board, i, j-1) && !solveRe(board, i, j+1) {
		solvePathMapping[i][j] = false
		return false
	}
	solvePathMapping[i][j] = false
	solveVisitedMapping[i][j] = 1
	return true
}

func solve(board [][]byte) {

	if len(board) <= 0 {
		return
	}

	solvePathMapping = make([][]bool, 0, len(board))
	solveVisitedMapping = make([][]int, 0, len(board))
	for i := 0; i < len(board); i++ {
		solvePathMapping = append(solvePathMapping, make([]bool, len(board[0])))
		solveVisitedMapping = append(solveVisitedMapping, make([]int, len(board[0])))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			solveVisitedMapping[i][j] = -1
			if board[i][j] == 88 {
				solveVisitedMapping[i][j] = 0
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if solveRe(board, i, j) {
				solveVisitedMapping[i][j] = 1
			} else {
				solveVisitedMapping[i][j] = 0
				board[i][j] = 88
			}
		}
	}
}
