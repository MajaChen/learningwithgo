package leetcode

import (
	"sort"
	"strings"
)

func triangleNumber(nums []int) int {

	if len(nums) <= 2 {
		return 0
	}

	count := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 0 { // 这里就应该加continue，也必须加continue，否则会得到荒唐的结果
			continue
		}
		for j, k := i+1, i+2; j < len(nums)-1; j++ {
			for ; k < len(nums) && nums[i]+nums[j] > nums[k]; k++ {
			}
			count += k - j - 1
		}
	}

	return count
}

func repeatedStringMatch(a string, b string) int {

	if a == b {
		return 1
	}

	mapping := make(map[string]bool)
	for i := 0; i < len(a); i++ {
		mapping[a[i:i+1]] = true
	}
	for j := 0; j < len(b); j++ {
		if _, ok := mapping[b[j:j+1]]; !ok {
			return -1
		}
	}

	for str, count, length := a, 1, len(a); length <= len(b)+2*len(a); length += len(a) {
		if strings.Index(str, b) >= 0 {
			return count
		}
		str += a
		count++
	}

	return -1
}
