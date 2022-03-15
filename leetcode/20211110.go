package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	k, index := 0, m
	for i := 0; i < len(nums2); i++ {
		// 寻找插入位置
		j := k
		for ; j < index && nums1[j] <= nums2[i]; j++ {
		}
		k = j
		// 从j开始向后移动元素
		for j = index - 1; j >= k; j-- {
			nums1[j+1] = nums1[j]
		}
		// 插入元素
		nums1[k] = nums2[i]
		index++
	}
}
