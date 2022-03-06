package leetcode

func generateTreesRe(nums []int) []*TreeNode {
	if len(nums) == 0 {
		return []*TreeNode{nil}
	}

	roots := make([]*TreeNode, 0)
	for i := 0; i < len(nums); i++ {
		leftChildren := generateTreesRe(nums[:i])
		rightChildren := generateTreesRe(nums[i+1:])
		for _, leftChild := range leftChildren {
			for _, rightChild := range rightChildren {
				root := &TreeNode{Val: nums[i]}
				root.Left = leftChild
				root.Right = rightChild
				roots = append(roots, root)
			}
		}
	}
	return roots
}

func generateTrees(n int) []*TreeNode {
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}
	return generateTreesRe(nums)
}

var treeNumsMapping []int

func numTreesRe(nums []int) int {

	var i int
	if i = len(nums); treeNumsMapping[i] > 0 {
		return treeNumsMapping[i]
	}

	var count int
	for i := 0; i < len(nums); i++ {
		leftCount := numTreesRe(nums[:i])
		rightCount := numTreesRe(nums[i+1:])
		count += leftCount * rightCount
	}

	treeNumsMapping[i] = count
	return count
}

func numTrees(n int) int {

	treeNumsMapping = make([]int, n+1)
	treeNumsMapping[0] = 1

	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}
	return numTreesRe(nums)
}

var traverseNums []int

func inorderTraverse(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if !inorderTraverse(root.Left) {
		return false
	}

	if len(traverseNums) > 0 && root.Val < traverseNums[len(traverseNums)-1] {
		return false
	}
	traverseNums = append(traverseNums, root.Val)

	if !inorderTraverse(root.Right) {
		return false
	}

	return true
}

func isValidBST(root *TreeNode) bool {
	traverseNums = make([]int, 0)
	return inorderTraverse(root)
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var q Queue = Queue{elems: make([]interface{}, 0)}
	q.Push(root)
	count, tmpRes := 0, make([]int, 0)
	for !q.IsEmpty() {
		size := q.Size()
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
		count = 0
		res = append(res, tmpRes)
		//tmpRes = tmpRes[:0] 禁止内存复用，否则会出错
		tmpRes = make([]int, 0, q.Size())
	}
	return res
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; {
		num := nums[i]
		nums[i] = nums[j]
		nums[j] = num
		i++
		j--
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var q Queue = Queue{elems: make([]interface{}, 0)}
	q.Push(root)

	for layer := 1; !q.IsEmpty(); layer++ {

		size, count, ares := q.Size(), 0, make([]int, 0)
		for count < size {
			node := q.Pop().(*TreeNode)
			ares = append(ares, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
			count++
		}

		if layer%2 == 0 {
			reverse(ares)
		}
		res = append(res, ares)
	}
	return res
}
