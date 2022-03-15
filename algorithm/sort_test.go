package algorithm

import "testing"

func TestListInsertSort(t *testing.T) {
	head := &floatListNode{}
	three := &floatListNode{val: 3}
	five := &floatListNode{val: 5}
	four := &floatListNode{val: 4}
	head.next = three
	three.next = five
	five.next = four

	bucket := &bucket{head: head, tail: head}
	bucket.sort()
	t.Error("test failed")
}

func TestBucketSort(t *testing.T) {
	nums := []float32{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68}
	nums = BucketSort(nums)
	t.Errorf("nums is %v", nums)
}

func TestCountSort(t *testing.T) {
	nums := []int{2, 5, 3, 0, 2, 3, 0, 3}
	nums = CountSort(nums, 5)
	t.Errorf("nums is %v", nums)
}

func TestCountSort2(t *testing.T) {
	nums := []int{2}
	nums = CountSort(nums, 2)
	t.Errorf("nums is %v", nums)
}

func TestQuickSort(t *testing.T) {
	nums := []int{2, 4, 1, 3}
	nums = QuickSort(nums)
	t.Errorf("nums is %v", nums)
}

func TestQuickSort2(t *testing.T) {
	nums := []int{49, 38, 65, 97, 76, 13, 27, 49}
	nums = QuickSort(nums)
	t.Errorf("nums is %v", nums)
}

func TestMergeSort(t *testing.T) {
	nums := []int{3, 2, 4, 1}
	nums = MergeSort(nums)
	t.Errorf("nums is %v", nums)
}

func TestMergeSort2(t *testing.T) {
	nums := []int{49, 38, 65, 97, 76, 13, 27, 49}
	nums = MergeSort(nums)
	t.Errorf("nums is %v", nums)
}

func TestBuildHeap(t *testing.T) {
	nums := []int{0, 5, 7, 1}
	nums = buildHeap(nums, 3)
	t.Errorf("nums is %v", nums)
}

func TestBuildHeap2(t *testing.T) {
	nums := []int{0, 13, 38, 27, 49, 76, 65, 49, 97}
	nums = buildHeap(nums, 8)
	t.Errorf("nums is %v", nums)
}

func TestHeapSort(t *testing.T) {
	nums := []int{0, 13, 38, 27, 49, 76, 65, 49, 97}
	nums = HeapSort(nums, 8)
	t.Errorf("nums is %v", nums)
}

func TestTopK(t *testing.T) {
	nums := []int{-1, 5, 4, 8, 2, 6, 0, 7}
	top3 := TopK(nums, 3)
	t.Errorf("top3 is %v", top3)
}
