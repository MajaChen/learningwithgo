package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func maxAbsValExpr(arr1 []int, arr2 []int) int {

	sign, maxAbsVal := []int{-1, 1}, math.MinInt32
	for _, a := range sign {
		for _, b := range sign {
			for _, c := range sign {
				maxDistance, minDistance := math.MinInt32, math.MaxInt32
				for i := 0; i < len(arr1); i++ {
					k := arr1[i]*a + arr2[i]*b + i*c
					if k > maxDistance {
						maxDistance = k
					}
					if k < minDistance {
						minDistance = k
					}
				}
				if maxDistance-minDistance > maxAbsVal {
					maxAbsVal = maxDistance - minDistance
				}
			}
		}
	}

	return maxAbsVal
}

//leetcode submit region end(Prohibit modification and deletion)
