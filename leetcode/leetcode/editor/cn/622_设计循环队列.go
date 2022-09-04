package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
type MyCircularQueue struct {
	arr []int
	h   int
	t   int
	c   int
	s   int
}

func MyCircularQueueConstructor(k int) MyCircularQueue {
	return MyCircularQueue{
		arr: make([]int, k),
		s:   k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.c == this.s {
		return false
	}
	this.arr[this.h] = value
	this.h = (this.h + 1) % this.s
	this.c++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.c == 0 {
		return false
	}
	this.c--
	this.t = (this.t + 1) % this.s
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.c == 0 {
		return -1
	}
	return this.arr[this.t]
}

func (this *MyCircularQueue) Rear() int {
	if this.c == 0 {
		return -1
	}
	return this.arr[(this.h-1+this.s)%this.s]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.c == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.c == this.s
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := CBTInserterConstructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)
