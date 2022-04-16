package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

var elemIndexMapping map[int]int

func buildTreeRe(preorder, inorder []int, offset int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	i := elemIndexMapping[preorder[0]] - offset
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTreeRe(preorder[1:i+1], inorder[:i], offset),
		Right: buildTreeRe(preorder[i+1:], inorder[i+1:], offset+i+1),
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	elemIndexMapping = make(map[int]int)
	for index, elem := range inorder {
		elemIndexMapping[elem] = index
	}

	return buildTreeRe(preorder, inorder, 0)
}

//leetcode submit region end(Prohibit modification and deletion)
