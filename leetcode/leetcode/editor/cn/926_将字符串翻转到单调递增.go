package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func minFlipsMonoIncr(s string) int {
	runes, counters, N := []rune(s), make([]int, len(s)+1), len(s)
	for i := 1; i <= N; i++ {
		counters[i] = counters[i-1] + int(runes[i-1]-48)
	}

	minAns := math.MaxInt32
	for i := 0; i <= N; i++ {
		if counter := counters[i] + (N - i) - (counters[N] - counters[i]); counter < minAns {
			minAns = counter
		}
	}

	return minAns
}

//leetcode submit region end(Prohibit modification and deletion)
