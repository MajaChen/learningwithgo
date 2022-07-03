package leetcode

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)
func diffWaysToCompute(expression string) []int {
	nums, opts := make([]int, 0, len(expression)), make([]string, 0, len(expression))
	for i := 0; i < len(expression); {
		if expression[i] >= 48 && expression[i] <= 57 {
			var j int
			for j = i; j < len(expression) && (expression[j] >= 48 && expression[j] <= 57); j++ {
			}
			num, _ := strconv.Atoi(expression[i:j])
			nums = append(nums, num)
			i = j
		} else {
			opts = append(opts, expression[i:i+1])
			i += 1
		}
	}

	dp, n := make([][][]int, 0, len(nums)), len(nums)
	for i := 0; i < n; i++ {
		dp = append(dp, make([][]int, n))
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < len(nums) && i+j-1 < len(nums); j++ {
			//[i, i+j-1]
			if l, r := j, i+j-1; l == r {
				dp[l][r] = []int{nums[l]}
			} else {
				var ns []int
				for t := l; t < r; t++ {
					nsl, nsr := dp[l][t], dp[t+1][r]
					for _, x := range nsl {
						for _, y := range nsr {
							if opts[t] == "+" {
								ns = append(ns, x+y)
							} else if opts[t] == "-" {
								ns = append(ns, x-y)
							} else {
								ns = append(ns, x*y)
							}
						}
					}
				}
				dp[l][r] = ns
			}
		}
	}

	return dp[0][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
