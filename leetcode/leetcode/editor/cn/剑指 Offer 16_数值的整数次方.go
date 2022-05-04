package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {

	if n < 0 {
		return 1 / myPowWithPositiveExpo(x, -n)
	} else {
		return myPowWithPositiveExpo(x, n)
	}
}

func myPowWithPositiveExpo(x float64, n int) float64 {
	mm := 1.0
	for n > 0 {
		c, m := 1, x
		for k := 1.0; c <= n; {
			m *= k
			k = m
			c *= 2
		}
		n -= c / 2
		mm *= m
	}

	return mm
}

//leetcode submit region end(Prohibit modification and deletion)
