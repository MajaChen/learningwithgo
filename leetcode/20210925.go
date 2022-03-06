package leetcode

import "math"

func increasingTriplet3(nums []int) bool {
	nums1, nums2 := math.MaxInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums1 {
			nums1 = nums[i]
		} else if nums[i] < nums2 {
			nums2 = nums[i]
		} else {
			return true
		}
	}
	return false
}

func increasingTriplet2(nums []int) int {
	nums1, nums2, nums3 := math.MaxInt32, math.MaxInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums1 {
			nums1 = nums[i]
		} else if nums[i] < nums2 {
			nums2 = nums[i]
		} else if nums[i] < nums3 {
			nums3 = nums[i]
		} else {
			nums1 = nums2
			nums2 = nums3
			nums3 = nums[i]
		}
	}
	if nums1 == math.MaxInt32 || nums2 == math.MaxInt32 || nums3 == math.MaxInt32 {
		return 0
	} else {
		return nums1 + nums2 + nums3
	}
}
