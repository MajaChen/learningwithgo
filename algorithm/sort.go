package algorithm

import "math"

// 堆结点“沉淀”
// 基于递归
func heapNodeSink(heap []int, leftBound, rightBound int) {
	// 叶结点直接返回
	if leftBound*2 > rightBound {
		return
	}

	// 非叶结点交换左右子节点的最大值
	left, right, maxIndex := math.MinInt32, math.MinInt32, leftBound*2
	if leftBound*2 <= rightBound {
		left = heap[leftBound*2]
	}
	if leftBound*2+1 <= rightBound {
		right = heap[leftBound*2+1]
	}

	if left < right {
		maxIndex += 1
	}
	if heap[maxIndex] > heap[leftBound] {
		swap(maxIndex, leftBound, heap)
	}

	if maxIndex%2 == 0 {
		heapNodeSink(heap, leftBound*2, rightBound)
	} else {
		heapNodeSink(heap, leftBound*2+1, rightBound)
	}
}

func buildHeap(nums []int, N int) []int {
	for left, right := int(N/2), N; left >= 1; left-- {
		heapNodeSink(nums, left, right)
	}
	return nums
}

const RANDOMNUMBER = 0

// HeapSort 堆排序
func HeapSort(nums []int, N int) []int {
	heap := buildHeap(nums, N)
	for leftBound, rightBound := 1, N; rightBound > leftBound; {
		swap(leftBound, rightBound, heap)
		rightBound--
		heapNodeSink(heap, leftBound, rightBound)
	}
	return heap
}

// nums 左右两边的元素已经有序，
// 将数组nums左右半边的的元素进行归并,并存入nums left到right范围
func merge(nums []int, left, mid, right int) {
	// 必须新建数组
	newNums := make([]int, right-left+1)

	i, j, count := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			newNums[count] = nums[i]
			i++
		} else {
			newNums[count] = nums[j]
			j++
		}
		count++
	}

	for i <= mid {
		newNums[count] = nums[i]
		count++
		i++
	}

	for j <= right {
		newNums[count] = nums[j]
		count++
		j++
	}

	// 将newNums元素拷贝回原数组
	for i := 0; i < len(newNums); i++ {
		nums[left+i] = newNums[i]
	}
}

// 将数组nums从范围left到right的元素排好序，存入nums的left到right范围
func mergeSortRe(nums []int, left, right int) {
	if left == right {
		return
	}

	mid := (left + right) / 2
	mergeSortRe(nums, left, mid)
	mergeSortRe(nums, mid+1, right)
	merge(nums, left, mid, right)
}

// MergeSort 归并排序
func MergeSort(nums []int) []int {
	mergeSortRe(nums, 0, len(nums)-1)
	return nums
}

func findPivot(nums []int, left, right int) ([]int, int) {

	pivot := nums[left]
	low, high, pivotIndex := left, right, left
	for high > low {
		for ; nums[high] >= pivot && high > low; high-- {
		}
		nums[low] = nums[high]
		pivotIndex = high

		for ; nums[low] <= pivot && high > low; low++ {
		}
		nums[high] = nums[low]
		pivotIndex = low
	}

	nums[pivotIndex] = pivot
	return nums, pivotIndex
}

func quickSortRe(nums []int, left, right int) []int {

	if left >= right {
		return nums
	}

	var pivotIndex int
	nums, pivotIndex = findPivot(nums, left, right)
	nums = quickSortRe(nums, left, pivotIndex-1)
	nums = quickSortRe(nums, pivotIndex+1, right)
	return nums
}

// QuickSort 快速排序
func QuickSort(nums []int) []int {
	return quickSortRe(nums, 0, len(nums)-1)
}

// CountSort 计数排序
func CountSort(nums []int, k int) []int {
	countNums := make([]int, k+1)
	for _, num := range nums {
		countNums[num]++
	}

	for i := 1; i < len(countNums); i++ {
		countNums[i] += countNums[i-1]
	}

	sortedNums := make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		sortedNums[countNums[nums[i]]] = nums[i]
		countNums[nums[i]] -= 1
	}
	return sortedNums[1:]
}

type floatListNode struct {
	val  float32
	next *floatListNode
}

type bucket struct {
	head *floatListNode
	tail *floatListNode
}

func (bucket *bucket) sort() {
	newHead := &floatListNode{}
	p, q := bucket.head.next, newHead
	for p != nil {
		pNext := p.next
		for q.next != nil && p.val > q.next.val {
			q = q.next
		}
		p.next = q.next
		q.next = p

		p = pNext
	}
	bucket.head = newHead
}

// BucketSort 桶排序
// nums中的整数在区间0-1上均匀分布
func BucketSort(nums []float32) []float32 {
	n := len(nums)
	buckets := make([]*bucket, n)
	for _, num := range nums {
		bucketIndex := int(num * float32(n))
		if buckets[bucketIndex] == nil {
			bucket := &bucket{}
			bucket.head = &floatListNode{}
			bucket.tail = bucket.head
			buckets[bucketIndex] = bucket
		}
		buckets[bucketIndex].tail.next = &floatListNode{val: num}
		buckets[bucketIndex].tail = buckets[bucketIndex].tail.next
	}

	k := 0
	for _, bucket := range buckets {
		if bucket != nil {
			bucket.sort()
			for p := bucket.head.next; p != nil; p = p.next {
				nums[k] = p.val
				k++
			}
		}
	}
	return nums
}
