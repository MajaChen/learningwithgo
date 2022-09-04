package leetcode

import ."learning/common"

//leetcode submit region begin(Prohibit modification and deletion)

func largestComponentSize(nums []int) int {
	set := InitDisjointSet()
	for _, num := range nums {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				set.Union(num, i)
				set.Union(num, num/i)
			}
		}
	}

	counters, ans := make(map[int]int), 0
	for _, num := range nums {
		root := set.Find(num).(int)
		counters[root] = counters[root] + 1
		if counters[root] > ans {
			ans = counters[root]
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
