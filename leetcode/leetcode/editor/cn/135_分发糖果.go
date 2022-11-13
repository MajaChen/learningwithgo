package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, 0, n)
	for i := 0; i < n; i++ {
		candies = append(candies, 1)
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	ans := 0
	for _, candy := range candies {
		ans += candy
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
