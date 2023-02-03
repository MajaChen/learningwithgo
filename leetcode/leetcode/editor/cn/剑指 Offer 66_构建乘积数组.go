package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func constructArr(a []int) []int {
	n := len(a)
	if n <= 1 {
		return []int{}
	}
	l, r, ans := make([]int, n), make([]int, n), make([]int, n)
	l[0] = a[0]
	for i := 1; i < n; i++ {
		l[i] = l[i-1] * a[i]
	}
	r[n-1] = a[n-1]
	for i := n - 2; i >= 0; i-- {
		r[i] = r[i+1] * a[i]
	}
	for i := 1; i < n-1; i++ {
		ans[i] = l[i-1] * r[i+1]
	}
	ans[0] = r[1]
	ans[n-1] = l[n-2]
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
