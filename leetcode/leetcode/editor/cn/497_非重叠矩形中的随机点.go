package leetcode

import (
	"math/rand"
	"sort"
	"time"
)

type Solution struct {
	n        int
	counters []int
	rects    [][]int
	sum      int
}

func Constructor(rects [][]int) Solution {
	n, counters := len(rects), make([]int, len(rects)+1)
	for i := 1; i <= n; i++ {
		counters[i] = counters[i-1] + (rects[i-1][2]-rects[i-1][0]+1)*(rects[i-1][3]-rects[i-1][1]+1)
	}
	return Solution{
		n:        n,
		counters: counters[1:],
		rects:    rects,
		sum:      counters[n],
	}
}

func rangeRand(m, n int) int {
	rand.Seed(time.Now().UnixNano())
	return m + rand.Intn(n-m+1)
}

func (this *Solution) Pick() []int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(this.sum + 1)
	i := sort.Search(len(this.counters), func(x int) bool {
		return this.counters[x] >= r
	})

	rect := this.rects[i]
	x, y := rangeRand(rect[0], rect[2]), rangeRand(rect[1], rect[3])
	return []int{x, y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := FinderConstructor(rects);
 * param_1 := obj.Pick();
 */
