package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
type MyCalendar struct {
	full    map[int]bool
	partial map[int]bool
}

func MyCalendarConstructor() MyCalendar {
	return MyCalendar{full: make(map[int]bool), partial: make(map[int]bool)}
}

func (c *MyCalendar) queryMyCalendar(s, e, l, r, pos int) bool {
	if e < l || s > r {
		return false
	}
	if s <= l && e >= r {
		return c.partial[pos]
	}
	if c.full[pos] {
		return true
	}

	m := (l + r) >> 1
	return c.queryMyCalendar(s, e, l, m, 2*pos) || c.queryMyCalendar(s, e, m+1, r, 2*pos+1)
}

func (c *MyCalendar) updateMyCalendar(s, e, l, r, pos int) {
	if e < l || s > r {
		return
	}
	if s <= l && e >= r {
		c.full[pos] = true
		c.partial[pos] = true
		return
	}
	m := (l + r) >> 1
	c.updateMyCalendar(s, e, l, m, 2*pos)
	c.updateMyCalendar(s, e, m+1, r, 2*pos+1)
	c.partial[pos] = true
	return
}

func (this *MyCalendar) Book(start int, end int) bool {
	if this.queryMyCalendar(start, end-1, 0, 1e9, 1) {
		return false
	}

	this.updateMyCalendar(start, end-1, 0, 1e9, 1)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := EncodeConstructor();
 * param_1 := obj.Book(start,end);
 */
//leetcode submit region end(Prohibit modification and deletion)
