package leetcode

import "math"

type MinStack struct {
	elems  []int64
	index  int
	minVal int64
}

// 规定新、旧max的差值大于0，使max与val之间的差值大于0
// 规定新、旧min的差值大于0，使min与val之间的差值大于0

func MinStackConstructor() MinStack {
	return MinStack{elems: make([]int64, 0), minVal: math.MaxInt32, index: -1}
}

var padding int64

func (this *MinStack) addElem(elem int64) {

	for this.index++; this.index > len(this.elems)-1; {
		this.elems = append(this.elems, padding)
	}
	this.elems[this.index] = elem
}

func (this *MinStack) Push(val int) {
	val64 := int64(val)
	if val64 <= this.minVal {
		this.addElem(this.minVal)
		this.addElem(0)
		this.minVal = val64
		return
	}

	this.addElem(val64 - this.minVal)
	return
}

func (this *MinStack) Pop() {
	if this.elems[this.index] == 0 {
		this.index--
		this.minVal = this.elems[this.index]
		this.index--
	} else {
		this.index--
	}
}

func (this *MinStack) Top() int {
	return int(this.minVal + this.elems[this.index])
}

func (this *MinStack) GetMin() int {
	return int(this.minVal)
}

type MaxStack struct {
	elems  []int64
	index  int
	maxVal int64
}

func MaxStackConstructor() MaxStack {
	return MaxStack{elems: make([]int64, 0), maxVal: math.MinInt32, index: -1}
}

func (this *MaxStack) Push(val int) {
	val64 := int64(val)
	if val64 >= this.maxVal {
		diff := val64 - this.maxVal
		this.index++
		this.elems = append(this.elems, diff)
		this.maxVal = val64
	}
	this.index++
	this.elems = append(this.elems, this.maxVal-val64)
	return
}

func (this *MaxStack) Pop() {
	if this.elems[this.index] == 0 {
		this.index--
		diff := this.elems[this.index]
		this.maxVal = this.maxVal - diff
		this.index--
	} else {
		this.index--
	}
	if this.index < 0 {
		this.maxVal = math.MinInt32
	}
	this.elems = this.elems[:this.index+1]
}

func (this *MaxStack) Top() int {
	return int(this.maxVal - this.elems[this.index])
}

func (this *MaxStack) GetMax() int {
	return int(this.maxVal)
}
