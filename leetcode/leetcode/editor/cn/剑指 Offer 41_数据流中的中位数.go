package leetcode

import . "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)
type MedianFinder struct {
	bh *Heap
	sh *Heap
	n  int
}

func Constructor() MedianFinder {
	return MedianFinder{
		bh: HeapConsturctor(func(x, y int) bool { return x > y }),
		sh: HeapConsturctor(func(x, y int) bool { return x < y })}
}

func (this *MedianFinder) AddNum(num int) {
	if this.n%2 == 0 {
		this.sh.Add(num)
		this.bh.Add(this.sh.Poll())
	} else {
		this.bh.Add(num)
		this.sh.Add(this.bh.Poll())
	}
	this.n += 1
	return
}

func (this *MedianFinder) FindMedian() float64 {
	if this.n%2 == 0 {
		return (float64(this.bh.Top()) + float64(this.sh.Top())) / 2
	} else {
		return float64(this.bh.Top())
	}
}

//leetcode submit region end(Prohibit modification and deletion)
