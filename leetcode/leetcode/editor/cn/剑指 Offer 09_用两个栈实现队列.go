package leetcode

import . "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)

type CQueue struct {
	s1, s2 Stack
}

func CQueueConstructor() CQueue {
	return CQueue{s1: Stack{}, s2: Stack{}}
}

func (this *CQueue) AppendTail(value int) {
	this.s1.Push(value)
}

func (this *CQueue) DeleteHead() int {

	if this.s1.IsEmpty() && this.s2.IsEmpty() {
		return -1
	}

	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}

	return this.s2.Pop().(int)
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := CQueueConstructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
//leetcode submit region end(Prohibit modification and deletion)
