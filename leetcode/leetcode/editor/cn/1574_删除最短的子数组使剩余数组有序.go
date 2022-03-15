package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

func findLengthOfShortestSubarray(arr []int) int {

	left, right, n, ans := 0, len(arr)-1, len(arr), len(arr)-1
	for ; left+1 < n && arr[left] <= arr[left+1]; left++ {
	}
	for ; right-1 >= 0 && arr[right-1] <= arr[right]; right-- {
	}

	if left == n-1 || right == 0 {
		return 0
	}
	ans = int(math.Min(float64(n-left-1), float64(right)))

	for i, j := 0, right; i <= left && j < n; {
		if arr[i] <= arr[j] {
			if j-i-1 < ans {
				ans = j - i - 1
			}
			i++
		} else {
			j++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
