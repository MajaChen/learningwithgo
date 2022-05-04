package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

const DEFAULT_STACK_LENGTH = 64

type MinStack struct {
	elems   []int
	minElem int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{elems: make([]int, 0, DEFAULT_STACK_LENGTH), minElem: math.MaxInt32}
}

func (this *MinStack) Push(x int) {
	if x <= this.minElem {
		this.elems = append(this.elems, this.minElem)
		this.elems = append(this.elems, 0)
		this.minElem = x
	} else {
		this.elems = append(this.elems, x-this.minElem)
	}
}

func (this *MinStack) Pop() {
	if this.elems[len(this.elems)-1] == 0 {
		this.minElem = this.elems[len(this.elems)-2]
		this.elems = this.elems[:len(this.elems)-2]
	} else {
		this.elems = this.elems[:len(this.elems)-1]
	}
}

func (this *MinStack) Top() int {
	if this.elems[len(this.elems)-1] == 0 {
		return this.minElem
	} else {
		return this.elems[len(this.elems)-1] + this.minElem
	}
}

func (this *MinStack) Min() int {
	return this.minElem
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := CQueueConstructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
//leetcode submit region end(Prohibit modification and deletion)
