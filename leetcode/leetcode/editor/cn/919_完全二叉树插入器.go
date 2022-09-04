package leetcode

import . "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)
type CBTInserter struct {
	root  *TreeNode
	queue *Queue
}

func CBTInserterConstructor(root *TreeNode) CBTInserter {
	inserter := CBTInserter{
		queue: &Queue{make([]interface{}, 0)},
	}
	inserter.internalInsert(root)
	return inserter
}

func (this *CBTInserter) internalInsert(root *TreeNode) {
	q := Queue{make([]interface{}, 0)}
	q.Push(root)
	this.root = root
	for !q.IsEmpty() {
		node := q.Pop().(*TreeNode)
		if node.Left != nil {
			q.Push(node.Left)
		}
		if node.Right != nil {
			q.Push(node.Right)
		}
		if node.Left == nil || node.Right == nil {
			this.queue.Push(node)
		}
	}
}

func (this *CBTInserter) Insert(val int) int {
	node := &TreeNode{Val: val}
	if t := this.queue.GetTop().(*TreeNode); t.Left == nil {
		t.Left = node
		this.queue.Push(node)
		return t.Val
	} else if t.Right == nil {
		t.Right = node
		this.queue.Push(node)
		return t.Val
	} else {
		this.queue.Pop()
		return this.Insert(val)
	}
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := MyCalendarTwoConstructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
//leetcode submit region end(Prohibit modification and deletion)
