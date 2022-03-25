package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

type CustomStack struct {
	elems   []int
	maxSize int
}

func CustomStackConstructor(maxSize int) CustomStack {
	return CustomStack{elems: make([]int, 0), maxSize: maxSize}
}

func (this *CustomStack) Push(x int) {

	if len(this.elems) < this.maxSize {
		this.elems = append(this.elems, x)
	}
}

func (this *CustomStack) Pop() int {

	if len(this.elems) == 0 {
		return -1
	}

	elem := this.elems[len(this.elems)-1]
	this.elems = this.elems[:len(this.elems)-1]
	return elem
}

func (this *CustomStack) Increment(k int, val int) {

	for i := 0; i < k && i < len(this.elems); i++ {
		this.elems[i] += val
	}
}

//leetcode submit region end(Prohibit modification and deletion)
