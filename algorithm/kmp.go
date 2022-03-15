package algorithm

var next []int

// 两大要点：
// 1.字符串的偏移，值传递
// 2.建立next数组过程中的回溯
func IndexKMP(nums1, nums2 []int, pos int) int {

	next = make([]int, len(nums2))
	getNextVal(nums2, next)

	var i, j int
	for i, j = pos, 1; i <= nums1[0] && j <= nums2[0]; {
		if j == 0 || nums1[i] == nums2[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j > nums2[0] {
		return i - nums2[0]
	} else {
		return 0
	}
}

func getNextVal(nums, next []int) {

	next[1] = 0
	var i, j int
	for i, j = 1, 0; i < nums[0]; {
		if j == 0 || nums[i] == nums[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
}
