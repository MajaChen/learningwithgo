package common

import "testing"

func TestHeap(t *testing.T) {

	var c = func(x, y int) bool { return x > y }
	heap := HeapConsturctor(c)
	heap.Add(13)
	heap.Add(2)
	heap.Add(44)
	heap.Add(3)
	heap.Add(30)
	heap.Add(20)
	heap.Add(1)
	heap.Add(23)
	heap.Add(9)
	t.Error(heap.Top())
}
