package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

type pair struct {
	max int
	min int
}

type MyCalendarTwo map[int]pair

func Constructor() MyCalendarTwo {
	return make(map[int]pair)
}

func (c MyCalendarTwo) update(s, e, l, r, v, idx int) {
	if e < l || s > r {
		return
	}
	if s <= l && e >= r {
		p := c[idx]
		p.min += v
		p.max += v
		c[idx] = p
		return
	}
	m := (l + r) >> 1
	c.update(s, e, l, m, v, 2*idx)
	c.update(s, e, m+1, r, v, 2*idx+1)
	p := c[idx]
	p.max = p.min + max(c[2*idx].max, c[2*idx+1].max)
	c[idx] = p
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	c := *this
	if this.update(start, end-1, 0, 1e9, 1, 1); c[1].max >= 3 {
		this.update(start, end-1, 0, 1e9, -1, 1)
		return false
	}
	return true
}

/**
* Your MyCalendarTwo object will be instantiated and called as such:
* obj := MovingAverageConstructor();
* param_1 := obj.Book(start,end);
 */
//leetcode submit region end(Prohibit modification and deletion)
