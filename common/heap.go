package common

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

const DefaultHeapLength = 64

type Heap struct {
	elems []int
	l     int
	c     func(int, int) bool
}

func HeapConsturctor(c func(int, int) bool) *Heap {
	return &Heap{elems: make([]int, 1, DefaultHeapLength), c: c}
}

func (heap *Heap) Add(x int) {
	heap.elems = append(heap.elems, x)
	heap.l += 1
	heap.siftUp(1, heap.l)
}

func (heap *Heap) Poll() int {
	elem := heap.elems[1]
	swap(heap.elems, 1, heap.l)
	heap.elems = heap.elems[:heap.l]
	heap.l -= 1
	heap.siftDown(1, heap.l)
	return elem
}

func (heap Heap) Top() int {
	return heap.elems[1]
}

func (heap *Heap) siftDown(l, r int) {
	for i := l; i*2 <= r; {
		j := i * 2
		if i*2+1 <= r && heap.c(heap.elems[i*2+1], heap.elems[i*2]) {
			j += 1
		}

		if heap.c(heap.elems[j], heap.elems[i]) {
			swap(heap.elems, j, i)
		} else {
			break
		}
		i = j
	}
}

func (heap *Heap) siftUp(l, r int) {
	for r/2 >= l {
		if heap.c(heap.elems[r], heap.elems[r/2]) {
			swap(heap.elems, r, r/2)
			r /= 2
		} else {
			break
		}
	}
}
