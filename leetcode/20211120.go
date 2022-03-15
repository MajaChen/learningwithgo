package leetcode

func twoSumV3(nums []int, target int) []int {
	for l, r := 0, len(nums)-1; l < r; {
		if sum := nums[l] + nums[r]; sum == target {
			return []int{l + 1, r + 1}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return []int{}
}

// step1:首选判断nums[mid]是否等于target,若是直接返回；
// step2:若不是，首先判断nums[mid]是否大于等于nums[left],若是则说明左边是一个单调递增的序列，那么进一步对左序列展开判断：
// step3:若target在左序列中，保留左，舍弃右；若不在左序列中，保留右，舍弃左；
// 重新回到step2，若否则说明右边是一个单调递增的序列，那么进一步对右序列展开类似判断
func search(nums []int, target int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if nums[left] <= target && nums[mid] > target { // 左边选中
				right = mid - 1
			} else { // 左边舍弃
				left = mid + 1
			}
		} else {
			if nums[right] >= target && nums[mid] < target { // 右边选中
				left = mid + 1
			} else { // 右边舍弃
				right = mid - 1
			}
		}
	}
	return -1
}
