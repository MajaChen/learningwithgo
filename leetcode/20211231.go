package leetcode

import "math"

func adjust_Big(heap []int, leftBound, rightBound int) {

	for i := leftBound; i*2 <= rightBound; {
		maxIndex := i * 2
		if i*2+1 <= rightBound && heap[i*2+1] > heap[i*2] {
			maxIndex += 1
		}

		if heap[maxIndex] > heap[i] {
			swap(maxIndex, i, heap)
		} else {
			break
		}
		i = maxIndex
	}
}

func adjust_Small(heap []int, leftBound, rightBound int) {

	for i := leftBound; i*2 <= rightBound; {
		minIndex := i * 2
		if i*2+1 <= rightBound && heap[i*2+1] < heap[i*2] {
			minIndex += 1
		}

		if heap[minIndex] < heap[i] {
			swap(minIndex, i, heap)
		} else {
			break
		}
		i = minIndex
	}
}

func buildHeap(heap []int, N int) {
	for i := N / 2; i > 0; i-- {
		adjust_Small(heap, i, N)
	}
}

type HeapSet struct {
	elements []int
	length   int
}

func BuildHeapSet() *HeapSet {
	return &HeapSet{elements: make([]int, 1)}
}

func (set *HeapSet) Push(val int) {

	set.elements = append(set.elements, val)
	set.length++
	buildHeap(set.elements, set.length)
}

func (set *HeapSet) Pop() int {

	val := set.elements[1]
	swap(1, set.length, set.elements)
	set.length--
	set.elements = set.elements[:set.length+1]
	adjust_Small(set.elements, 1, set.length)
	return val
}

func getKthMagicNumber(k int) int {

	mapping := make(map[int]bool)
	mapping[1] = true
	set := BuildHeapSet()
	set.Push(1)

	var cur int
	for i := 1; i <= k; i++ {
		cur = set.Pop()
		if _, ok := mapping[cur*3]; !ok {
			set.Push(cur * 3)
			mapping[cur*3] = true
		}
		if _, ok := mapping[cur*5]; !ok {
			set.Push(cur * 5)
			mapping[cur*5] = true
		}
		if _, ok := mapping[cur*7]; !ok {
			set.Push(cur * 7)
			mapping[cur*7] = true
		}
		mapping[cur] = true
	}
	return cur
}

func findSquare(matrix [][]int) []int {

	if len(matrix) <= 0 {
		return []int{}
	}

	mapping, m, n, maxLen := make([][]int, 0), len(matrix), len(matrix[0]), math.MinInt32
	for i := 0; i < m; i++ {
		mapping = append(mapping, make([]int, n))
	}

	var r, c, size int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				var l1, l2, l3 int
				if i-1 >= 0 {
					l1 = mapping[i-1][j]
				}
				if j-1 >= 0 {
					l2 = mapping[i][j-1]
				}
				if i-1 >= 0 && j-1 >= 0 {
					l3 = mapping[i-1][j-1]
				}
				if mapping[i][j] = min(l1, l2, l3) + 1; mapping[i][j] > maxLen {
					l := mapping[i][j]
					r, c, size = i-l, j-l, l
					maxLen = l
				}
			}
		}
	}

	if maxLen == math.MinInt32 {
		return []int{}
	}
	return []int{r, c, size}
}
