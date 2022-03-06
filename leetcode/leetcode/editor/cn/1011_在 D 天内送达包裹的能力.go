package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

func canFinish(weights []int, days, capacity int) bool {

	var d int
	var j int
	for i := 0; i < len(weights) && d <= days; {
		if j+weights[i] > capacity {
			d++
			j = 0
		} else {
			j += weights[i]
			i++
		}
	}

	if j > 0 {
		d++
	}
	return d <= days
}
func shipWithinDays(weights []int, days int) int {

	var sum, max int
	max = math.MinInt32
	for i := 0; i < len(weights); i++ {
		sum += weights[i]
		if weights[i] > max {
			max = weights[i]
		}
	}

	var l, r int
	for l, r = max, sum; l < r; {
		if mid := (l + r) / 2; canFinish(weights, days, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

//leetcode submit region end(Prohibit modification and deletion)
