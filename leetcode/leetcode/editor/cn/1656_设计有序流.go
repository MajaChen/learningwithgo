package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
type OrderedStream struct {
	ss  []string
	ptr int
}

func OrderedStreamConstructor(n int) OrderedStream {
	return OrderedStream{
		ss:  make([]string, n+1),
		ptr: 1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.ss[idKey] = value
	var result []string
	if this.ss[this.ptr] != "" {
		var i int
		for i = this.ptr; i < len(this.ss) && this.ss[i] != ""; i++ {
		}
		result = this.ss[this.ptr:i]
		this.ptr = i

	}
	return result
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := MyCircularQueueConstructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
