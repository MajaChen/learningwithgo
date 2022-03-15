package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func getNumSum(nums []int, l, r int) int {
	var sum int
	for i := l; i <= r; i++ {
		sum += nums[i]
	}
	return sum
}

func getMax(i1, i2, i3 int) int {

	if i1 >= i2 && i1 >= i3 {
		return i1
	}

	if i2 >= i1 && i2 >= i3 {
		return i2
	}

	return i3
}
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {

	l, m, n := firstLen, secondLen, len(nums)
	suma, summ, suml, sum := make([]int, n+1), make([]int, n+1), make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		suma[i] = suma[i-1] + nums[i-1]

		summ[i] = summ[i-1]
		if i-m >= 0 {
			if j := getNumSum(nums, i-m, i-1); j > summ[i] {
				summ[i] = j
			}
		}

		suml[i] = suml[i-1]
		if i-l >= 0 {
			if j := getNumSum(nums, i-l, i-1); j > suml[i] {
				suml[i] = j
			}
		}

		var s1, s2, s3 int
		s1 = sum[i-1]

		if i-m >= 0 {
			s2 = suma[i] - suma[i-m] + suml[i-m]
		}

		if i-l >= 0 {
			s3 = suma[i] - suma[i-l] + summ[i-l]
		}

		sum[i] = getMax(s1, s2, s3)
	}

	return sum[n]
}

//leetcode submit region end(Prohibit modification and deletion)
