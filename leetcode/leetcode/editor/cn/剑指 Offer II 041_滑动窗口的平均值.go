package leetcode

import "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)

type MovingAverage struct {
	size int
	len  int
	sum  int
	head *common.ListNode
	tail *common.ListNode
}

func MovingAverageConstructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	newNode := &common.ListNode{Val: val}
	this.len++
	this.sum += val
	if this.tail == nil {
		this.head = newNode
		this.tail = newNode
	} else {
		this.tail.Next = newNode
		this.tail = newNode
		if this.len > this.size {
			this.len = this.size
			this.sum -= this.head.Val
			this.head = this.head.Next
		}
	}
	return float64(this.sum) / float64(this.len)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := MovingAverageConstructor(size);
 * param_1 := obj.Next(val);
 */
//leetcode submit region end(Prohibit modification and deletion)
