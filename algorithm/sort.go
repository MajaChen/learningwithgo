package algorithm

import "math"

// 堆结点“沉淀”
// 基于递归
func heapNodeSink(heap []int, leftBound, rightBound int) []int {
	// 叶结点直接返回
	if leftBound*2 >= rightBound {
		return heap
	}

	// 非叶结点交换左右子节点的最大值
	left, right, maxIndex := math.MinInt32, math.MinInt32, leftBound*2
	if leftBound*2 < rightBound {
		left = heap[leftBound*2]
	}
	if leftBound*2+1 < rightBound {
		right = heap[leftBound*2+1]
	}

	if left < right {
		maxIndex += 1
	}
	if heap[maxIndex] > heap[leftBound] {
		swap(maxIndex, leftBound, heap)
	}

	return heap
}

func buildHeap(nums []int) []int {
	for left, right := (len(nums)/2)-1, len(nums)-1; left >= 0; left-- {
		nums = heapNodeSink(nums, left, right)
	}
	return nums
}

// 堆排序
func HeapSort(nums []int) []int {
	heap := buildHeap(nums)
	for leftBound, rightBound := 0, len(nums); rightBound >= leftBound; rightBound-- {
		swap(leftBound, rightBound, heap)
		rightBound--
		heapNodeSink(heap, leftBound, rightBound)
	}
	return heap
}

func MergeSortRe(nums []int, left, right int) []int {
	if left == right {
		return nums
	}

	mid := (left + right) / 2
	lsubNums := MergeSortRe(nums, left, mid)
	rsubNums := MergeSortRe(nums, mid+1, right)
	lsubNums = append(lsubNums, math.MaxInt32)
	rsubNums = append(rsubNums, math.MaxInt32)

	for i, j := 0, 0; i < len(lsubNums) && j < len(rsubNums); {
		if lsubNums[i] < rsubNums[j] {
			nums[i+j] = lsubNums[i]
			i++

		} else {
			nums[i+j] = rsubNums[j]
			j++
		}
	}

	return nums[:len(nums)-1]
}

// 归并排序
func MergeSort(nums []int) []int {
	return MergeSortRe(nums, 0, len(nums)-1)
}

// 快速排序
