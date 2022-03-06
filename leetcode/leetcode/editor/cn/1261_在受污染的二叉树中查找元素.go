package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

type FindElements struct {
	root *TreeNode
}

func buildFindTree(root *TreeNode, parentVal, addVal int) {

	if root == nil {
		return
	}

	root.Val = parentVal*2 + addVal
	buildFindTree(root.Left, root.Val, 1)
	buildFindTree(root.Right, root.Val, 2)
}

func FindTreeConstructor(root *TreeNode) FindElements {

	root.Val = 0
	buildFindTree(root.Left, root.Val, 1)
	buildFindTree(root.Right, root.Val, 2)
	return FindElements{root: root}
}

func find(root *TreeNode, target int) bool {

	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}

	return find(root.Left, target) || find(root.Right, target)
}

func (this *FindElements) Find(target int) bool {

	return find(this.root, target)
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := FindTreeConstructor(root);
 * param_1 := obj.Find(target);
 */
//leetcode submit region end(Prohibit modification and deletion)
