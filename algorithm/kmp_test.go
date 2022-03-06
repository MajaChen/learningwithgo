package algorithm

import "testing"

func TestIndexKMP(t *testing.T) {

	nums1, nums2 := []int{7, 5, 6, 1, 3, 5, 6, 7}, []int{3, 5, 6, 7}
	i := IndexKMP(nums1, nums2, 1)
	t.Error(i)
}
